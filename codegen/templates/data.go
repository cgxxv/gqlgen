package templates

var data = map[string]string{
	"args.gotpl":      "\t{{- if . }}args := map[string]interface{}{} {{end}}\n\t{{- range $i, $arg := . }}\n\t\tvar arg{{$i}} {{$arg.Signature }}\n\t\tif tmp, ok := field.Args[{{$arg.GQLName|quote}}]; ok {\n\t\t\tvar err error\n\t\t\t{{$arg.Unmarshal (print \"arg\" $i) \"tmp\" }}\n\t\t\tif err != nil {\n\t\t\t\tec.Error(ctx, err)\n\t\t\t\t{{- if $arg.Object.Stream }}\n\t\t\t\t\treturn nil\n\t\t\t\t{{- else }}\n\t\t\t\t\treturn graphql.Null\n\t\t\t\t{{- end }}\n\t\t\t}\n\t\t} {{ if $arg.Default }} else {\n\t\t\tvar tmp interface{} = {{ $arg.Default | dump }}\n\t\t\tvar err error\n\t\t\t{{$arg.Unmarshal (print \"arg\" $i) \"tmp\" }}\n\t\t\tif err != nil {\n\t\t\t\tec.Error(ctx, err)\n\t\t\t\t{{- if $arg.Object.Stream }}\n\t\t\t\t\treturn nil\n\t\t\t\t{{- else }}\n\t\t\t\t\treturn graphql.Null\n\t\t\t\t{{- end }}\n\t\t\t}\n\t\t}\n\t\t{{end }}\n\t\targs[{{$arg.GQLName|quote}}] = arg{{$i}}\n\t{{- end -}}\n",
	"field.gotpl":     "{{ $field := . }}\n{{ $object := $field.Object }}\n\n{{- if $object.Stream }}\n\tfunc (ec *executionContext) _{{$object.GQLType}}_{{$field.GQLName}}(ctx context.Context, field graphql.CollectedField) func() graphql.Marshaler {\n\t\t{{- template \"args.gotpl\" $field.Args }}\n\t\tctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{Field: field})\n\t\tresults, err := ec.resolvers.{{ $field.ShortInvocation }}\n\t\tif err != nil {\n\t\t\tec.Error(ctx, err)\n\t\t\treturn nil\n\t\t}\n\t\treturn func() graphql.Marshaler {\n\t\t\tres, ok := <-results\n\t\t\tif !ok {\n\t\t\t\treturn nil\n\t\t\t}\n\t\t\tvar out graphql.OrderedMap\n\t\t\tout.Add(field.Alias, func() graphql.Marshaler { {{ $field.WriteJson }} }())\n\t\t\treturn &out\n\t\t}\n\t}\n{{ else }}\n\tfunc (ec *executionContext) _{{$object.GQLType}}_{{$field.GQLName}}(ctx context.Context, field graphql.CollectedField, {{if not $object.Root}}obj *{{$object.FullName}}{{end}}) graphql.Marshaler {\n\t\t{{- template \"args.gotpl\" $field.Args }}\n\n\t\t{{- if $field.IsConcurrent }}\n\t\t\tctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{\n\t\t\t\tObject: {{$object.GQLType|quote}},\n\t\t\t\tArgs: {{if $field.Args }}args{{else}}nil{{end}},\n\t\t\t\tField: field,\n\t\t\t})\n\t\t\treturn graphql.Defer(func() (ret graphql.Marshaler) {\n\t\t\t\tdefer func() {\n\t\t\t\t\tif r := recover(); r != nil {\n\t\t\t\t\t\tuserErr := ec.Recover(ctx, r)\n\t\t\t\t\t\tec.Error(ctx, userErr)\n\t\t\t\t\t\tret = graphql.Null\n\t\t\t\t\t}\n\t\t\t\t}()\n\t\t{{ else }}\n\t\t\trctx := graphql.GetResolverContext(ctx)\n\t\t\trctx.Object = {{$object.GQLType|quote}}\n\t\t\trctx.Args = {{if $field.Args }}args{{else}}nil{{end}}\n\t\t\trctx.Field = field\n\t\t\trctx.PushField(field.Alias)\n\t\t\tdefer rctx.Pop()\n\t\t{{- end }}\n\t\t\tresTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {\n\t\t\t\t{{- if $field.IsResolver }}\n\t\t\t\t\treturn ec.resolvers.{{ $field.ShortInvocation }}\n\t\t\t\t{{- else if $field.GoVarName }}\n\t\t\t\t\treturn obj.{{$field.GoVarName}}, nil\n\t\t\t\t{{- else if $field.GoMethodName }}\n\t\t\t\t\t{{- if $field.NoErr }}\n\t\t\t\t\t\treturn {{$field.GoMethodName}}({{ $field.CallArgs }}), nil\n\t\t\t\t\t{{- else }}\n\t\t\t\t\t\treturn {{$field.GoMethodName}}({{ $field.CallArgs }})\n\t\t\t\t\t{{- end }}\n\t\t\t\t{{- end }}\n\t\t\t})\n\t\t\tif resTmp == nil {\n\t\t\t\treturn graphql.Null\n\t\t\t}\n\t\t\tres := resTmp.({{$field.Signature}})\n\t\t\t{{ $field.WriteJson }}\n\t\t{{- if $field.IsConcurrent }}\n\t\t\t})\n\t\t{{- end }}\n\t}\n{{ end }}\n",
	"generated.gotpl": "// Code generated by github.com/vektah/gqlgen, DO NOT EDIT.\n\npackage {{ .PackageName }}\n\nimport (\n{{- range $import := .Imports }}\n\t{{- $import.Write }}\n{{ end }}\n)\n\n// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.\nfunc NewExecutableSchema(resolvers ResolverRoot) graphql.ExecutableSchema {\n\treturn &executableSchema{resolvers: resolvers}\n}\n\ntype Config struct {\n\tresolvers ResolverRoot\n\t{{ if .Directives }}directive DirectiveRoot{{ end }}\n}\n\ntype ResolverRoot interface {\n{{- range $object := .Objects -}}\n\t{{ if $object.HasResolvers -}}\n\t\t{{$object.GQLType}}() {{$object.GQLType}}Resolver\n\t{{ end }}\n{{- end }}\n}\n\n{{ if .Directives }}\ntype DirectiveRoot struct {\n{{ range $directive := .Directives }}\n\t{{$directive.Name}} graphql.FieldMiddleware\n{{ end }}\n}\n{{ end }}\n\n{{- range $object := .Objects -}}\n\t{{ if $object.HasResolvers }}\n\t\ttype {{$object.GQLType}}Resolver interface {\n\t\t{{ range $field := $object.Fields -}}\n\t\t\t{{ $field.ShortResolverDeclaration }}\n\t\t{{ end }}\n\t\t}\n\t{{- end }}\n{{- end }}\n\ntype executableSchema struct {\n\tresolvers      ResolverRoot\n}\n\nfunc (e *executableSchema) Schema() *ast.Schema {\n\treturn parsedSchema\n}\n\nfunc (e *executableSchema) Query(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {\n\t{{- if .QueryRoot }}\n\t\tec := executionContext{graphql.GetRequestContext(ctx), e.resolvers}\n\n\t\tbuf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {\n\t\t\tdata := ec._{{.QueryRoot.GQLType}}(ctx, op.SelectionSet)\n\t\t\tvar buf bytes.Buffer\n\t\t\tdata.MarshalGQL(&buf)\n\t\t\treturn buf.Bytes()\n\t\t})\n\n\t\treturn &graphql.Response{\n\t\t\tData:   buf,\n\t\t\tErrors: ec.Errors,\n\t\t}\n\t{{- else }}\n\t\treturn graphql.ErrorResponse(ctx, \"queries are not supported\")\n\t{{- end }}\n}\n\nfunc (e *executableSchema) Mutation(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {\n\t{{- if .MutationRoot }}\n\t\tec := executionContext{graphql.GetRequestContext(ctx), e.resolvers}\n\n\t\tbuf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {\n\t\t\tdata := ec._{{.MutationRoot.GQLType}}(ctx, op.SelectionSet)\n\t\t\tvar buf bytes.Buffer\n\t\t\tdata.MarshalGQL(&buf)\n\t\t\treturn buf.Bytes()\n\t\t})\n\n\t\treturn &graphql.Response{\n\t\t\tData:   buf,\n\t\t\tErrors: ec.Errors,\n\t\t}\n\t{{- else }}\n\t\treturn graphql.ErrorResponse(ctx, \"mutations are not supported\")\n\t{{- end }}\n}\n\nfunc (e *executableSchema) Subscription(ctx context.Context, op *ast.OperationDefinition) func() *graphql.Response {\n\t{{- if .SubscriptionRoot }}\n\t\tec := executionContext{graphql.GetRequestContext(ctx), e.resolvers}\n\n\t\tnext := ec._{{.SubscriptionRoot.GQLType}}(ctx, op.SelectionSet)\n\t\tif ec.Errors != nil {\n\t\t\treturn graphql.OneShot(&graphql.Response{Data: []byte(\"null\"), Errors: ec.Errors})\n\t\t}\n\n\t\tvar buf bytes.Buffer\n\t\treturn func() *graphql.Response {\n\t\t\tbuf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {\n\t\t\t\tbuf.Reset()\n\t\t\t\tdata := next()\n\n\t\t\t\tif data == nil {\n\t\t\t\t\treturn nil\n\t\t\t\t}\n\t\t\t\tdata.MarshalGQL(&buf)\n\t\t\t\treturn buf.Bytes()\n\t\t\t})\n\n\t\t\treturn &graphql.Response{\n\t\t\t\tData:   buf,\n\t\t\t\tErrors: ec.Errors,\n\t\t\t}\n\t\t}\n\t{{- else }}\n\t\treturn graphql.OneShot(graphql.ErrorResponse(ctx, \"subscriptions are not supported\"))\n\t{{- end }}\n}\n\ntype executionContext struct {\n\t*graphql.RequestContext\n\n\tresolvers ResolverRoot\n}\n\n{{- range $object := .Objects }}\n\t{{ template \"object.gotpl\" $object }}\n\n\t{{- range $field := $object.Fields }}\n\t\t{{ template \"field.gotpl\" $field }}\n\t{{ end }}\n{{- end}}\n\n{{- range $interface := .Interfaces }}\n\t{{ template \"interface.gotpl\" $interface }}\n{{- end }}\n\n{{- range $input := .Inputs }}\n\t{{ template \"input.gotpl\" $input }}\n{{- end }}\n\nfunc (ec *executionContext) introspectSchema() *introspection.Schema {\n\treturn introspection.WrapSchema(parsedSchema)\n}\n\nfunc (ec *executionContext) introspectType(name string) *introspection.Type {\n\treturn introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name])\n}\n\nvar parsedSchema = gqlparser.MustLoadSchema(\n\t&ast.Source{Name: {{.SchemaFilename|quote}}, Input: {{.SchemaRaw|rawQuote}}},\n)\n",
	"input.gotpl":     "\t{{- if .IsMarshaled }}\n\tfunc Unmarshal{{ .GQLType }}(v interface{}) ({{.FullName}}, error) {\n\t\tvar it {{.FullName}}\n\t\tvar asMap = v.(map[string]interface{})\n\t\t{{ range $field := .Fields}}\n\t\t\t{{- if $field.Default}}\n\t\t\t\tif _, present := asMap[{{$field.GQLName|quote}}] ; !present {\n\t\t\t\t\tasMap[{{$field.GQLName|quote}}] = {{ $field.Default | dump }}\n\t\t\t\t}\n\t\t\t{{- end}}\n\t\t{{- end }}\n\n\t\tfor k, v := range asMap {\n\t\t\tswitch k {\n\t\t\t{{- range $field := .Fields }}\n\t\t\tcase {{$field.GQLName|quote}}:\n\t\t\t\tvar err error\n\t\t\t\t{{ $field.Unmarshal (print \"it.\" $field.GoVarName) \"v\" }}\n\t\t\t\tif err != nil {\n\t\t\t\t\treturn it, err\n\t\t\t\t}\n\t\t\t{{- end }}\n\t\t\t}\n\t\t}\n\n\t\treturn it, nil\n\t}\n\t{{- end }}\n",
	"interface.gotpl": "{{- $interface := . }}\n\nfunc (ec *executionContext) _{{$interface.GQLType}}(ctx context.Context, sel ast.SelectionSet, obj *{{$interface.FullName}}) graphql.Marshaler {\n\tswitch obj := (*obj).(type) {\n\tcase nil:\n\t\treturn graphql.Null\n\t{{- range $implementor := $interface.Implementors }}\n\t\t{{- if $implementor.ValueReceiver }}\n\t\t\tcase {{$implementor.FullName}}:\n\t\t\t\treturn ec._{{$implementor.GQLType}}(ctx, sel, &obj)\n\t\t{{- end}}\n\t\tcase *{{$implementor.FullName}}:\n\t\t\treturn ec._{{$implementor.GQLType}}(ctx, sel, obj)\n\t{{- end }}\n\tdefault:\n\t\tpanic(fmt.Errorf(\"unexpected type %T\", obj))\n\t}\n}\n",
	"models.gotpl":    "// Code generated by github.com/vektah/gqlgen, DO NOT EDIT.\n\npackage {{ .PackageName }}\n\nimport (\n{{- range $import := .Imports }}\n\t{{- $import.Write }}\n{{ end }}\n)\n\n{{ range $model := .Models }}\n\t{{- if .IsInterface }}\n\t\ttype {{.GoType}} interface {}\n\t{{- else }}\n\t\ttype {{.GoType}} struct {\n\t\t\t{{- range $field := .Fields }}\n\t\t\t\t{{- if $field.GoVarName }}\n\t\t\t\t\t{{ $field.GoVarName }} {{$field.Signature}} `json:\"{{$field.GQLName}}\"`\n\t\t\t\t{{- else }}\n\t\t\t\t\t{{ $field.GoFKName }} {{$field.GoFKType}}\n\t\t\t\t{{- end }}\n\t\t\t{{- end }}\n\t\t}\n\t{{- end }}\n{{- end}}\n\n{{ range $enum := .Enums }}\n\ttype {{.GoType}} string\n\tconst (\n\t{{ range $value := .Values -}}\n\t\t{{with .Description}} {{.|prefixLines \"// \"}} {{end}}\n\t\t{{$enum.GoType}}{{ .Name|toCamel }} {{$enum.GoType}} = {{.Name|quote}}\n\t{{- end }}\n\t)\n\n\tfunc (e {{.GoType}}) IsValid() bool {\n\t\tswitch e {\n\t\tcase {{ range $index, $element := .Values}}{{if $index}},{{end}}{{ $enum.GoType }}{{ $element.Name|toCamel }}{{end}}:\n\t\t\treturn true\n\t\t}\n\t\treturn false\n\t}\n\n\tfunc (e {{.GoType}}) String() string {\n\t\treturn string(e)\n\t}\n\n\tfunc (e *{{.GoType}}) UnmarshalGQL(v interface{}) error {\n\t\tstr, ok := v.(string)\n\t\tif !ok {\n\t\t\treturn fmt.Errorf(\"enums must be strings\")\n\t\t}\n\n\t\t*e = {{.GoType}}(str)\n\t\tif !e.IsValid() {\n\t\t\treturn fmt.Errorf(\"%s is not a valid {{.GQLType}}\", str)\n\t\t}\n\t\treturn nil\n\t}\n\n\tfunc (e {{.GoType}}) MarshalGQL(w io.Writer) {\n\t\tfmt.Fprint(w, strconv.Quote(e.String()))\n\t}\n\n{{- end }}\n",
	"object.gotpl":    "{{ $object := . }}\n\nvar {{ $object.GQLType|lcFirst}}Implementors = {{$object.Implementors}}\n\n// nolint: gocyclo, errcheck, gas, goconst\n{{- if .Stream }}\nfunc (ec *executionContext) _{{$object.GQLType}}(ctx context.Context, sel ast.SelectionSet) func() graphql.Marshaler {\n\tfields := graphql.CollectFields(ctx, sel, {{$object.GQLType|lcFirst}}Implementors)\n\tctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{\n\t\tObject: {{$object.GQLType|quote}},\n\t})\n\tif len(fields) != 1 {\n\t\tec.Errorf(ctx, \"must subscribe to exactly one stream\")\n\t\treturn nil\n\t}\n\n\tswitch fields[0].Name {\n\t{{- range $field := $object.Fields }}\n\tcase \"{{$field.GQLName}}\":\n\t\treturn ec._{{$object.GQLType}}_{{$field.GQLName}}(ctx, fields[0])\n\t{{- end }}\n\tdefault:\n\t\tpanic(\"unknown field \" + strconv.Quote(fields[0].Name))\n\t}\n}\n{{- else }}\nfunc (ec *executionContext) _{{$object.GQLType}}(ctx context.Context, sel ast.SelectionSet{{if not $object.Root}}, obj *{{$object.FullName}} {{end}}) graphql.Marshaler {\n\tfields := graphql.CollectFields(ctx, sel, {{$object.GQLType|lcFirst}}Implementors)\n\t{{if $object.Root}}\n\t\tctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{\n\t\t\tObject: {{$object.GQLType|quote}},\n\t\t})\n\t{{end}}\n\tout := graphql.NewOrderedMap(len(fields))\n\tfor i, field := range fields {\n\t\tout.Keys[i] = field.Alias\n\n\t\tswitch field.Name {\n\t\tcase \"__typename\":\n\t\t\tout.Values[i] = graphql.MarshalString({{$object.GQLType|quote}})\n\t\t{{- range $field := $object.Fields }}\n\t\tcase \"{{$field.GQLName}}\":\n\t\t\tout.Values[i] = ec._{{$object.GQLType}}_{{$field.GQLName}}(ctx, field{{if not $object.Root}}, obj{{end}})\n\t\t{{- end }}\n\t\tdefault:\n\t\t\tpanic(\"unknown field \" + strconv.Quote(field.Name))\n\t\t}\n\t}\n\n\treturn out\n}\n{{- end }}\n",
}
