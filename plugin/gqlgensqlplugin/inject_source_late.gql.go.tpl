{{- range $key, $value := .Refs.ReferObjects}}

input {{$key}}Ref {
	{{- range $fieldKey, $field := $value.InputRefFields}}
  {{$field.Name}}: {{$field.RefGqlType}}
	{{- end}}
}

{{- end}}

{{- range $key, $value := .List}}
enum {{$value.TypeName}}HasFilter {
	{{- range $fieldKey, $field := $value.Fields}}
  {{$field.Name}}
	{{- end}}
}

enum {{$value.TypeName}}Orderable {
	{{- range $fieldKey, $field := $value.OrderAbleFields}}
  {{$field.Name}}
	{{- end}}
}

input {{$value.TypeName}}Filter {
	{{$value.PrimaryField.Name}}: [{{$value.PrimaryField.GqlType}}!]
	has: [{{$value.TypeName}}HasFilter]
	and: [{{$value.TypeName}}Filter]
	or: [{{$value.TypeName}}Filter]
	not: [{{$value.TypeName}}Filter]
}

input {{$value.TypeName}}Order {
	asc: {{$value.TypeName}}Orderable
	desc: {{$value.TypeName}}Orderable
	then: {{$value.TypeName}}Order
}

type {{$value.TypeName}}AggregateResult {
	count: Int
	{{- range $fieldKey, $field := $value.AggregateFields}}
  {{$field.Name}}Min: {{$field.GqlType}}
	{{$field.Name}}Max: {{$field.GqlType}}
	{{- end}}
}



type Add{{$value.TypeName}}Payload {
	{{$value.TypeName}}(filter: {{$value.TypeName}}Filter, order: {{$value.TypeName}}Order, first: Int, offset: Int): [{{$value.TypeName}}]
}

type Update{{$value.TypeName}}Payload {
	{{$value.TypeName}}(filter: {{$value.TypeName}}Filter, order: {{$value.TypeName}}Order, first: Int, offset: Int): [{{$value.TypeName}}]
	numIds: Int
}

type Delete{{$value.TypeName}}Payload {
	{{$value.TypeName}}(filter: {{$value.TypeName}}Filter, order: {{$value.TypeName}}Order, first: Int, offset: Int): [{{$value.TypeName}}]
	numIds: Int
	msg: String
}

input Add{{$value.TypeName}}Input{
	{{- range $fieldKey, $field := $value.Fields}}
  {{$field.Name}}: {{$field.RefGqlType}}!
	{{- end}}
}

input {{$value.TypeName}}Patch{
	{{- range $fieldKey, $field := $value.InputPatchFields}}
  {{$field.Name}}: {{$field.RefGqlType}}
	{{- end}}
}

input Update{{$value.TypeName}}Input{
	filter: {{$value.TypeName}}Filter
	set: {{$value.TypeName}}Patch
	remove: {{$value.TypeName}}Patch
}

{{- end}}
extend type Query {
	{{- range $key, $value := .List}}
		{{- if $value.Query.Get}}
  get{{$value.TypeName}}({{$value.PrimaryField.Name}}: {{$value.PrimaryField.GqlType}}!): {{$value.TypeName}}{{ range $directiveKey, $directive := $value.Query.DirectiveExt}} {{$directive}}{{end}}
		{{- end}}
		{{- if $value.Query.Query}}
  query{{$value.TypeName}}(filter: {{$value.TypeName}}Filter, order: {{$value.TypeName}}Order, first: Int, offset: Int): [{{$value.TypeName}}]{{ range $directiveKey, $directive := $value.Query.DirectiveExt}} {{$directive}}{{end}}
		{{- end}}
		{{- if $value.Query.Aggregate}}
  aggregate{{$value.TypeName}}(filter: {{$value.TypeName}}Filter): {{$value.TypeName}}AggregateResult{{ range $directiveKey, $directive := $value.Query.DirectiveExt}} {{$directive}}{{end}}
			{{- end}}
	{{- end}}
}

extend type Mutation {
	{{- range $key, $value := .List}}
	{{- if $value.Mutation.Add}}
	add{{$value.TypeName}}(input: [Add{{$value.TypeName}}Input!]!): Add{{$value.TypeName}}Payload
	{{- end}}
	{{- if $value.Mutation.Update}}
	update{{$value.TypeName}}(input: Update{{$value.TypeName}}Input!):  Update{{$value.TypeName}}Payload
	{{- end}}
	{{- if $value.Mutation.Delete}}
	delete{{$value.TypeName}}(filter: {{$value.TypeName}}Filter! ): Delete{{$value.TypeName}}Payload
	{{- end}}
	{{- end}}
}