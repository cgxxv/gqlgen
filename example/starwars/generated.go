package starwars

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
)

// region    ************************** generated!.gotpl **************************

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Droid() DroidResolver
	FriendsConnection() FriendsConnectionResolver
	Human() HumanResolver
	Mutation() MutationResolver
	Query() QueryResolver
	Starship() StarshipResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Droid struct {
		Id                func(childComplexity int) int
		Name              func(childComplexity int) int
		Friends           func(childComplexity int) int
		FriendsConnection func(childComplexity int, first *int, after *string) int
		AppearsIn         func(childComplexity int) int
		PrimaryFunction   func(childComplexity int) int
	}

	FriendsConnection struct {
		TotalCount func(childComplexity int) int
		Edges      func(childComplexity int) int
		Friends    func(childComplexity int) int
		PageInfo   func(childComplexity int) int
	}

	FriendsEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	Human struct {
		Id                func(childComplexity int) int
		Name              func(childComplexity int) int
		Height            func(childComplexity int, unit LengthUnit) int
		Mass              func(childComplexity int) int
		Friends           func(childComplexity int) int
		FriendsConnection func(childComplexity int, first *int, after *string) int
		AppearsIn         func(childComplexity int) int
		Starships         func(childComplexity int) int
	}

	Mutation struct {
		CreateReview func(childComplexity int, episode Episode, review Review) int
	}

	PageInfo struct {
		StartCursor func(childComplexity int) int
		EndCursor   func(childComplexity int) int
		HasNextPage func(childComplexity int) int
	}

	Query struct {
		Hero      func(childComplexity int, episode *Episode) int
		Reviews   func(childComplexity int, episode Episode, since *time.Time) int
		Search    func(childComplexity int, text string) int
		Character func(childComplexity int, id string) int
		Droid     func(childComplexity int, id string) int
		Human     func(childComplexity int, id string) int
		Starship  func(childComplexity int, id string) int
	}

	Review struct {
		Stars      func(childComplexity int) int
		Commentary func(childComplexity int) int
		Time       func(childComplexity int) int
	}

	Starship struct {
		Id      func(childComplexity int) int
		Name    func(childComplexity int) int
		Length  func(childComplexity int, unit *LengthUnit) int
		History func(childComplexity int) int
	}
}

type DroidResolver interface {
	Friends(ctx context.Context, obj *Droid) ([]Character, error)
	FriendsConnection(ctx context.Context, obj *Droid, first *int, after *string) (FriendsConnection, error)
}
type FriendsConnectionResolver interface {
	Edges(ctx context.Context, obj *FriendsConnection) ([]FriendsEdge, error)
	Friends(ctx context.Context, obj *FriendsConnection) ([]Character, error)
}
type HumanResolver interface {
	Friends(ctx context.Context, obj *Human) ([]Character, error)
	FriendsConnection(ctx context.Context, obj *Human, first *int, after *string) (FriendsConnection, error)

	Starships(ctx context.Context, obj *Human) ([]Starship, error)
}
type MutationResolver interface {
	CreateReview(ctx context.Context, episode Episode, review Review) (*Review, error)
}
type QueryResolver interface {
	Hero(ctx context.Context, episode *Episode) (Character, error)
	Reviews(ctx context.Context, episode Episode, since *time.Time) ([]Review, error)
	Search(ctx context.Context, text string) ([]SearchResult, error)
	Character(ctx context.Context, id string) (Character, error)
	Droid(ctx context.Context, id string) (*Droid, error)
	Human(ctx context.Context, id string) (*Human, error)
	Starship(ctx context.Context, id string) (*Starship, error)
}
type StarshipResolver interface {
	Length(ctx context.Context, obj *Starship, unit *LengthUnit) (float64, error)
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "Droid.id":
		if e.complexity.Droid.Id == nil {
			break
		}

		return e.complexity.Droid.Id(childComplexity), true

	case "Droid.name":
		if e.complexity.Droid.Name == nil {
			break
		}

		return e.complexity.Droid.Name(childComplexity), true

	case "Droid.friends":
		if e.complexity.Droid.Friends == nil {
			break
		}

		return e.complexity.Droid.Friends(childComplexity), true

	case "Droid.friendsConnection":
		if e.complexity.Droid.FriendsConnection == nil {
			break
		}

		args, err := ec.field_Droid_friendsConnection_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Droid.FriendsConnection(childComplexity, args["first"].(*int), args["after"].(*string)), true

	case "Droid.appearsIn":
		if e.complexity.Droid.AppearsIn == nil {
			break
		}

		return e.complexity.Droid.AppearsIn(childComplexity), true

	case "Droid.primaryFunction":
		if e.complexity.Droid.PrimaryFunction == nil {
			break
		}

		return e.complexity.Droid.PrimaryFunction(childComplexity), true

	case "FriendsConnection.totalCount":
		if e.complexity.FriendsConnection.TotalCount == nil {
			break
		}

		return e.complexity.FriendsConnection.TotalCount(childComplexity), true

	case "FriendsConnection.edges":
		if e.complexity.FriendsConnection.Edges == nil {
			break
		}

		return e.complexity.FriendsConnection.Edges(childComplexity), true

	case "FriendsConnection.friends":
		if e.complexity.FriendsConnection.Friends == nil {
			break
		}

		return e.complexity.FriendsConnection.Friends(childComplexity), true

	case "FriendsConnection.pageInfo":
		if e.complexity.FriendsConnection.PageInfo == nil {
			break
		}

		return e.complexity.FriendsConnection.PageInfo(childComplexity), true

	case "FriendsEdge.cursor":
		if e.complexity.FriendsEdge.Cursor == nil {
			break
		}

		return e.complexity.FriendsEdge.Cursor(childComplexity), true

	case "FriendsEdge.node":
		if e.complexity.FriendsEdge.Node == nil {
			break
		}

		return e.complexity.FriendsEdge.Node(childComplexity), true

	case "Human.id":
		if e.complexity.Human.Id == nil {
			break
		}

		return e.complexity.Human.Id(childComplexity), true

	case "Human.name":
		if e.complexity.Human.Name == nil {
			break
		}

		return e.complexity.Human.Name(childComplexity), true

	case "Human.height":
		if e.complexity.Human.Height == nil {
			break
		}

		args, err := ec.field_Human_height_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Human.Height(childComplexity, args["unit"].(LengthUnit)), true

	case "Human.mass":
		if e.complexity.Human.Mass == nil {
			break
		}

		return e.complexity.Human.Mass(childComplexity), true

	case "Human.friends":
		if e.complexity.Human.Friends == nil {
			break
		}

		return e.complexity.Human.Friends(childComplexity), true

	case "Human.friendsConnection":
		if e.complexity.Human.FriendsConnection == nil {
			break
		}

		args, err := ec.field_Human_friendsConnection_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Human.FriendsConnection(childComplexity, args["first"].(*int), args["after"].(*string)), true

	case "Human.appearsIn":
		if e.complexity.Human.AppearsIn == nil {
			break
		}

		return e.complexity.Human.AppearsIn(childComplexity), true

	case "Human.starships":
		if e.complexity.Human.Starships == nil {
			break
		}

		return e.complexity.Human.Starships(childComplexity), true

	case "Mutation.createReview":
		if e.complexity.Mutation.CreateReview == nil {
			break
		}

		args, err := ec.field_Mutation_createReview_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateReview(childComplexity, args["episode"].(Episode), args["review"].(Review)), true

	case "PageInfo.startCursor":
		if e.complexity.PageInfo.StartCursor == nil {
			break
		}

		return e.complexity.PageInfo.StartCursor(childComplexity), true

	case "PageInfo.endCursor":
		if e.complexity.PageInfo.EndCursor == nil {
			break
		}

		return e.complexity.PageInfo.EndCursor(childComplexity), true

	case "PageInfo.hasNextPage":
		if e.complexity.PageInfo.HasNextPage == nil {
			break
		}

		return e.complexity.PageInfo.HasNextPage(childComplexity), true

	case "Query.hero":
		if e.complexity.Query.Hero == nil {
			break
		}

		args, err := ec.field_Query_hero_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Hero(childComplexity, args["episode"].(*Episode)), true

	case "Query.reviews":
		if e.complexity.Query.Reviews == nil {
			break
		}

		args, err := ec.field_Query_reviews_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Reviews(childComplexity, args["episode"].(Episode), args["since"].(*time.Time)), true

	case "Query.search":
		if e.complexity.Query.Search == nil {
			break
		}

		args, err := ec.field_Query_search_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Search(childComplexity, args["text"].(string)), true

	case "Query.character":
		if e.complexity.Query.Character == nil {
			break
		}

		args, err := ec.field_Query_character_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Character(childComplexity, args["id"].(string)), true

	case "Query.droid":
		if e.complexity.Query.Droid == nil {
			break
		}

		args, err := ec.field_Query_droid_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Droid(childComplexity, args["id"].(string)), true

	case "Query.human":
		if e.complexity.Query.Human == nil {
			break
		}

		args, err := ec.field_Query_human_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Human(childComplexity, args["id"].(string)), true

	case "Query.starship":
		if e.complexity.Query.Starship == nil {
			break
		}

		args, err := ec.field_Query_starship_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Starship(childComplexity, args["id"].(string)), true

	case "Review.stars":
		if e.complexity.Review.Stars == nil {
			break
		}

		return e.complexity.Review.Stars(childComplexity), true

	case "Review.commentary":
		if e.complexity.Review.Commentary == nil {
			break
		}

		return e.complexity.Review.Commentary(childComplexity), true

	case "Review.time":
		if e.complexity.Review.Time == nil {
			break
		}

		return e.complexity.Review.Time(childComplexity), true

	case "Starship.id":
		if e.complexity.Starship.Id == nil {
			break
		}

		return e.complexity.Starship.Id(childComplexity), true

	case "Starship.name":
		if e.complexity.Starship.Name == nil {
			break
		}

		return e.complexity.Starship.Name(childComplexity), true

	case "Starship.length":
		if e.complexity.Starship.Length == nil {
			break
		}

		args, err := ec.field_Starship_length_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Starship.Length(childComplexity, args["unit"].(*LengthUnit)), true

	case "Starship.history":
		if e.complexity.Starship.History == nil {
			break
		}

		return e.complexity.Starship.History(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Query(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {
	ec := executionContext{graphql.GetRequestContext(ctx), e}

	buf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {
		data := ec._Query(ctx, op.SelectionSet)
		var buf bytes.Buffer
		data.MarshalGQL(&buf)
		return buf.Bytes()
	})

	return &graphql.Response{
		Data:       buf,
		Errors:     ec.Errors,
		Extensions: ec.Extensions,
	}
}

func (e *executableSchema) Mutation(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {
	ec := executionContext{graphql.GetRequestContext(ctx), e}

	buf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {
		data := ec._Mutation(ctx, op.SelectionSet)
		var buf bytes.Buffer
		data.MarshalGQL(&buf)
		return buf.Bytes()
	})

	return &graphql.Response{
		Data:       buf,
		Errors:     ec.Errors,
		Extensions: ec.Extensions,
	}
}

func (e *executableSchema) Subscription(ctx context.Context, op *ast.OperationDefinition) func() *graphql.Response {
	return graphql.OneShot(graphql.ErrorResponse(ctx, "subscriptions are not supported"))
}

type executionContext struct {
	*graphql.RequestContext
	*executableSchema
}

func (ec *executionContext) FieldMiddleware(ctx context.Context, obj interface{}, next graphql.Resolver) (ret interface{}) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = nil
		}
	}()
	res, err := ec.ResolverMiddleware(ctx, next)
	if err != nil {
		ec.Error(ctx, err)
		return nil
	}
	return res
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var parsedSchema = gqlparser.MustLoadSchema(
	&ast.Source{Name: "schema.graphql", Input: `# The query type, represents all of the entry points into our object graph
type Query {
    hero(episode: Episode = NEWHOPE): Character
    reviews(episode: Episode!, since: Time): [Review!]!
    search(text: String!): [SearchResult!]!
    character(id: ID!): Character
    droid(id: ID!): Droid
    human(id: ID!): Human
    starship(id: ID!): Starship
}
# The mutation type, represents all updates we can make to our data
type Mutation {
    createReview(episode: Episode!, review: ReviewInput!): Review
}
# The episodes in the Star Wars trilogy
enum Episode {
    # Star Wars Episode IV: A New Hope, released in 1977.
    NEWHOPE
    # Star Wars Episode V: The Empire Strikes Back, released in 1980.
    EMPIRE
    # Star Wars Episode VI: Return of the Jedi, released in 1983.
    JEDI
}
# A character from the Star Wars universe
interface Character {
    # The ID of the character
    id: ID!
    # The name of the character
    name: String!
    # The friends of the character, or an empty list if they have none
    friends: [Character!]
    # The friends of the character exposed as a connection with edges
    friendsConnection(first: Int, after: ID): FriendsConnection!
    # The movies this character appears in
    appearsIn: [Episode!]!
}
# Units of height
enum LengthUnit {
    # The standard unit around the world
    METER
    # Primarily used in the United States
    FOOT
}
# A humanoid creature from the Star Wars universe
type Human implements Character {
    # The ID of the human
    id: ID!
    # What this human calls themselves
    name: String!
    # Height in the preferred unit, default is meters
    height(unit: LengthUnit = METER): Float!
    # Mass in kilograms, or null if unknown
    mass: Float
    # This human's friends, or an empty list if they have none
    friends: [Character!]
    # The friends of the human exposed as a connection with edges
    friendsConnection(first: Int, after: ID): FriendsConnection!
    # The movies this human appears in
    appearsIn: [Episode!]!
    # A list of starships this person has piloted, or an empty list if none
    starships: [Starship!]
}
# An autonomous mechanical character in the Star Wars universe
type Droid implements Character {
    # The ID of the droid
    id: ID!
    # What others call this droid
    name: String!
    # This droid's friends, or an empty list if they have none
    friends: [Character!]
    # The friends of the droid exposed as a connection with edges
    friendsConnection(first: Int, after: ID): FriendsConnection!
    # The movies this droid appears in
    appearsIn: [Episode!]!
    # This droid's primary function
    primaryFunction: String
}
# A connection object for a character's friends
type FriendsConnection {
    # The total number of friends
    totalCount: Int!
    # The edges for each of the character's friends.
    edges: [FriendsEdge!]
    # A list of the friends, as a convenience when edges are not needed.
    friends: [Character!]
    # Information for paginating this connection
    pageInfo: PageInfo!
}
# An edge object for a character's friends
type FriendsEdge {
    # A cursor used for pagination
    cursor: ID!
    # The character represented by this friendship edge
    node: Character
}
# Information for paginating this connection
type PageInfo {
    startCursor: ID!
    endCursor: ID!
    hasNextPage: Boolean!
}
# Represents a review for a movie
type Review {
    # The number of stars this review gave, 1-5
    stars: Int!
    # Comment about the movie
    commentary: String
    # when the review was posted
    time: Time
}
# The input object sent when someone is creating a new review
input ReviewInput {
    # 0-5 stars
    stars: Int!
    # Comment about the movie, optional
    commentary: String
    # when the review was posted
    time: Time
}
type Starship {
    # The ID of the starship
    id: ID!
    # The name of the starship
    name: String!
    # Length of the starship, along the longest axis
    length(unit: LengthUnit = METER): Float!
    # coordinates tracking this ship
    history: [[Int!]!]!
}
union SearchResult = Human | Droid | Starship
scalar Time
`},
)

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_Droid_friendsConnection_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *int
	if tmp, ok := rawArgs["first"]; ok {
		arg0, err = ec.unmarshalOInt2ᚖint(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["first"] = arg0
	var arg1 *string
	if tmp, ok := rawArgs["after"]; ok {
		arg1, err = ec.unmarshalOID2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["after"] = arg1
	return args, nil
}

func (ec *executionContext) field_Human_friendsConnection_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *int
	if tmp, ok := rawArgs["first"]; ok {
		arg0, err = ec.unmarshalOInt2ᚖint(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["first"] = arg0
	var arg1 *string
	if tmp, ok := rawArgs["after"]; ok {
		arg1, err = ec.unmarshalOID2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["after"] = arg1
	return args, nil
}

func (ec *executionContext) field_Human_height_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 LengthUnit
	if tmp, ok := rawArgs["unit"]; ok {
		arg0, err = ec.unmarshalOLengthUnit2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐLengthUnit(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["unit"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_createReview_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 Episode
	if tmp, ok := rawArgs["episode"]; ok {
		arg0, err = ec.unmarshalNEpisode2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐEpisode(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["episode"] = arg0
	var arg1 Review
	if tmp, ok := rawArgs["review"]; ok {
		arg1, err = ec.unmarshalNReviewInput2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐReview(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["review"] = arg1
	return args, nil
}

func (ec *executionContext) field_Query___type_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["name"]; ok {
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["name"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_character_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_droid_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_hero_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *Episode
	if tmp, ok := rawArgs["episode"]; ok {
		arg0, err = ec.unmarshalOEpisode2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐEpisode(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["episode"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_human_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_reviews_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 Episode
	if tmp, ok := rawArgs["episode"]; ok {
		arg0, err = ec.unmarshalNEpisode2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐEpisode(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["episode"] = arg0
	var arg1 *time.Time
	if tmp, ok := rawArgs["since"]; ok {
		arg1, err = ec.unmarshalOTime2ᚖtimeᚐTime(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["since"] = arg1
	return args, nil
}

func (ec *executionContext) field_Query_search_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["text"]; ok {
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["text"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_starship_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Starship_length_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *LengthUnit
	if tmp, ok := rawArgs["unit"]; ok {
		arg0, err = ec.unmarshalOLengthUnit2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐLengthUnit(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["unit"] = arg0
	return args, nil
}

func (ec *executionContext) field___Type_enumValues_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		arg0, err = ec.unmarshalOBoolean2bool(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil
}

func (ec *executionContext) field___Type_fields_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		arg0, err = ec.unmarshalOBoolean2bool(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Droid_id(ctx context.Context, field graphql.CollectedField, obj *Droid) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Droid",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Droid_name(ctx context.Context, field graphql.CollectedField, obj *Droid) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Droid",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Droid_friends(ctx context.Context, field graphql.CollectedField, obj *Droid) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Droid",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Droid().Friends(rctx, obj)
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]Character)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOCharacter2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐCharacter(ctx, field.Selections, res)
}

func (ec *executionContext) _Droid_friendsConnection(ctx context.Context, field graphql.CollectedField, obj *Droid) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Droid",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Droid_friendsConnection_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Droid().FriendsConnection(rctx, obj, args["first"].(*int), args["after"].(*string))
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(FriendsConnection)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNFriendsConnection2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐFriendsConnection(ctx, field.Selections, res)
}

func (ec *executionContext) _Droid_appearsIn(ctx context.Context, field graphql.CollectedField, obj *Droid) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Droid",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.AppearsIn, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]Episode)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNEpisode2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐEpisode(ctx, field.Selections, res)
}

func (ec *executionContext) _Droid_primaryFunction(ctx context.Context, field graphql.CollectedField, obj *Droid) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Droid",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PrimaryFunction, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _FriendsConnection_totalCount(ctx context.Context, field graphql.CollectedField, obj *FriendsConnection) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "FriendsConnection",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.TotalCount(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _FriendsConnection_edges(ctx context.Context, field graphql.CollectedField, obj *FriendsConnection) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "FriendsConnection",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.FriendsConnection().Edges(rctx, obj)
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]FriendsEdge)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOFriendsEdge2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐFriendsEdge(ctx, field.Selections, res)
}

func (ec *executionContext) _FriendsConnection_friends(ctx context.Context, field graphql.CollectedField, obj *FriendsConnection) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "FriendsConnection",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.FriendsConnection().Friends(rctx, obj)
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]Character)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOCharacter2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐCharacter(ctx, field.Selections, res)
}

func (ec *executionContext) _FriendsConnection_pageInfo(ctx context.Context, field graphql.CollectedField, obj *FriendsConnection) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "FriendsConnection",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PageInfo(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(PageInfo)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNPageInfo2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐPageInfo(ctx, field.Selections, res)
}

func (ec *executionContext) _FriendsEdge_cursor(ctx context.Context, field graphql.CollectedField, obj *FriendsEdge) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "FriendsEdge",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Cursor, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _FriendsEdge_node(ctx context.Context, field graphql.CollectedField, obj *FriendsEdge) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "FriendsEdge",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Node, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(Character)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOCharacter2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐCharacter(ctx, field.Selections, res)
}

func (ec *executionContext) _Human_id(ctx context.Context, field graphql.CollectedField, obj *Human) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Human",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Human_name(ctx context.Context, field graphql.CollectedField, obj *Human) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Human",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Human_height(ctx context.Context, field graphql.CollectedField, obj *Human) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Human",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Human_height_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Height(args["unit"].(LengthUnit)), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(float64)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) _Human_mass(ctx context.Context, field graphql.CollectedField, obj *Human) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Human",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Mass, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(float64)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) _Human_friends(ctx context.Context, field graphql.CollectedField, obj *Human) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Human",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Human().Friends(rctx, obj)
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]Character)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOCharacter2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐCharacter(ctx, field.Selections, res)
}

func (ec *executionContext) _Human_friendsConnection(ctx context.Context, field graphql.CollectedField, obj *Human) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Human",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Human_friendsConnection_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Human().FriendsConnection(rctx, obj, args["first"].(*int), args["after"].(*string))
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(FriendsConnection)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNFriendsConnection2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐFriendsConnection(ctx, field.Selections, res)
}

func (ec *executionContext) _Human_appearsIn(ctx context.Context, field graphql.CollectedField, obj *Human) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Human",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.AppearsIn, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]Episode)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNEpisode2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐEpisode(ctx, field.Selections, res)
}

func (ec *executionContext) _Human_starships(ctx context.Context, field graphql.CollectedField, obj *Human) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Human",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Human().Starships(rctx, obj)
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]Starship)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOStarship2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐStarship(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_createReview(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Mutation",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_createReview_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().CreateReview(rctx, args["episode"].(Episode), args["review"].(Review))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Review)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOReview2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐReview(ctx, field.Selections, res)
}

func (ec *executionContext) _PageInfo_startCursor(ctx context.Context, field graphql.CollectedField, obj *PageInfo) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "PageInfo",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.StartCursor, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _PageInfo_endCursor(ctx context.Context, field graphql.CollectedField, obj *PageInfo) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "PageInfo",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.EndCursor, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _PageInfo_hasNextPage(ctx context.Context, field graphql.CollectedField, obj *PageInfo) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "PageInfo",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.HasNextPage, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_hero(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_hero_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Hero(rctx, args["episode"].(*Episode))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(Character)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOCharacter2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐCharacter(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_reviews(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_reviews_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Reviews(rctx, args["episode"].(Episode), args["since"].(*time.Time))
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]Review)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNReview2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐReview(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_search(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_search_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Search(rctx, args["text"].(string))
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]SearchResult)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNSearchResult2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐSearchResult(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_character(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_character_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Character(rctx, args["id"].(string))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(Character)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOCharacter2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐCharacter(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_droid(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_droid_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Droid(rctx, args["id"].(string))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Droid)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalODroid2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐDroid(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_human(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_human_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Human(rctx, args["id"].(string))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Human)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOHuman2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐHuman(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_starship(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_starship_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Starship(rctx, args["id"].(string))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Starship)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOStarship2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐStarship(ctx, field.Selections, res)
}

func (ec *executionContext) _Query___type(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query___type_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectType(args["name"].(string))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) _Query___schema(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectSchema()
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Schema)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__Schema2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx, field.Selections, res)
}

func (ec *executionContext) _Review_stars(ctx context.Context, field graphql.CollectedField, obj *Review) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Review",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Stars, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _Review_commentary(ctx context.Context, field graphql.CollectedField, obj *Review) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Review",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Commentary, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _Review_time(ctx context.Context, field graphql.CollectedField, obj *Review) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Review",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Time, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(time.Time)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) _Starship_id(ctx context.Context, field graphql.CollectedField, obj *Starship) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Starship",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Starship_name(ctx context.Context, field graphql.CollectedField, obj *Starship) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Starship",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Starship_length(ctx context.Context, field graphql.CollectedField, obj *Starship) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Starship",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Starship_length_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Starship().Length(rctx, obj, args["unit"].(*LengthUnit))
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(float64)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) _Starship_history(ctx context.Context, field graphql.CollectedField, obj *Starship) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Starship",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.History, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([][]int)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNInt2ᚕᚕint(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_locations(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Locations, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalN__DirectiveLocation2ᚕstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_type(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_type(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_defaultValue(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DefaultValue, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_types(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Types(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalN__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_queryType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.QueryType(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_mutationType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.MutationType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_subscriptionType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SubscriptionType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_directives(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Directives(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Directive)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalN__Directive2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_kind(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Kind(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalN__TypeKind2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_fields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field___Type_fields_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Fields(args["includeDeprecated"].(bool)), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Field)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__Field2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_interfaces(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Interfaces(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_possibleTypes(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PossibleTypes(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_enumValues(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field___Type_enumValues_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.EnumValues(args["includeDeprecated"].(bool)), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.EnumValue)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__EnumValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_inputFields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.InputFields(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_ofType(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.OfType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputReviewInput(ctx context.Context, v interface{}) (Review, error) {
	var it Review
	var asMap = v.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "stars":
			var err error
			it.Stars, err = ec.unmarshalNInt2int(ctx, v)
			if err != nil {
				return it, err
			}
		case "commentary":
			var err error
			it.Commentary, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "time":
			var err error
			it.Time, err = ec.unmarshalOTime2timeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

func (ec *executionContext) _Character(ctx context.Context, sel ast.SelectionSet, obj *Character) graphql.Marshaler {
	switch obj := (*obj).(type) {
	case nil:
		return graphql.Null
	case Human:
		return ec._Human(ctx, sel, &obj)
	case *Human:
		return ec._Human(ctx, sel, obj)
	case Droid:
		return ec._Droid(ctx, sel, &obj)
	case *Droid:
		return ec._Droid(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

func (ec *executionContext) _SearchResult(ctx context.Context, sel ast.SelectionSet, obj *SearchResult) graphql.Marshaler {
	switch obj := (*obj).(type) {
	case nil:
		return graphql.Null
	case Human:
		return ec._Human(ctx, sel, &obj)
	case *Human:
		return ec._Human(ctx, sel, obj)
	case Droid:
		return ec._Droid(ctx, sel, &obj)
	case *Droid:
		return ec._Droid(ctx, sel, obj)
	case Starship:
		return ec._Starship(ctx, sel, &obj)
	case *Starship:
		return ec._Starship(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var droidImplementors = []string{"Droid", "Character"}

func (ec *executionContext) _Droid(ctx context.Context, sel ast.SelectionSet, obj *Droid) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, droidImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Droid")
		case "id":
			out.Values[i] = ec._Droid_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "name":
			out.Values[i] = ec._Droid_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "friends":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Droid_friends(ctx, field, obj)
				return res
			})
		case "friendsConnection":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Droid_friendsConnection(ctx, field, obj)
				if res == graphql.Null {
					invalid = true
				}
				return res
			})
		case "appearsIn":
			out.Values[i] = ec._Droid_appearsIn(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "primaryFunction":
			out.Values[i] = ec._Droid_primaryFunction(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var friendsConnectionImplementors = []string{"FriendsConnection"}

func (ec *executionContext) _FriendsConnection(ctx context.Context, sel ast.SelectionSet, obj *FriendsConnection) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, friendsConnectionImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("FriendsConnection")
		case "totalCount":
			out.Values[i] = ec._FriendsConnection_totalCount(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "edges":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._FriendsConnection_edges(ctx, field, obj)
				return res
			})
		case "friends":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._FriendsConnection_friends(ctx, field, obj)
				return res
			})
		case "pageInfo":
			out.Values[i] = ec._FriendsConnection_pageInfo(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var friendsEdgeImplementors = []string{"FriendsEdge"}

func (ec *executionContext) _FriendsEdge(ctx context.Context, sel ast.SelectionSet, obj *FriendsEdge) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, friendsEdgeImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("FriendsEdge")
		case "cursor":
			out.Values[i] = ec._FriendsEdge_cursor(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "node":
			out.Values[i] = ec._FriendsEdge_node(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var humanImplementors = []string{"Human", "Character"}

func (ec *executionContext) _Human(ctx context.Context, sel ast.SelectionSet, obj *Human) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, humanImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Human")
		case "id":
			out.Values[i] = ec._Human_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "name":
			out.Values[i] = ec._Human_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "height":
			out.Values[i] = ec._Human_height(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "mass":
			out.Values[i] = ec._Human_mass(ctx, field, obj)
		case "friends":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Human_friends(ctx, field, obj)
				return res
			})
		case "friendsConnection":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Human_friendsConnection(ctx, field, obj)
				if res == graphql.Null {
					invalid = true
				}
				return res
			})
		case "appearsIn":
			out.Values[i] = ec._Human_appearsIn(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "starships":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Human_starships(ctx, field, obj)
				return res
			})
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var mutationImplementors = []string{"Mutation"}

func (ec *executionContext) _Mutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, mutationImplementors)

	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "Mutation",
	})

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Mutation")
		case "createReview":
			out.Values[i] = ec._Mutation_createReview(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var pageInfoImplementors = []string{"PageInfo"}

func (ec *executionContext) _PageInfo(ctx context.Context, sel ast.SelectionSet, obj *PageInfo) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, pageInfoImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("PageInfo")
		case "startCursor":
			out.Values[i] = ec._PageInfo_startCursor(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "endCursor":
			out.Values[i] = ec._PageInfo_endCursor(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "hasNextPage":
			out.Values[i] = ec._PageInfo_hasNextPage(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var queryImplementors = []string{"Query"}

func (ec *executionContext) _Query(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, queryImplementors)

	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "Query",
	})

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "hero":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_hero(ctx, field)
				return res
			})
		case "reviews":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_reviews(ctx, field)
				if res == graphql.Null {
					invalid = true
				}
				return res
			})
		case "search":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_search(ctx, field)
				if res == graphql.Null {
					invalid = true
				}
				return res
			})
		case "character":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_character(ctx, field)
				return res
			})
		case "droid":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_droid(ctx, field)
				return res
			})
		case "human":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_human(ctx, field)
				return res
			})
		case "starship":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_starship(ctx, field)
				return res
			})
		case "__type":
			out.Values[i] = ec._Query___type(ctx, field)
		case "__schema":
			out.Values[i] = ec._Query___schema(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var reviewImplementors = []string{"Review"}

func (ec *executionContext) _Review(ctx context.Context, sel ast.SelectionSet, obj *Review) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, reviewImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Review")
		case "stars":
			out.Values[i] = ec._Review_stars(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "commentary":
			out.Values[i] = ec._Review_commentary(ctx, field, obj)
		case "time":
			out.Values[i] = ec._Review_time(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var starshipImplementors = []string{"Starship"}

func (ec *executionContext) _Starship(ctx context.Context, sel ast.SelectionSet, obj *Starship) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, starshipImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Starship")
		case "id":
			out.Values[i] = ec._Starship_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "name":
			out.Values[i] = ec._Starship_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "length":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Starship_length(ctx, field, obj)
				if res == graphql.Null {
					invalid = true
				}
				return res
			})
		case "history":
			out.Values[i] = ec._Starship_history(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var __DirectiveImplementors = []string{"__Directive"}

func (ec *executionContext) ___Directive(ctx context.Context, sel ast.SelectionSet, obj *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __DirectiveImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":
			out.Values[i] = ec.___Directive_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___Directive_description(ctx, field, obj)
		case "locations":
			out.Values[i] = ec.___Directive_locations(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "args":
			out.Values[i] = ec.___Directive_args(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var __EnumValueImplementors = []string{"__EnumValue"}

func (ec *executionContext) ___EnumValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __EnumValueImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":
			out.Values[i] = ec.___EnumValue_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___EnumValue_description(ctx, field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___EnumValue_isDeprecated(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "deprecationReason":
			out.Values[i] = ec.___EnumValue_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var __FieldImplementors = []string{"__Field"}

func (ec *executionContext) ___Field(ctx context.Context, sel ast.SelectionSet, obj *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __FieldImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":
			out.Values[i] = ec.___Field_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___Field_description(ctx, field, obj)
		case "args":
			out.Values[i] = ec.___Field_args(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "type":
			out.Values[i] = ec.___Field_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "isDeprecated":
			out.Values[i] = ec.___Field_isDeprecated(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "deprecationReason":
			out.Values[i] = ec.___Field_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var __InputValueImplementors = []string{"__InputValue"}

func (ec *executionContext) ___InputValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __InputValueImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":
			out.Values[i] = ec.___InputValue_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___InputValue_description(ctx, field, obj)
		case "type":
			out.Values[i] = ec.___InputValue_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "defaultValue":
			out.Values[i] = ec.___InputValue_defaultValue(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var __SchemaImplementors = []string{"__Schema"}

func (ec *executionContext) ___Schema(ctx context.Context, sel ast.SelectionSet, obj *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __SchemaImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "types":
			out.Values[i] = ec.___Schema_types(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "queryType":
			out.Values[i] = ec.___Schema_queryType(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "mutationType":
			out.Values[i] = ec.___Schema_mutationType(ctx, field, obj)
		case "subscriptionType":
			out.Values[i] = ec.___Schema_subscriptionType(ctx, field, obj)
		case "directives":
			out.Values[i] = ec.___Schema_directives(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var __TypeImplementors = []string{"__Type"}

func (ec *executionContext) ___Type(ctx context.Context, sel ast.SelectionSet, obj *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __TypeImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":
			out.Values[i] = ec.___Type_kind(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "name":
			out.Values[i] = ec.___Type_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Type_description(ctx, field, obj)
		case "fields":
			out.Values[i] = ec.___Type_fields(ctx, field, obj)
		case "interfaces":
			out.Values[i] = ec.___Type_interfaces(ctx, field, obj)
		case "possibleTypes":
			out.Values[i] = ec.___Type_possibleTypes(ctx, field, obj)
		case "enumValues":
			out.Values[i] = ec.___Type_enumValues(ctx, field, obj)
		case "inputFields":
			out.Values[i] = ec.___Type_inputFields(ctx, field, obj)
		case "ofType":
			out.Values[i] = ec.___Type_ofType(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNBoolean2bool(ctx context.Context, v interface{}) (bool, error) {
	return graphql.UnmarshalBoolean(v)
}

func (ec *executionContext) marshalNBoolean2bool(ctx context.Context, sel ast.SelectionSet, v bool) graphql.Marshaler {
	return graphql.MarshalBoolean(v)
}

func (ec *executionContext) marshalNCharacter2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐCharacter(ctx context.Context, sel ast.SelectionSet, v Character) graphql.Marshaler {
	return ec._Character(ctx, sel, &v)
}

func (ec *executionContext) unmarshalNEpisode2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐEpisode(ctx context.Context, v interface{}) (Episode, error) {
	var res Episode
	return res, res.UnmarshalGQL(v)
}

func (ec *executionContext) marshalNEpisode2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐEpisode(ctx context.Context, sel ast.SelectionSet, v Episode) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalNEpisode2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐEpisode(ctx context.Context, v interface{}) ([]Episode, error) {
	var vSlice []interface{}
	if v != nil {
		if tmp1, ok := v.([]interface{}); ok {
			vSlice = tmp1
		} else {
			vSlice = []interface{}{v}
		}
	}
	var err error
	res := make([]Episode, len(vSlice))
	for i := range vSlice {
		res[i], err = ec.unmarshalNEpisode2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐEpisode(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalNEpisode2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐEpisode(ctx context.Context, sel ast.SelectionSet, v []Episode) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNEpisode2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐEpisode(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) unmarshalNFloat2float64(ctx context.Context, v interface{}) (float64, error) {
	return graphql.UnmarshalFloat(v)
}

func (ec *executionContext) marshalNFloat2float64(ctx context.Context, sel ast.SelectionSet, v float64) graphql.Marshaler {
	return graphql.MarshalFloat(v)
}

func (ec *executionContext) marshalNFriendsConnection2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐFriendsConnection(ctx context.Context, sel ast.SelectionSet, v FriendsConnection) graphql.Marshaler {
	return ec._FriendsConnection(ctx, sel, &v)
}

func (ec *executionContext) marshalNFriendsEdge2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐFriendsEdge(ctx context.Context, sel ast.SelectionSet, v FriendsEdge) graphql.Marshaler {
	return ec._FriendsEdge(ctx, sel, &v)
}

func (ec *executionContext) unmarshalNID2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalID(v)
}

func (ec *executionContext) marshalNID2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	return graphql.MarshalID(v)
}

func (ec *executionContext) unmarshalNInt2int(ctx context.Context, v interface{}) (int, error) {
	return graphql.UnmarshalInt(v)
}

func (ec *executionContext) marshalNInt2int(ctx context.Context, sel ast.SelectionSet, v int) graphql.Marshaler {
	return graphql.MarshalInt(v)
}

func (ec *executionContext) unmarshalNInt2ᚕint(ctx context.Context, v interface{}) ([]int, error) {
	var vSlice []interface{}
	if v != nil {
		if tmp1, ok := v.([]interface{}); ok {
			vSlice = tmp1
		} else {
			vSlice = []interface{}{v}
		}
	}
	var err error
	res := make([]int, len(vSlice))
	for i := range vSlice {
		res[i], err = ec.unmarshalNInt2int(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalNInt2ᚕint(ctx context.Context, sel ast.SelectionSet, v []int) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	for i := range v {
		ret[i] = ec.marshalNInt2int(ctx, sel, v[i])
	}

	return ret
}

func (ec *executionContext) unmarshalNInt2ᚕᚕint(ctx context.Context, v interface{}) ([][]int, error) {
	var vSlice []interface{}
	if v != nil {
		if tmp1, ok := v.([]interface{}); ok {
			vSlice = tmp1
		} else {
			vSlice = []interface{}{v}
		}
	}
	var err error
	res := make([][]int, len(vSlice))
	for i := range vSlice {
		res[i], err = ec.unmarshalNInt2ᚕint(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalNInt2ᚕᚕint(ctx context.Context, sel ast.SelectionSet, v [][]int) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	for i := range v {
		ret[i] = ec.marshalNInt2ᚕint(ctx, sel, v[i])
	}

	return ret
}

func (ec *executionContext) marshalNPageInfo2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐPageInfo(ctx context.Context, sel ast.SelectionSet, v PageInfo) graphql.Marshaler {
	return ec._PageInfo(ctx, sel, &v)
}

func (ec *executionContext) marshalNReview2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐReview(ctx context.Context, sel ast.SelectionSet, v Review) graphql.Marshaler {
	return ec._Review(ctx, sel, &v)
}

func (ec *executionContext) marshalNReview2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐReview(ctx context.Context, sel ast.SelectionSet, v []Review) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNReview2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐReview(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) unmarshalNReviewInput2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐReview(ctx context.Context, v interface{}) (Review, error) {
	return ec.unmarshalInputReviewInput(ctx, v)
}

func (ec *executionContext) marshalNSearchResult2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐSearchResult(ctx context.Context, sel ast.SelectionSet, v SearchResult) graphql.Marshaler {
	return ec._SearchResult(ctx, sel, &v)
}

func (ec *executionContext) marshalNSearchResult2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐSearchResult(ctx context.Context, sel ast.SelectionSet, v []SearchResult) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNSearchResult2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐSearchResult(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalNStarship2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐStarship(ctx context.Context, sel ast.SelectionSet, v Starship) graphql.Marshaler {
	return ec._Starship(ctx, sel, &v)
}

func (ec *executionContext) unmarshalNString2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalNString2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	return graphql.MarshalString(v)
}

func (ec *executionContext) marshalN__Directive2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx context.Context, sel ast.SelectionSet, v introspection.Directive) graphql.Marshaler {
	return ec.___Directive(ctx, sel, &v)
}

func (ec *executionContext) marshalN__Directive2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx context.Context, sel ast.SelectionSet, v []introspection.Directive) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Directive2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) unmarshalN__DirectiveLocation2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalN__DirectiveLocation2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	return graphql.MarshalString(v)
}

func (ec *executionContext) unmarshalN__DirectiveLocation2ᚕstring(ctx context.Context, v interface{}) ([]string, error) {
	var vSlice []interface{}
	if v != nil {
		if tmp1, ok := v.([]interface{}); ok {
			vSlice = tmp1
		} else {
			vSlice = []interface{}{v}
		}
	}
	var err error
	res := make([]string, len(vSlice))
	for i := range vSlice {
		res[i], err = ec.unmarshalN__DirectiveLocation2string(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalN__DirectiveLocation2ᚕstring(ctx context.Context, sel ast.SelectionSet, v []string) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__DirectiveLocation2string(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalN__EnumValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx context.Context, sel ast.SelectionSet, v introspection.EnumValue) graphql.Marshaler {
	return ec.___EnumValue(ctx, sel, &v)
}

func (ec *executionContext) marshalN__Field2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx context.Context, sel ast.SelectionSet, v introspection.Field) graphql.Marshaler {
	return ec.___Field(ctx, sel, &v)
}

func (ec *executionContext) marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx context.Context, sel ast.SelectionSet, v introspection.InputValue) graphql.Marshaler {
	return ec.___InputValue(ctx, sel, &v)
}

func (ec *executionContext) marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx context.Context, sel ast.SelectionSet, v []introspection.InputValue) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v introspection.Type) graphql.Marshaler {
	return ec.___Type(ctx, sel, &v)
}

func (ec *executionContext) marshalN__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v []introspection.Type) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v *introspection.Type) graphql.Marshaler {
	if v == nil {
		if !ec.HasError(graphql.GetResolverContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec.___Type(ctx, sel, v)
}

func (ec *executionContext) unmarshalN__TypeKind2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalN__TypeKind2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	return graphql.MarshalString(v)
}

func (ec *executionContext) unmarshalOBoolean2bool(ctx context.Context, v interface{}) (bool, error) {
	return graphql.UnmarshalBoolean(v)
}

func (ec *executionContext) marshalOBoolean2bool(ctx context.Context, sel ast.SelectionSet, v bool) graphql.Marshaler {
	return graphql.MarshalBoolean(v)
}

func (ec *executionContext) marshalOCharacter2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐCharacter(ctx context.Context, sel ast.SelectionSet, v Character) graphql.Marshaler {
	return ec._Character(ctx, sel, &v)
}

func (ec *executionContext) marshalOCharacter2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐCharacter(ctx context.Context, sel ast.SelectionSet, v []Character) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNCharacter2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐCharacter(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalODroid2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐDroid(ctx context.Context, sel ast.SelectionSet, v Droid) graphql.Marshaler {
	return ec._Droid(ctx, sel, &v)
}

func (ec *executionContext) marshalODroid2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐDroid(ctx context.Context, sel ast.SelectionSet, v *Droid) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Droid(ctx, sel, v)
}

func (ec *executionContext) unmarshalOEpisode2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐEpisode(ctx context.Context, v interface{}) (Episode, error) {
	var res Episode
	return res, res.UnmarshalGQL(v)
}

func (ec *executionContext) marshalOEpisode2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐEpisode(ctx context.Context, sel ast.SelectionSet, v Episode) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalOEpisode2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐEpisode(ctx context.Context, v interface{}) (*Episode, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOEpisode2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐEpisode(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOEpisode2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐEpisode(ctx context.Context, sel ast.SelectionSet, v *Episode) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return v
}

func (ec *executionContext) unmarshalOFloat2float64(ctx context.Context, v interface{}) (float64, error) {
	return graphql.UnmarshalFloat(v)
}

func (ec *executionContext) marshalOFloat2float64(ctx context.Context, sel ast.SelectionSet, v float64) graphql.Marshaler {
	return graphql.MarshalFloat(v)
}

func (ec *executionContext) marshalOFriendsEdge2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐFriendsEdge(ctx context.Context, sel ast.SelectionSet, v []FriendsEdge) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNFriendsEdge2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐFriendsEdge(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalOHuman2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐHuman(ctx context.Context, sel ast.SelectionSet, v Human) graphql.Marshaler {
	return ec._Human(ctx, sel, &v)
}

func (ec *executionContext) marshalOHuman2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐHuman(ctx context.Context, sel ast.SelectionSet, v *Human) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Human(ctx, sel, v)
}

func (ec *executionContext) unmarshalOID2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalID(v)
}

func (ec *executionContext) marshalOID2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	return graphql.MarshalID(v)
}

func (ec *executionContext) unmarshalOID2ᚖstring(ctx context.Context, v interface{}) (*string, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOID2string(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOID2ᚖstring(ctx context.Context, sel ast.SelectionSet, v *string) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOID2string(ctx, sel, *v)
}

func (ec *executionContext) unmarshalOInt2int(ctx context.Context, v interface{}) (int, error) {
	return graphql.UnmarshalInt(v)
}

func (ec *executionContext) marshalOInt2int(ctx context.Context, sel ast.SelectionSet, v int) graphql.Marshaler {
	return graphql.MarshalInt(v)
}

func (ec *executionContext) unmarshalOInt2ᚖint(ctx context.Context, v interface{}) (*int, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOInt2int(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOInt2ᚖint(ctx context.Context, sel ast.SelectionSet, v *int) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOInt2int(ctx, sel, *v)
}

func (ec *executionContext) unmarshalOLengthUnit2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐLengthUnit(ctx context.Context, v interface{}) (LengthUnit, error) {
	var res LengthUnit
	return res, res.UnmarshalGQL(v)
}

func (ec *executionContext) marshalOLengthUnit2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐLengthUnit(ctx context.Context, sel ast.SelectionSet, v LengthUnit) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalOLengthUnit2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐLengthUnit(ctx context.Context, v interface{}) (*LengthUnit, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOLengthUnit2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐLengthUnit(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOLengthUnit2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐLengthUnit(ctx context.Context, sel ast.SelectionSet, v *LengthUnit) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return v
}

func (ec *executionContext) marshalOReview2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐReview(ctx context.Context, sel ast.SelectionSet, v Review) graphql.Marshaler {
	return ec._Review(ctx, sel, &v)
}

func (ec *executionContext) marshalOReview2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐReview(ctx context.Context, sel ast.SelectionSet, v *Review) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Review(ctx, sel, v)
}

func (ec *executionContext) marshalOStarship2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐStarship(ctx context.Context, sel ast.SelectionSet, v Starship) graphql.Marshaler {
	return ec._Starship(ctx, sel, &v)
}

func (ec *executionContext) marshalOStarship2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐStarship(ctx context.Context, sel ast.SelectionSet, v []Starship) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNStarship2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐStarship(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalOStarship2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋstarwarsᚐStarship(ctx context.Context, sel ast.SelectionSet, v *Starship) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Starship(ctx, sel, v)
}

func (ec *executionContext) unmarshalOString2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalOString2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	return graphql.MarshalString(v)
}

func (ec *executionContext) unmarshalOString2ᚖstring(ctx context.Context, v interface{}) (*string, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOString2string(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOString2ᚖstring(ctx context.Context, sel ast.SelectionSet, v *string) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOString2string(ctx, sel, *v)
}

func (ec *executionContext) unmarshalOTime2timeᚐTime(ctx context.Context, v interface{}) (time.Time, error) {
	return graphql.UnmarshalTime(v)
}

func (ec *executionContext) marshalOTime2timeᚐTime(ctx context.Context, sel ast.SelectionSet, v time.Time) graphql.Marshaler {
	return graphql.MarshalTime(v)
}

func (ec *executionContext) unmarshalOTime2ᚖtimeᚐTime(ctx context.Context, v interface{}) (*time.Time, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOTime2timeᚐTime(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOTime2ᚖtimeᚐTime(ctx context.Context, sel ast.SelectionSet, v *time.Time) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOTime2timeᚐTime(ctx, sel, *v)
}

func (ec *executionContext) marshalO__EnumValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx context.Context, sel ast.SelectionSet, v []introspection.EnumValue) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__EnumValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__Field2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx context.Context, sel ast.SelectionSet, v []introspection.Field) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Field2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx context.Context, sel ast.SelectionSet, v []introspection.InputValue) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__Schema2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx context.Context, sel ast.SelectionSet, v introspection.Schema) graphql.Marshaler {
	return ec.___Schema(ctx, sel, &v)
}

func (ec *executionContext) marshalO__Schema2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx context.Context, sel ast.SelectionSet, v *introspection.Schema) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.___Schema(ctx, sel, v)
}

func (ec *executionContext) marshalO__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v introspection.Type) graphql.Marshaler {
	return ec.___Type(ctx, sel, &v)
}

func (ec *executionContext) marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v []introspection.Type) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		rctx := &graphql.ResolverContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(i int) {
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v *introspection.Type) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
