package templates

import (
	"bytes"
	"fmt"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"github.com/99designs/gqlgen/internal/imports"
	"github.com/pkg/errors"
)

// this is done with a global because subtemplates currently get called in functions. Lets aim to remove this eventually.
var CurrentImports *Imports

type Options struct {
	PackageName     string
	Filename        string
	RegionTags      bool
	GeneratedHeader bool
	Data            interface{}
}

func Render(cfg Options) error {
	if CurrentImports != nil {
		panic(fmt.Errorf("recursive or concurrent call to RenderToFile detected"))
	}
	CurrentImports = &Imports{destDir: filepath.Dir(cfg.Filename)}

	// load path relative to calling source file
	_, callerFile, _, _ := runtime.Caller(1)
	rootDir := filepath.Dir(callerFile)

	t := template.New("").Funcs(Funcs())

	var roots []string
	// load all the templates in the directory
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		name := filepath.ToSlash(strings.TrimPrefix(path, rootDir+string(os.PathSeparator)))
		if !strings.HasSuffix(info.Name(), ".gotpl") {
			return nil
		}
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		t, err = t.New(name).Parse(string(b))
		if err != nil {
			return errors.Wrap(err, cfg.Filename)
		}

		roots = append(roots, name)

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "locating templates")
	}

	// then execute all the important looking ones in order, adding them to the same file
	sort.Slice(roots, func(i, j int) bool {
		// important files go first
		if strings.HasSuffix(roots[i], "!.gotpl") {
			return true
		}
		if strings.HasSuffix(roots[j], "!.gotpl") {
			return false
		}
		return roots[i] < roots[j]
	})
	var buf bytes.Buffer
	for _, root := range roots {
		if cfg.RegionTags {
			buf.WriteString("\n// region    " + center(70, "*", " "+root+" ") + "\n")
		}
		err = t.Lookup(root).Execute(&buf, cfg.Data)
		if err != nil {
			return errors.Wrap(err, root)
		}
		if cfg.RegionTags {
			buf.WriteString("\n// endregion " + center(70, "*", " "+root+" ") + "\n")
		}
	}

	var result bytes.Buffer
	if cfg.GeneratedHeader {
		result.WriteString("// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.\n\n")
	}
	result.WriteString("package ")
	result.WriteString(cfg.PackageName)
	result.WriteString("\n\n")
	result.WriteString("import (\n")
	result.WriteString(CurrentImports.String())
	result.WriteString(")\n")
	_, err = buf.WriteTo(&result)
	if err != nil {
		return err
	}
	CurrentImports = nil

	return write(cfg.Filename, result.Bytes())
}

func center(width int, pad string, s string) string {
	if len(s)+2 > width {
		return s
	}
	lpad := (width - len(s)) / 2
	rpad := width - (lpad + len(s))
	return strings.Repeat(pad, lpad) + s + strings.Repeat(pad, rpad)
}

func Funcs() template.FuncMap {
	return template.FuncMap{
		"ucFirst":       ucFirst,
		"lcFirst":       lcFirst,
		"quote":         strconv.Quote,
		"rawQuote":      rawQuote,
		"dump":          Dump,
		"ref":           ref,
		"ts":            TypeIdentifier,
		"call":          Call,
		"prefixLines":   prefixLines,
		"notNil":        notNil,
		"reserveImport": CurrentImports.Reserve,
		"lookupImport":  CurrentImports.Lookup,
		"go":            ToGo,
		"goPrivate":     ToGoPrivate,
		"add": func(a, b int) int {
			return a + b
		},
		"render": func(filename string, tpldata interface{}) (*bytes.Buffer, error) {
			return render(resolveName(filename, 0), tpldata)
		},
	}
}

func ucFirst(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func lcFirst(s string) string {
	if s == "" {
		return ""
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

func isDelimiter(c rune) bool {
	return c == '-' || c == '_' || unicode.IsSpace(c)
}

func ref(p types.Type) string {
	return CurrentImports.LookupType(p)
}

var pkgReplacer = strings.NewReplacer(
	"/", "ᚋ",
	".", "ᚗ",
	"-", "ᚑ",
)

func TypeIdentifier(t types.Type) string {
	res := ""
	for {
		switch it := t.(type) {
		case *types.Pointer:
			t.Underlying()
			res += "ᚖ"
			t = it.Elem()
		case *types.Slice:
			res += "ᚕ"
			t = it.Elem()
		case *types.Named:
			res += pkgReplacer.Replace(it.Obj().Pkg().Path())
			res += "ᚐ"
			res += it.Obj().Name()
			return res
		case *types.Basic:
			res += it.Name()
			return res
		case *types.Map:
			res += "map"
			return res
		case *types.Interface:
			res += "interface"
			return res
		default:
			panic(fmt.Errorf("unexpected type %T", it))
		}
	}
}

func Call(p *types.Func) string {
	pkg := CurrentImports.Lookup(p.Pkg().Path())

	if pkg != "" {
		pkg += "."
	}

	if p.Type() != nil {
		// make sure the returned type is listed in our imports.
		ref(p.Type().(*types.Signature).Results().At(0).Type())
	}

	return pkg + p.Name()
}

func ToGo(name string) string {
	runes := make([]rune, 0, len(name))

	wordWalker(name, func(word string, hasCommonInitial bool) {
		if !hasCommonInitial {
			word = ucFirst(strings.ToLower(word))
		}
		runes = append(runes, []rune(word)...)
	})

	return string(runes)
}

func ToGoPrivate(name string) string {
	runes := make([]rune, 0, len(name))

	first := true
	wordWalker(name, func(word string, hasCommonInitial bool) {
		if first {
			word = strings.ToLower(word)
			first = false
		} else if !hasCommonInitial {
			word = ucFirst(strings.ToLower(word))
		}
		runes = append(runes, []rune(word)...)
	})

	return sanitizeKeywords(string(runes))
}

func wordWalker(str string, f func(word string, hasCommonInitial bool)) {

	skipRune := func(r rune) bool {
		switch r {
		case '-', '_':
			return true
		default:
			return false
		}
	}

	runes := []rune(str)
	w, i := 0, 0 // index of start of word, scan
	hasCommonInitial := false
	for i+1 <= len(runes) {
		eow := false // whether we hit the end of a word
		if i+1 == len(runes) {
			eow = true
		} else if skipRune(runes[i+1]) {
			// underscore; shift the remainder forward over any run of underscores
			eow = true
			n := 1
			for i+n+1 < len(runes) && skipRune(runes[i+n+1]) {
				n++
			}

			// Leave at most one underscore if the underscore is between two digits
			if i+n+1 < len(runes) && unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i+n+1]) {
				n--
			}

			copy(runes[i+1:], runes[i+n+1:])
			runes = runes[:len(runes)-n]
		} else if unicode.IsLower(runes[i]) && !unicode.IsLower(runes[i+1]) {
			// lower->non-lower
			eow = true
		}
		i++

		// [w,i) is a word.
		word := string(runes[w:i])
		if !eow && commonInitialisms[word] && !unicode.IsLower(runes[i]) {
			// through
			// split IDFoo → ID, Foo
			// but URLs → URLs
		} else if !eow {
			if commonInitialisms[word] {
				hasCommonInitial = true
			}
			continue
		}

		if u := strings.ToUpper(word); commonInitialisms[u] {
			hasCommonInitial = true
			word = u
		}

		f(word, hasCommonInitial)
		hasCommonInitial = false
		w = i
	}
}

var keywords = []string{
	"break",
	"default",
	"func",
	"interface",
	"select",
	"case",
	"defer",
	"go",
	"map",
	"struct",
	"chan",
	"else",
	"goto",
	"package",
	"switch",
	"const",
	"fallthrough",
	"if",
	"range",
	"type",
	"continue",
	"for",
	"import",
	"return",
	"var",
	"_",
}

// sanitizeKeywords prevents collisions with go keywords for arguments to resolver functions
func sanitizeKeywords(name string) string {
	for _, k := range keywords {
		if name == k {
			return name + "Arg"
		}
	}
	return name
}

// copy from https://github.com/golang/lint/blob/06c8688daad7faa9da5a0c2f163a3d14aac986ca/lint.go#L679
func lintName(name string) string {
	// Fast path for simple cases: "_" and all lowercase.
	if name == "_" {
		return name
	}
	allLower := true
	for _, r := range name {
		if !unicode.IsLower(r) {
			allLower = false
			break
		}
	}
	if allLower {
		return name
	}

	// Split camelCase at any lower->upper transition, and split on underscores.
	// Check each word for common initialisms.
	runes := []rune(name)
	w, i := 0, 0 // index of start of word, scan
	for i+1 <= len(runes) {
		eow := false // whether we hit the end of a word
		if i+1 == len(runes) {
			eow = true
		} else if runes[i+1] == '_' {
			// underscore; shift the remainder forward over any run of underscores
			eow = true
			n := 1
			for i+n+1 < len(runes) && runes[i+n+1] == '_' {
				n++
			}

			// Leave at most one underscore if the underscore is between two digits
			if i+n+1 < len(runes) && unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i+n+1]) {
				n--
			}

			copy(runes[i+1:], runes[i+n+1:])
			runes = runes[:len(runes)-n]
		} else if unicode.IsLower(runes[i]) && !unicode.IsLower(runes[i+1]) {
			// lower->non-lower
			eow = true
		}
		i++
		if !eow {
			continue
		}

		// [w,i) is a word.
		word := string(runes[w:i])
		if u := strings.ToUpper(word); commonInitialisms[u] {
			// Keep consistent case, which is lowercase only at the start.
			if w == 0 && unicode.IsLower(runes[w]) {
				u = strings.ToLower(u)
			}
			// All the common initialisms are ASCII,
			// so we can replace the bytes exactly.
			copy(runes[w:], []rune(u))
		} else if w > 0 && strings.ToLower(word) == word {
			// already all lowercase, and not the first word, so uppercase the first character.
			runes[w] = unicode.ToUpper(runes[w])
		}
		w = i
	}
	return string(runes)
}

// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}

func rawQuote(s string) string {
	return "`" + strings.Replace(s, "`", "`+\"`\"+`", -1) + "`"
}

func notNil(field string, data interface{}) bool {
	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return false
	}
	val := v.FieldByName(field)

	return val.IsValid() && !val.IsNil()
}

func Dump(val interface{}) string {
	switch val := val.(type) {
	case int:
		return strconv.Itoa(val)
	case int64:
		return fmt.Sprintf("%d", val)
	case float64:
		return fmt.Sprintf("%f", val)
	case string:
		return strconv.Quote(val)
	case bool:
		return strconv.FormatBool(val)
	case nil:
		return "nil"
	case []interface{}:
		var parts []string
		for _, part := range val {
			parts = append(parts, Dump(part))
		}
		return "[]interface{}{" + strings.Join(parts, ",") + "}"
	case map[string]interface{}:
		buf := bytes.Buffer{}
		buf.WriteString("map[string]interface{}{")
		var keys []string
		for key := range val {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			data := val[key]

			buf.WriteString(strconv.Quote(key))
			buf.WriteString(":")
			buf.WriteString(Dump(data))
			buf.WriteString(",")
		}
		buf.WriteString("}")
		return buf.String()
	default:
		panic(fmt.Errorf("unsupported type %T", val))
	}
}

func prefixLines(prefix, s string) string {
	return prefix + strings.Replace(s, "\n", "\n"+prefix, -1)
}

func resolveName(name string, skip int) string {
	if name[0] == '.' {
		// load path relative to calling source file
		_, callerFile, _, _ := runtime.Caller(skip + 1)
		return filepath.Join(filepath.Dir(callerFile), name[1:])
	}

	// load path relative to this directory
	_, callerFile, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(callerFile), name)
}

func render(filename string, tpldata interface{}) (*bytes.Buffer, error) {
	t := template.New("").Funcs(Funcs())

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	t, err = t.New(filepath.Base(filename)).Parse(string(b))
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	return buf, t.Execute(buf, tpldata)
}

func write(filename string, b []byte) error {
	err := os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		return errors.Wrap(err, "failed to create directory")
	}

	formatted, err := imports.Prune(filename, b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gofmt failed on %s: %s\n", filepath.Base(filename), err.Error())
		formatted = b
	}

	err = ioutil.WriteFile(filename, formatted, 0644)
	if err != nil {
		return errors.Wrapf(err, "failed to write %s", filename)
	}

	return nil
}
