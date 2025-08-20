data "iosxe_{{snakeCase .Name}}" "example" {
{{- range  .Attributes}}
{{- if or .Id .Reference}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if or (eq .Type "StringList") (eq .Type "StringSet")}}["{{.Example}}"]{{else if or (eq .Type "Int64List") (eq .Type "Int64Set")}}[{{.Example}}]{{else}}{{.Example}}{{end}}
{{- end}}
{{- end}}
}
