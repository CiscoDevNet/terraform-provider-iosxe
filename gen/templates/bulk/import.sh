terraform import iosxe_{{snakeCase .BulkName}}.example "{{range $i, $e := (bulkParentAttributes .)}}{{if $i}},{{end}}<{{.TfName}}>{{end}}"
