resource "iosxe_{{snakeCase .BulkName}}" "example" {
{{- range bulkParentAttributes .}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else}}{{.Example}}{{end}}
{{- end}}
  items = {
    "{{bulkMapKeyExample .}}" = {
    {{- range bulkItemAttributes .}}
    {{- if and (not .ExcludeTest) (not .ExcludeExample) (not ( len .TestTags))}}
    {{- if or (eq .Type "List") (eq .Type "Set")}}
      {{.TfName}} = [
        {
          {{- range .Attributes}}
          {{- if and (not .ExcludeTest) (not .ExcludeExample) (not ( len .TestTags))}}
          {{- if or (eq .Type "List") (eq .Type "Set")}}
          {{.TfName}} = [
            {
              {{- range .Attributes}}
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
  }
}
