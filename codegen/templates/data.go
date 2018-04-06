package templates

var data = map[string]string{
	"args.gotpl":      "\t{{- if . }}args := map[string]interface{}{} {{end}}\n\t{{- range $i, $arg := . }}\n\t\tvar arg{{$i}} {{$arg.Signature }}\n\t\tif tmp, ok := field.Args[{{$arg.GQLName|quote}}]; ok {\n\t\t\tvar err error\n\t\t\t{{$arg.Unmarshal (print \"arg\" $i) \"tmp\" }}\n\t\t\tif err != nil {\n\t\t\t\tec.Error(err)\n\t\t\t\t{{- if $arg.Object.Stream }}\n\t\t\t\t\treturn nil\n\t\t\t\t{{- else }}\n\t\t\t\t\treturn graphql.Null\n\t\t\t\t{{- end }}\n\t\t\t}\n\t\t} {{ if $arg.Default }} else {\n\t\t\tvar tmp interface{} = {{ $arg.Default | dump }}\n\t\t\tvar err error\n\t\t\t{{$arg.Unmarshal (print \"arg\" $i) \"tmp\" }}\n\t\t\tif err != nil {\n\t\t\t\tec.Error(err)\n\t\t\t\t{{- if $arg.Object.Stream }}\n\t\t\t\t\treturn nil\n\t\t\t\t{{- else }}\n\t\t\t\t\treturn graphql.Null\n\t\t\t\t{{- end }}\n\t\t\t}\n\t\t}\n\t\t{{end }}\n\t\targs[{{$arg.GQLName|quote}}] = arg{{$i}}\n\t{{- end -}}\n",
	"field.gotpl":     "{{ $field := . }}\n{{ $object := $field.Object }}\n\n{{- if $object.Stream }}\n\tfunc (ec *executionContext) _{{$object.GQLType}}_{{$field.GQLName}}(ctx context.Context, field graphql.CollectedField) func() graphql.Marshaler {\n\t\t{{- template \"args.gotpl\" $field.Args }}\n\t\trctx := graphql.WithResolverContext(ctx, &graphql.ResolverContext{Field: field})\n\t\tresults, err := ec.resolvers.{{ $object.GQLType }}_{{ $field.GQLName }}({{ $field.CallArgs }})\n\t\tif err != nil {\n\t\t\tec.Error(err)\n\t\t\treturn nil\n\t\t}\n\t\treturn func() graphql.Marshaler {\n\t\t\tres, ok := <-results\n\t\t\tif !ok {\n\t\t\t\treturn nil\n\t\t\t}\n\t\t\tvar out graphql.OrderedMap\n\t\t\tout.Add(field.Alias, func() graphql.Marshaler { {{ $field.WriteJson }} }())\n\t\t\treturn &out\n\t\t}\n\t}\n{{ else }}\n\tfunc (ec *executionContext) _{{$object.GQLType}}_{{$field.GQLName}}(ctx context.Context, field graphql.CollectedField, {{if not $object.Root}}obj *{{$object.FullName}}{{end}}) graphql.Marshaler {\n\t\t{{- template \"args.gotpl\" $field.Args }}\n\n\t\t{{- if $field.IsConcurrent }}\n\t\t\treturn graphql.Defer(func() (ret graphql.Marshaler) {\n\t\t\t\tdefer func() {\n\t\t\t\t\tif r := recover(); r != nil {\n\t\t\t\t\t\tuserErr := ec.Recover(r)\n\t\t\t\t\t\tec.Error(userErr)\n\t\t\t\t\t\tret = graphql.Null\n\t\t\t\t\t}\n\t\t\t\t}()\n\t\t{{- end }}\n\n\t\t\t{{- if $field.GoVarName }}\n\t\t\t\tres := obj.{{$field.GoVarName}}\n\t\t\t{{- else if $field.GoMethodName }}\n\t\t\t\t{{- if $field.NoErr }}\n\t\t\t\t\tres := {{$field.GoMethodName}}({{ $field.CallArgs }})\n\t\t\t\t{{- else }}\n\t\t\t\t\tres, err := {{$field.GoMethodName}}({{ $field.CallArgs }})\n\t\t\t\t\tif err != nil {\n\t\t\t\t\t\tec.Error(err)\n\t\t\t\t\t\treturn graphql.Null\n\t\t\t\t\t}\n\t\t\t\t{{- end }}\n\t\t\t{{- else }}\n\t\t\t\trctx := graphql.WithResolverContext(ctx, &graphql.ResolverContext{\n\t\t\t\t\tObject: {{$object.GQLType|quote}},\n\t\t\t\t\tArgs: {{if $field.Args }}args{{else}}nil{{end}},\n\t\t\t\t\tField: field,\n\t\t\t\t})\n\t\t\t\tresTmp, err := ec.Middleware(rctx, func(rctx context.Context) (interface{}, error) {\n\t\t\t\t\treturn ec.resolvers.{{ $object.GQLType }}_{{ $field.GQLName }}({{ $field.CallArgs }})\n\t\t\t\t})\n\t\t\t\tif err != nil {\n\t\t\t\t\tec.Error(err)\n\t\t\t\t\treturn graphql.Null\n\t\t\t\t}\n\t\t\t\tif resTmp == nil {\n\t\t\t\t\treturn graphql.Null\n\t\t\t\t}\n\t\t\t\tres := resTmp.({{$field.Signature}})\n\t\t\t{{- end }}\n\t\t\t{{ $field.WriteJson }}\n\t\t{{- if $field.IsConcurrent }}\n\t\t\t})\n\t\t{{- end }}\n\t}\n{{ end }}\n",
	"generated.gotpl": "// This file was generated by github.com/vektah/gqlgen, DO NOT EDIT\n\npackage {{ .PackageName }}\n\nimport (\n{{- range $import := .Imports }}\n\t{{- $import.Write }}\n{{ end }}\n)\n\nfunc MakeExecutableSchema(resolvers Resolvers) graphql.ExecutableSchema {\n\treturn &executableSchema{resolvers: resolvers}\n}\n\ntype Resolvers interface {\n{{- range $object := .Objects -}}\n\t{{ range $field := $object.Fields -}}\n\t\t{{ $field.ResolverDeclaration }}\n\t{{ end }}\n{{- end }}\n}\n\ntype executableSchema struct {\n\tresolvers      Resolvers\n}\n\nfunc (e *executableSchema) Schema() *schema.Schema {\n\treturn parsedSchema\n}\n\nfunc (e *executableSchema) Query(ctx context.Context, op *query.Operation) *graphql.Response {\n\t{{- if .QueryRoot }}\n\t\tec := executionContext{graphql.GetRequestContext(ctx), e.resolvers}\n\n\t\tdata := ec._{{.QueryRoot.GQLType}}(ctx, op.Selections)\n\t\tvar buf bytes.Buffer\n\t\tdata.MarshalGQL(&buf)\n\n\t\treturn &graphql.Response{\n\t\t\tData:   buf.Bytes(),\n\t\t\tErrors: ec.Errors,\n\t\t}\n\t{{- else }}\n\t\treturn &graphql.Response{Errors: []*errors.QueryError{ {Message: \"queries are not supported\"} }}\n\t{{- end }}\n}\n\nfunc (e *executableSchema) Mutation(ctx context.Context, op *query.Operation) *graphql.Response {\n\t{{- if .MutationRoot }}\n\t\tec := executionContext{graphql.GetRequestContext(ctx), e.resolvers}\n\n\t\tdata := ec._{{.MutationRoot.GQLType}}(ctx, op.Selections)\n\t\tvar buf bytes.Buffer\n\t\tdata.MarshalGQL(&buf)\n\n\t\treturn &graphql.Response{\n\t\t\tData:   buf.Bytes(),\n\t\t\tErrors: ec.Errors,\n\t\t}\n\t{{- else }}\n\t\treturn &graphql.Response{Errors: []*errors.QueryError{ {Message: \"mutations are not supported\"} }}\n\t{{- end }}\n}\n\nfunc (e *executableSchema) Subscription(ctx context.Context, op *query.Operation) func() *graphql.Response {\n\t{{- if .SubscriptionRoot }}\n\t\tec := executionContext{graphql.GetRequestContext(ctx), e.resolvers}\n\n\t\tnext := ec._{{.SubscriptionRoot.GQLType}}(ctx, op.Selections)\n\t\tif ec.Errors != nil {\n\t\t\treturn graphql.OneShot(&graphql.Response{Data: []byte(\"null\"), Errors: ec.Errors})\n\t\t}\n\n\t\tvar buf bytes.Buffer\n\t\treturn func() *graphql.Response {\n\t\t\tbuf.Reset()\n\t\t\tdata := next()\n\t\t\tif data == nil {\n\t\t\t\treturn nil\n\t\t\t}\n\t\t\tdata.MarshalGQL(&buf)\n\n\t\t\terrs := ec.Errors\n\t\t\tec.Errors = nil\n\t\t\treturn &graphql.Response{\n\t\t\t\tData:   buf.Bytes(),\n\t\t\t\tErrors: errs,\n\t\t\t}\n\t\t}\n\t{{- else }}\n\t\treturn graphql.OneShot(&graphql.Response{Errors: []*errors.QueryError{ {Message: \"subscriptions are not supported\"} }})\n\t{{- end }}\n}\n\ntype executionContext struct {\n\t*graphql.RequestContext\n\n\tresolvers Resolvers\n}\n\n{{- range $object := .Objects }}\n\t{{ template \"object.gotpl\" $object }}\n\n\t{{- range $field := $object.Fields }}\n\t\t{{ template \"field.gotpl\" $field }}\n\t{{ end }}\n{{- end}}\n\n{{- range $interface := .Interfaces }}\n\t{{ template \"interface.gotpl\" $interface }}\n{{- end }}\n\n{{- range $input := .Inputs }}\n\t{{ template \"input.gotpl\" $input }}\n{{- end }}\n\nfunc (ec *executionContext) introspectSchema() *introspection.Schema {\n\treturn introspection.WrapSchema(parsedSchema)\n}\n\nfunc (ec *executionContext) introspectType(name string) *introspection.Type {\n\tt := parsedSchema.Resolve(name)\n\tif t == nil {\n\t\treturn nil\n\t}\n\treturn introspection.WrapType(t)\n}\n\nvar parsedSchema = schema.MustParse({{.SchemaRaw|rawQuote}})\n",
	"input.gotpl":     "\t{{- if .IsMarshaled }}\n\tfunc Unmarshal{{ .GQLType }}(v interface{}) ({{.FullName}}, error) {\n\t\tvar it {{.FullName}}\n\n\t\tfor k, v := range v.(map[string]interface{}) {\n\t\t\tswitch k {\n\t\t\t{{- range $field := .Fields }}\n\t\t\tcase {{$field.GQLName|quote}}:\n\t\t\t\tvar err error\n\t\t\t\t{{ $field.Unmarshal (print \"it.\" $field.GoVarName) \"v\" }}\n\t\t\t\tif err != nil {\n\t\t\t\t\treturn it, err\n\t\t\t\t}\n\t\t\t{{- end }}\n\t\t\t}\n\t\t}\n\n\t\treturn it, nil\n\t}\n\t{{- end }}\n",
	"interface.gotpl": "{{- $interface := . }}\n\nfunc (ec *executionContext) _{{$interface.GQLType}}(ctx context.Context, sel []query.Selection, obj *{{$interface.FullName}}) graphql.Marshaler {\n\tswitch obj := (*obj).(type) {\n\tcase nil:\n\t\treturn graphql.Null\n\t{{- range $implementor := $interface.Implementors }}\n\t\t{{- if $implementor.ValueReceiver }}\n\t\t\tcase {{$implementor.FullName}}:\n\t\t\t\treturn ec._{{$implementor.GQLType}}(ctx, sel, &obj)\n\t\t{{- end}}\n\t\tcase *{{$implementor.FullName}}:\n\t\t\treturn ec._{{$implementor.GQLType}}(ctx, sel, obj)\n\t{{- end }}\n\tdefault:\n\t\tpanic(fmt.Errorf(\"unexpected type %T\", obj))\n\t}\n}\n",
	"models.gotpl":    "// This file was generated by github.com/vektah/gqlgen, DO NOT EDIT\n\npackage {{ .PackageName }}\n\nimport (\n{{- range $import := .Imports }}\n\t{{- $import.Write }}\n{{ end }}\n)\n\n{{ range $model := .Models }}\n\t{{- if .IsInterface }}\n\t\ttype {{.GoType}} interface {}\n\t{{- else }}\n\t\ttype {{.GoType}} struct {\n\t\t\t{{- range $field := .Fields }}\n\t\t\t\t{{- if $field.GoVarName }}\n\t\t\t\t\t{{ $field.GoVarName }} {{$field.Signature}}\n\t\t\t\t{{- else }}\n\t\t\t\t\t{{ $field.GoFKName }} {{$field.GoFKType}}\n\t\t\t\t{{- end }}\n\t\t\t{{- end }}\n\t\t}\n\t{{- end }}\n{{- end}}\n\n{{ range $enum := .Enums }}\n\ttype {{.GoType}} string\n\tconst (\n\t{{ range $value := .Values }}\n\t\t{{$enum.GoType}}{{ .Name|toCamel }} {{$enum.GoType}} = {{.Name|quote}} {{with .Description}} // {{.}} {{end}}\n\t{{- end }}\n\t)\n\n\tfunc (e {{.GoType}}) IsValid() bool {\n\t\tswitch e {\n\t\tcase {{ range $index, $element := .Values}}{{if $index}},{{end}}{{ $enum.GoType }}{{ $element.Name|toCamel }}{{end}}:\n\t\t\treturn true\n\t\t}\n\t\treturn false\n\t}\n\n\tfunc (e {{.GoType}}) String() string {\n\t\treturn string(e)\n\t}\n\n\tfunc (e *{{.GoType}}) UnmarshalGQL(v interface{}) error {\n\t\tstr, ok := v.(string)\n\t\tif !ok {\n\t\t\treturn fmt.Errorf(\"enums must be strings\")\n\t\t}\n\n\t\t*e = {{.GoType}}(str)\n\t\tif !e.IsValid() {\n\t\t\treturn fmt.Errorf(\"%s is not a valid {{.GQLType}}\", str)\n\t\t}\n\t\treturn nil\n\t}\n\n\tfunc (e {{.GoType}}) MarshalGQL(w io.Writer) {\n\t\tfmt.Fprint(w, strconv.Quote(e.String()))\n\t}\n\n{{- end }}\n",
	"object.gotpl":    "{{ $object := . }}\n\nvar {{ $object.GQLType|lcFirst}}Implementors = {{$object.Implementors}}\n\n// nolint: gocyclo, errcheck, gas, goconst\n{{- if .Stream }}\nfunc (ec *executionContext) _{{$object.GQLType}}(ctx context.Context, sel []query.Selection) func() graphql.Marshaler {\n\tfields := graphql.CollectFields(ec.Doc, sel, {{$object.GQLType|lcFirst}}Implementors, ec.Variables)\n\n\tif len(fields) != 1 {\n\t\tec.Errorf(\"must subscribe to exactly one stream\")\n\t\treturn nil\n\t}\n\n\tswitch fields[0].Name {\n\t{{- range $field := $object.Fields }}\n\tcase \"{{$field.GQLName}}\":\n\t\treturn ec._{{$object.GQLType}}_{{$field.GQLName}}(ctx, fields[0])\n\t{{- end }}\n\tdefault:\n\t\tpanic(\"unknown field \" + strconv.Quote(fields[0].Name))\n\t}\n}\n{{- else }}\nfunc (ec *executionContext) _{{$object.GQLType}}(ctx context.Context, sel []query.Selection{{if not $object.Root}}, obj *{{$object.FullName}} {{end}}) graphql.Marshaler {\n\tfields := graphql.CollectFields(ec.Doc, sel, {{$object.GQLType|lcFirst}}Implementors, ec.Variables)\n\tout := graphql.NewOrderedMap(len(fields))\n\tfor i, field := range fields {\n\t\tout.Keys[i] = field.Alias\n\n\t\tswitch field.Name {\n\t\tcase \"__typename\":\n\t\t\tout.Values[i] = graphql.MarshalString({{$object.GQLType|quote}})\n\t\t{{- range $field := $object.Fields }}\n\t\tcase \"{{$field.GQLName}}\":\n\t\t\tout.Values[i] = ec._{{$object.GQLType}}_{{$field.GQLName}}(ctx, field{{if not $object.Root}}, obj{{end}})\n\t\t{{- end }}\n\t\tdefault:\n\t\t\tpanic(\"unknown field \" + strconv.Quote(field.Name))\n\t\t}\n\t}\n\n\treturn out\n}\n{{- end }}\n",
}
