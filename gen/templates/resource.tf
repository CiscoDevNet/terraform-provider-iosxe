resource "iosxe_{{snakeCase .Name}}" "example" {
{{- range  .Attributes}}
{{- if and (not .ExcludeTest) (not .ExcludeExample) (not ( len .TestTags))}}
{{- if eq .Type "List"}}
  {{.TfName}} = [
    {
      {{- range  .Attributes}}
      {{- if and (not .ExcludeTest) (not .ExcludeExample) (not ( len .TestTags))}}
      {{- if eq .Type "List"}}
        {{.TfName}} = [
          {
            {{- range  .Attributes}}
            {{- if and (not .ExcludeTest) (not .ExcludeExample) (not ( len .TestTags))}}
            {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if or (eq .Type "StringList") (eq .Type "StringSet")}}["{{.Example}}"]{{else if or (eq .Type "Int64List") (eq .Type "Int64Set")}}[{{.Example}}]{{else}}{{.Example}}{{end}}
            {{- end}}
            {{- end}}
          }
        ]
      {{- else}}
      {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if or (eq .Type "StringList") (eq .Type "StringSet")}}["{{.Example}}"]{{else if or (eq .Type "Int64List") (eq .Type "Int64Set")}}[{{.Example}}]{{else}}{{.Example}}{{end}}
      {{- end}}
      {{- end}}
      {{- end}}
    }
  ]
{{- else}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if or (eq .Type "StringList") (eq .Type "StringSet")}}["{{.Example}}"]{{else if or (eq .Type "Int64List") (eq .Type "Int64Set")}}[{{.Example}}]{{else}}{{.Example}}{{end}}
{{- end}}
{{- end}}
{{- end}}
}
