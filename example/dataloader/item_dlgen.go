// generated by github.com/vektah/dataloaden ; DO NOT EDIT

package dataloader

import (
	"sync"
	"time"
)

// ItemSliceLoader batches and caches requests
type ItemSliceLoader struct {
	// this method provides the data for the loader
	fetch func(keys []int) ([][]Item, []error)

	// how long to done before sending a batch
	wait time.Duration

	// this will limit the maximum number of keys to send in one batch, 0 = no limit
	maxBatch int

	// INTERNAL

	// lazily created cache
	cache map[int][]Item

	// the current batch. keys will continue to be collected until timeout is hit,
	// then everything will be sent to the fetch method and out to the listeners
	batch *itemSliceBatch

	// mutex to prevent races
	mu sync.Mutex
}

type itemSliceBatch struct {
	keys    []int
	data    [][]Item
	error   []error
	closing bool
	done    chan struct{}
}

// Load a item by key, batching and caching will be applied automatically
func (l *ItemSliceLoader) Load(key int) ([]Item, error) {
	return l.LoadThunk(key)()
}

// LoadThunk returns a function that when called will block waiting for a item.
// This method should be used if you want one goroutine to make requests to many
// different data loaders without blocking until the thunk is called.
func (l *ItemSliceLoader) LoadThunk(key int) func() ([]Item, error) {
	l.mu.Lock()
	if it, ok := l.cache[key]; ok {
		l.mu.Unlock()
		return func() ([]Item, error) {
			return it, nil
		}
	}
	if l.batch == nil {
		l.batch = &itemSliceBatch{done: make(chan struct{})}
	}
	batch := l.batch
	pos := batch.keyIndex(l, key)
	l.mu.Unlock()

	return func() ([]Item, error) {
		<-batch.done

		if batch.error[pos] == nil {
			l.mu.Lock()
			if l.cache == nil {
				l.cache = map[int][]Item{}
			}
			l.cache[key] = batch.data[pos]
			l.mu.Unlock()
		}

		return batch.data[pos], batch.error[pos]
	}
}

// LoadAll fetches many keys at once. It will be broken into appropriate sized
// sub batches depending on how the loader is configured
func (l *ItemSliceLoader) LoadAll(keys []int) ([][]Item, []error) {
	results := make([]func() ([]Item, error), len(keys))

	for i, key := range keys {
		results[i] = l.LoadThunk(key)
	}

	items := make([][]Item, len(keys))
	errors := make([]error, len(keys))
	for i, thunk := range results {
		items[i], errors[i] = thunk()
	}
	return items, errors
}

// Prime the cache with the provided key and value. If the key already exists, no change is made.
// (To forcefully prime the cache, clear the key first with loader.clear(key).prime(key, value).)
func (l *ItemSliceLoader) Prime(key int, value []Item) {
	l.mu.Lock()
	if _, found := l.cache[key]; !found {
		l.cache[key] = value
	}
	l.mu.Unlock()
}

// Clear the value at key from the cache, if it exists
func (l *ItemSliceLoader) Clear(key int) {
	l.mu.Lock()
	delete(l.cache, key)
	l.mu.Unlock()
}

// keyIndex will return the location of the key in the batch, if its not found
// it will add the key to the batch
func (b *itemSliceBatch) keyIndex(l *ItemSliceLoader, key int) int {
	for i, existingKey := range b.keys {
		if key == existingKey {
			return i
		}
	}

	pos := len(b.keys)
	b.keys = append(b.keys, key)
	if pos == 0 {
		go b.startTimer(l)
	}

	if l.maxBatch != 0 && pos >= l.maxBatch-1 {
		if !b.closing {
			b.closing = true
			l.batch = nil
			go b.end(l)
		}
	}

	return pos
}

func (b *itemSliceBatch) startTimer(l *ItemSliceLoader) {
	time.Sleep(l.wait)
	l.mu.Lock()

	// we must have hit a batch limit and are already finalizing this batch
	if b.closing {
		l.mu.Unlock()
		return
	}

	l.batch = nil
	l.mu.Unlock()

	b.end(l)
}

func (b *itemSliceBatch) end(l *ItemSliceLoader) {
	b.data, b.error = l.fetch(b.keys)
	close(b.done)
}
