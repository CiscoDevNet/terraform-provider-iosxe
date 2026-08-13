// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

//go:build ignore

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/openconfig/goyang/pkg/yang"
	"gopkg.in/yaml.v3"
)

const (
	definitionsPath     = "./gen/definitions/"
	fullDefinitionsPath = "./gen/full_definitions/"
	modelsPath          = "./gen/models/"
	providerTemplate    = "./gen/templates/provider.go"
	providerLocation    = "./internal/provider/provider.go"
	changelogTemplate   = "./gen/templates/changelog.md.tmpl"
	changelogLocation   = "./templates/guides/changelog.md.tmpl"
	changelogOriginal   = "./CHANGELOG.md"
)

type t struct {
	path   string
	prefix string
	suffix string
}

var templates = []t{
	{
		path:   "./gen/templates/model.go",
		prefix: "./internal/provider/model_iosxe_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/data_source.go",
		prefix: "./internal/provider/data_source_iosxe_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/data_source_test.go",
		prefix: "./internal/provider/data_source_iosxe_",
		suffix: "_test.go",
	},
	{
		path:   "./gen/templates/resource.go",
		prefix: "./internal/provider/resource_iosxe_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/resource_test.go",
		prefix: "./internal/provider/resource_iosxe_",
		suffix: "_test.go",
	},
	{
		path:   "./gen/templates/data-source.tf",
		prefix: "./examples/data-sources/iosxe_",
		suffix: "/data-source.tf",
	},
	{
		path:   "./gen/templates/resource.tf",
		prefix: "./examples/resources/iosxe_",
		suffix: "/resource.tf",
	},
	{
		path:   "./gen/templates/import.sh",
		prefix: "./examples/resources/iosxe_",
		suffix: "/import.sh",
	},
	{
		path:   "./gen/templates/bulk/model.go",
		prefix: "./internal/provider/model_iosxe_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/bulk/resource.go",
		prefix: "./internal/provider/resource_iosxe_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/bulk/resource_test.go",
		prefix: "./internal/provider/resource_iosxe_",
		suffix: "_test.go",
	},
	{
		path:   "./gen/templates/bulk/resource.tf",
		prefix: "./examples/resources/iosxe_",
		suffix: "/resource.tf",
	},
	{
		path:   "./gen/templates/bulk/import.sh",
		prefix: "./examples/resources/iosxe_",
		suffix: "/import.sh",
	},
}

// isBulkTemplate returns true if the template renders a bulk resource artifact.
func isBulkTemplate(templatePath string) bool {
	return strings.Contains(templatePath, "/bulk/")
}

// skipTemplate returns true if the template should not be rendered for the given config.
func skipTemplate(config YamlConfig, templatePath string) bool {
	if !isBulkTemplate(templatePath) {
		return false
	}
	if !config.BulkResource {
		return true
	}
	return config.SkipBulkResourceTest && templatePath == "./gen/templates/bulk/resource_test.go"
}

type YamlConfig struct {
	Name                    string                `yaml:"name"`
	BulkName                string                `yaml:"bulk_name"`
	Path                    string                `yaml:"path"`
	AugmentPath             string                `yaml:"augment_path"`
	NoDelete                bool                  `yaml:"no_delete"`
	NoDeleteAttributes      bool                  `yaml:"no_delete_attributes"`
	DefaultDeleteAttributes bool                  `yaml:"default_delete_attributes"`
	Wait                    bool                  `yaml:"wait"`
	RequestTimeout          int64                 `yaml:"request_timeout"`
	TestTags                []string              `yaml:"test_tags"`
	SkipMinimumTest         bool                  `yaml:"skip_minimum_test"`
	SkipBulkResourceTest    bool                  `yaml:"skip_bulk_resource_test"`
	NoAugmentConfig         bool                  `yaml:"no_augment_config"`
	BulkResource            bool                  `yaml:"bulk_resource"`
	DsDescription           string                `yaml:"ds_description"`
	ResDescription          string                `yaml:"res_description"`
	ResBulkDescription      string                `yaml:"res_bulk_description"`
	DocCategory             string                `yaml:"doc_category"`
	Attributes              []YamlConfigAttribute `yaml:"attributes"`
	TestPrerequisites       []YamlTest            `yaml:"test_prerequisites"`
}

type YamlConfigAttribute struct {
	YangName  string `yaml:"yang_name"`
	YangScope string `yaml:"yang_scope"`
	TfName    string `yaml:"tf_name"`
	XPath     string `yaml:"xpath"`
	Type      string `yaml:"type"`
	// "empty", "presence" or "boolean"
	TypeYangBool       string                `yaml:"type_yang_bool"`
	Id                 bool                  `yaml:"id"`
	Reference          bool                  `yaml:"reference"`
	Mandatory          bool                  `yaml:"mandatory"`
	Optional           bool                  `yaml:"optional"`
	WriteOnly          bool                  `yaml:"write_only"`
	Sensitive          bool                  `yaml:"sensitive"`
	ExcludeTest        bool                  `yaml:"exclude_test"`
	ExcludeExample     bool                  `yaml:"exclude_example"`
	Description        string                `yaml:"description"`
	Example            string                `yaml:"example"`
	AllowImportChanges bool                  `yaml:"allow_import_changes"`
	EnumValues         []string              `yaml:"enum_values"`
	MinInt             int64                 `yaml:"min_int"`
	MaxInt             int64                 `yaml:"max_int"`
	MinFloat           float64               `yaml:"min_float"`
	MaxFloat           float64               `yaml:"max_float"`
	StringPatterns     []string              `yaml:"string_patterns"`
	StringMinLength    int64                 `yaml:"string_min_length"`
	StringMaxLength    int64                 `yaml:"string_max_length"`
	DefaultValue       string                `yaml:"default_value"`
	RequiresReplace    bool                  `yaml:"requires_replace"`
	NoAugmentConfig    bool                  `yaml:"no_augment_config"`
	DeleteParent       bool                  `yaml:"delete_parent"`
	NoDelete           bool                  `yaml:"no_delete"`
	ReadFilter         string                `yaml:"read_filter"`
	NormalizeIPv6      bool                  `yaml:"normalize_ipv6"`
	TestTags           []string              `yaml:"test_tags"`
	Attributes         []YamlConfigAttribute `yaml:"attributes"`
}

type YamlTest struct {
	Path         string              `yaml:"path"`
	NoDelete     bool                `yaml:"no_delete"`
	Attributes   []YamlTestAttribute `yaml:"attributes"`
	Lists        []YamlTestList      `yaml:"lists"`
	Dependencies []string            `yaml:"dependencies"`
}

type YamlTestAttribute struct {
	Name      string `yaml:"name"`
	Value     string `yaml:"value"`
	Reference string `yaml:"reference"`
}

type YamlTestList struct {
	Name  string             `yaml:"name"`
	Key   string             `yaml:"key"`
	Items []YamlTestListItem `yaml:"items"`
}

type YamlTestListItem struct {
	Attributes []YamlTestAttribute `yaml:"attributes"`
}

// Templating helper function to get short YANG name without prefix (xxx:abc -> abc)
func ToYangShortName(s string) string {
	elements := strings.Split(s, "/")
	for i := range elements {
		if strings.Contains(elements[i], ":") {
			elements[i] = strings.Split(elements[i], ":")[1]
		}
	}
	return strings.Join(elements, "/")
}

// Templating helper function to convert TF name to GO name
func ToGoName(s string) string {
	var g []string

	p := strings.Split(s, "_")

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	s = strings.Join(g, "")
	return s
}

// Templating helper function to convert string to camel case
func CamelCase(s string) string {
	var g []string

	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	return strings.Join(g, "")
}

// Templating helper function to convert string to snake case
func SnakeCase(s string) string {
	var g []string

	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.ToLower(value))
	}
	return strings.Join(g, "_")
}

// Templating helper function to return true if id included in attributes
func HasId(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Id || attr.Reference {
			return true
		}
	}
	return false
}

// Templating helper function to check if any attribute is sensitive (recursive)
func HasSensitiveAttr(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Sensitive {
			return true
		}
		// Recursively check nested list/set attributes
		if (attr.Type == "List" || attr.Type == "Set") && len(attr.Attributes) > 0 {
			if HasSensitiveAttr(attr.Attributes) {
				return true
			}
		}
	}
	return false
}

// Templating helper function to remove last element of path
func RemoveLastPathElement(p string) string {
	return path.Dir(p)
}

// Templating helper function to support arithmetic addition
func Add(a, b int) int {
	return a + b
}

// GetImportExcludes returns a list of attributes to exclude from import testing
func GetImportExcludes(attributes []YamlConfigAttribute) []string {
	var excludes []string
	for _, attr := range attributes {
		if ((attr.TypeYangBool == "empty" || attr.TypeYangBool == "presence") && (attr.ExcludeTest || len(attr.TestTags) > 0)) || attr.AllowImportChanges {
			excludes = append(excludes, attr.TfName)
		}
		if len(attr.Attributes) > 0 {
			ca := GetImportExcludes(attr.Attributes)
			for _, c := range ca {
				excludes = append(excludes, attr.TfName+".0."+c)
			}
		}
	}
	return excludes
}

// Templating helper function to return all import attributes
func ImportAttributes(config YamlConfig) []YamlConfigAttribute {
	attributes := []YamlConfigAttribute{}
	for _, attr := range config.Attributes {
		if attr.Reference || attr.Id {
			attributes = append(attributes, attr)
		}
	}

	return attributes
}

func GetDeletePath(attribute YamlConfigAttribute) string {
	path := attribute.XPath
	if attribute.DeleteParent {
		return RemoveLastPathElement(path)
	}
	return path
}

func ReverseAttributes(attributes []YamlConfigAttribute) []YamlConfigAttribute {
	reversed := make([]YamlConfigAttribute, len(attributes))
	for i, v := range attributes {
		reversed[len(attributes)-1-i] = v
	}
	return reversed
}

// XPathAttributes filters attributes to only those with an XPath defined.
// Attributes without XPath (like disabled_vlans) require custom handling
// and should be skipped in standard body building templates.
func XPathAttributes(attributes []YamlConfigAttribute) []YamlConfigAttribute {
	var filtered []YamlConfigAttribute
	for _, attr := range attributes {
		if attr.XPath != "" {
			filtered = append(filtered, attr)
		}
	}
	return filtered
}

func ToRestconfPath(path string) string {
	return helpers.ConvertXPathToRestconfPath(path)
}

func ToDotPath(path string) string {
	return strings.ReplaceAll(path, "/", ".")
}

// BulkKeySeparator separates the individual key values of a composite bulk map key.
const BulkKeySeparator = ";"

// countFormatVerbs counts the fmt format verbs (e.g. %s, %v) in a path template.
// Escaped percent signs ("%%") are not counted.
func countFormatVerbs(p string) int {
	count := 0
	for i := 0; i < len(p)-1; i++ {
		if p[i] != '%' {
			continue
		}
		if p[i+1] == '%' {
			i++
			continue
		}
		count++
	}
	return count
}

// BulkParentXPath returns the XPath of the container holding the bulk items, which is the
// resource path without its last element.
// Example: "/Cisco-IOS-XE-native:native/interface/%s[name=%v]" -> "/Cisco-IOS-XE-native:native/interface"
func BulkParentXPath(config YamlConfig) string {
	return RemoveLastPathElement(config.Path)
}

// BulkParentAttributes returns the id/reference attributes that are consumed by the parent XPath.
// These become top-level (resource-level) attributes of the bulk resource.
func BulkParentAttributes(config YamlConfig) []YamlConfigAttribute {
	ids := ImportAttributes(config)
	n := countFormatVerbs(BulkParentXPath(config))
	if n > len(ids) {
		n = len(ids)
	}
	return ids[:n]
}

// BulkKeyAttributes returns the id/reference attributes that identify a single item within the
// parent container. Their values are joined by BulkKeySeparator to form the map key.
func BulkKeyAttributes(config YamlConfig) []YamlConfigAttribute {
	ids := ImportAttributes(config)
	n := countFormatVerbs(BulkParentXPath(config))
	if n > len(ids) {
		n = len(ids)
	}
	return ids[n:]
}

// BulkItemAttributes returns all attributes that are rendered inside a bulk item, meaning all
// attributes except the ones that identify the resource or an item within it.
func BulkItemAttributes(config YamlConfig) []YamlConfigAttribute {
	skip := make(map[string]bool)
	for _, attr := range ImportAttributes(config) {
		skip[attr.TfName] = true
	}
	attributes := []YamlConfigAttribute{}
	for _, attr := range config.Attributes {
		if skip[attr.TfName] {
			continue
		}
		attributes = append(attributes, attr)
	}
	return attributes
}

// bulkLastPathSegment returns the last element of the resource path, e.g. "%s[name=%v]".
func bulkLastPathSegment(config YamlConfig) string {
	elements := strings.Split(config.Path, "/")
	return elements[len(elements)-1]
}

// bulkElementNameTemplate returns the element name portion of the last path segment, meaning the
// part before the first XPath predicate. Example: "%s[name=%v]" -> "%s".
func bulkElementNameTemplate(config YamlConfig) string {
	segment := bulkLastPathSegment(config)
	if i := strings.Index(segment, "["); i >= 0 {
		return segment[:i]
	}
	return segment
}

// BulkElementNameIsVerb returns true if the element name of the bulk items is not a literal but
// provided by the first key attribute, e.g. "/native/interface/%s[name=%v]".
func BulkElementNameIsVerb(config YamlConfig) bool {
	return countFormatVerbs(bulkElementNameTemplate(config)) > 0
}

// BulkElementNameExpr returns a Go expression evaluating to the XML element name of a bulk item.
func BulkElementNameExpr(config YamlConfig) string {
	tmpl := bulkElementNameTemplate(config)
	if !BulkElementNameIsVerb(config) {
		return `"` + tmpl + `"`
	}
	if countFormatVerbs(tmpl) == 1 && strings.HasPrefix(tmpl, "%") && len(tmpl) == 2 {
		return "keyParts[0]"
	}
	return `fmt.Sprintf("` + tmpl + `", keyParts[0])`
}

// BulkKeyElementEnumValues returns the possible element names of a bulk item. This is only
// relevant when the element name is provided by a key attribute (see BulkElementNameIsVerb), in
// which case the enum values of that attribute enumerate every element name to look for when
// reading all items from a device.
func BulkKeyElementEnumValues(config YamlConfig) []string {
	if !BulkElementNameIsVerb(config) {
		return []string{bulkElementNameTemplate(config)}
	}
	keys := BulkKeyAttributes(config)
	if len(keys) == 0 {
		return nil
	}
	return keys[0].EnumValues
}

// BulkParentXPathArgs returns the fmt.Sprintf arguments to resolve the format verbs of the parent
// XPath from the resource-level attributes.
func BulkParentXPathArgs(config YamlConfig) string {
	var args []string
	for _, attr := range BulkParentAttributes(config) {
		args = append(args, fmt.Sprintf(`fmt.Sprintf("%%v", data.%s.Value%s())`, ToGoName(attr.TfName), attr.Type))
	}
	return strings.Join(args, ", ")
}

// BulkItemXPathArgs returns the fmt.Sprintf arguments to resolve the format verbs of the resource
// path from the resource-level attributes and the parsed map key.
func BulkItemXPathArgs(config YamlConfig) string {
	args := []string{}
	if a := BulkParentXPathArgs(config); a != "" {
		args = append(args, a)
	}
	for i := range BulkKeyAttributes(config) {
		args = append(args, fmt.Sprintf("keyParts[%d]", i))
	}
	return strings.Join(args, ", ")
}

// BulkMapKeyParse returns the Go statement parsing a bulk map key into its individual values.
func BulkMapKeyParse(keyVar string, config YamlConfig) string {
	return fmt.Sprintf(`keyParts := strings.SplitN(%s, "%s", %d)`, keyVar, BulkKeySeparator, len(BulkKeyAttributes(config)))
}

// BulkMapKeyExample returns the example map key of a bulk item, built from the example values of
// the key attributes.
func BulkMapKeyExample(config YamlConfig) string {
	var parts []string
	for _, attr := range BulkKeyAttributes(config) {
		parts = append(parts, attr.Example)
	}
	return strings.Join(parts, BulkKeySeparator)
}

// BulkMapKeyDescription documents the format of the bulk map key. The returned string is embedded
// in a Go string literal, hence the escaped newlines.
func BulkMapKeyDescription(config YamlConfig) string {
	keys := BulkKeyAttributes(config)
	if len(keys) == 0 {
		return ""
	}
	var names []string
	for _, attr := range keys {
		names = append(names, "`"+attr.TfName+"`")
	}
	var sb strings.Builder
	if len(keys) == 1 {
		sb.WriteString(fmt.Sprintf("\\n  - Map key: %s", names[0]))
	} else {
		sb.WriteString(fmt.Sprintf("\\n  - Map key: %s, joined by `%s`", strings.Join(names, ", "), BulkKeySeparator))
	}
	for _, attr := range keys {
		if len(attr.EnumValues) == 0 {
			continue
		}
		quoted := make([]string, len(attr.EnumValues))
		for i, v := range attr.EnumValues {
			quoted[i] = "`" + v + "`"
		}
		sb.WriteString(fmt.Sprintf("\\n  - Choices for `%s`: %s", attr.TfName, strings.Join(quoted, ", ")))
	}
	return sb.String()
}

// pluralize returns the plural form of the last word of s.
func pluralize(s string) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return s
	}
	last := words[len(words)-1]
	lower := strings.ToLower(last)
	switch {
	case strings.HasSuffix(lower, "y") && !strings.HasSuffix(lower, "ay") && !strings.HasSuffix(lower, "ey") &&
		!strings.HasSuffix(lower, "iy") && !strings.HasSuffix(lower, "oy") && !strings.HasSuffix(lower, "uy"):
		last = last[:len(last)-1] + "ies"
	case strings.HasSuffix(lower, "s"), strings.HasSuffix(lower, "x"), strings.HasSuffix(lower, "z"),
		strings.HasSuffix(lower, "ch"), strings.HasSuffix(lower, "sh"):
		last += "es"
	default:
		last += "s"
	}
	words[len(words)-1] = last
	return strings.Join(words, " ")
}

// initBulkNames fills in the derived bulk name and description. It does not depend on the YANG
// augmentation and is therefore applied to every config, so that provider.go can be rendered
// correctly even when the generator only regenerates a single resource.
func initBulkNames(config *YamlConfig) {
	if !config.BulkResource {
		return
	}
	if config.BulkName == "" {
		config.BulkName = pluralize(config.Name)
	}
	if config.ResBulkDescription == "" {
		config.ResBulkDescription = fmt.Sprintf("This resource can manage the %s configuration of multiple items in a single resource.", config.Name)
	}
}

// validateBulkConfig verifies that a bulk resource can actually be generated from the definition.
// It requires an augmented config, as the key attributes are derived from the YANG models.
func validateBulkConfig(config YamlConfig) {
	if !config.BulkResource {
		return
	}
	if len(BulkKeyAttributes(config)) == 0 {
		log.Fatalf("%q: 'bulk_resource' requires a path ending in a list element, got %q", config.Name, config.Path)
	}
	if BulkElementNameIsVerb(config) && len(BulkKeyElementEnumValues(config)) == 0 {
		log.Fatalf("%q: 'bulk_resource' requires 'enum_values' on attribute %q, as it provides the element name of an item",
			config.Name, BulkKeyAttributes(config)[0].TfName)
	}
}

// Map of templating functions
var functions = template.FuncMap{
	"toGoName":                 ToGoName,
	"camelCase":                CamelCase,
	"snakeCase":                SnakeCase,
	"hasId":                    HasId,
	"hasSensitiveAttr":         HasSensitiveAttr,
	"add":                      Add,
	"getImportExcludes":        GetImportExcludes,
	"importAttributes":         ImportAttributes,
	"getDeletePath":            GetDeletePath,
	"reverseAttributes":        ReverseAttributes,
	"xpathAttributes":          XPathAttributes,
	"toRestconfPath":           ToRestconfPath,
	"toDotPath":                ToDotPath,
	"bulkParentXPath":          BulkParentXPath,
	"bulkParentAttributes":     BulkParentAttributes,
	"bulkKeyAttributes":        BulkKeyAttributes,
	"bulkItemAttributes":       BulkItemAttributes,
	"bulkElementNameIsVerb":    BulkElementNameIsVerb,
	"bulkElementNameExpr":      BulkElementNameExpr,
	"bulkKeyElementEnumValues": BulkKeyElementEnumValues,
	"bulkParentXPathArgs":      BulkParentXPathArgs,
	"bulkItemXPathArgs":        BulkItemXPathArgs,
	"bulkMapKeyParse":          BulkMapKeyParse,
	"bulkMapKeyExample":        BulkMapKeyExample,
	"bulkMapKeyDescription":    BulkMapKeyDescription,
}

func resolvePath(e *yang.Entry, path string) *yang.Entry {
	pathElements := strings.Split(path, "/")

	for _, pathElement := range pathElements {
		if len(pathElement) > 0 {
			// remove XPath predicate (e.g., [name=value] or [name=%v])
			if strings.Contains(pathElement, "[") {
				pathElement = pathElement[:strings.Index(pathElement, "[")]
			}
			// remove namespace prefix (e.g., Cisco-IOS-XE-bgp:bgp -> bgp)
			if strings.Contains(pathElement, ":") {
				pathElement = pathElement[strings.Index(pathElement, ":")+1:]
			}
			if _, ok := e.Dir[pathElement]; !ok {
				panic(fmt.Sprintf("Failed to resolve YANG path: %s, element: %s", path, pathElement))
			}
			e = e.Dir[pathElement]
		}
	}

	return e
}

func addKeys(e *yang.Entry, config *YamlConfig) {
	first := true
	for {
		if e.Key != "" {
			keys := strings.Split(e.Key, " ")
			for _, key := range keys {
				var keyAttr *YamlConfigAttribute
				// check if key attribute already in config
				for i := range config.Attributes {
					if config.Attributes[i].YangScope != "" && config.Attributes[i].YangScope != e.Name {
						continue
					}
					if config.Attributes[i].YangName == key {
						keyAttr = &config.Attributes[i]
						break
					}
				}
				if keyAttr == nil {
					continue
				}
				if first {
					keyAttr.Id = true
					keyAttr.Reference = false
				} else {
					keyAttr.Id = false
					keyAttr.Reference = true
				}
				parseAttribute(e, keyAttr)
			}
		}
		first = false
		if e.Parent != nil {
			e = e.Parent
			continue
		}
		break
	}
}

func parseAttribute(e *yang.Entry, attr *YamlConfigAttribute) {
	leaf := resolvePath(e, attr.YangName)
	//fmt.Printf("%s, Entry: %+v\n\n", attr.YangName, e)
	//fmt.Printf("%s, Kind: %+v, ListAttr: %+v, Type: %+v\n\n", leaf.Name, leaf.Kind, leaf.ListAttr, leaf.Type)
	if leaf.Kind.String() == "Leaf" {
		if leaf.ListAttr != nil {
			if helpers.Contains([]string{"string", "union", "leafref", "enumeration"}, leaf.Type.Kind.String()) {
				if attr.Type == "" {
					attr.Type = "StringList"
				}
			} else if helpers.Contains([]string{"uint8", "uint16", "uint32", "uint64"}, leaf.Type.Kind.String()) {
				if attr.Type == "" {
					attr.Type = "Int64List"
				}
			} else {
				panic(fmt.Sprintf("Unknown leaf-list type, attribute: %s, type: %s", attr.YangName, leaf.Type.Kind.String()))
			}
			// TODO parse union type
		} else if helpers.Contains([]string{"string", "union", "leafref"}, leaf.Type.Kind.String()) {
			if attr.Type == "" {
				attr.Type = "String"
			}
			if leaf.Type.Length != nil {
				if attr.StringMinLength == 0 {
					attr.StringMinLength = int64(leaf.Type.Length[0].Min.Value)
				}
				max := leaf.Type.Length[0].Max.Value
				// hack to not introduce unsigned types
				if max > math.MaxInt64 {
					max = math.MaxInt64
				}
				if attr.StringMaxLength == 0 {
					attr.StringMaxLength = int64(max)
				}
			}
			if len(leaf.Type.Pattern) > 0 {
				if len(attr.StringPatterns) == 0 {
					attr.StringPatterns = leaf.Type.Pattern
				}
			}
		} else if helpers.Contains([]string{"uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64"}, leaf.Type.Kind.String()) {
			attr.Type = "Int64"
			if leaf.Type.Range != nil {
				if attr.MinInt == 0 {
					attr.MinInt = int64(leaf.Type.Range[0].Min.Value)
					if leaf.Type.Range[0].Min.Negative {
						attr.MinInt = -attr.MinInt
					}
				}
				max := leaf.Type.Range[0].Max.Value
				// hack to not introduce unsigned types
				if max > math.MaxInt64 {
					max = math.MaxInt64
				}
				if attr.MaxInt == 0 {
					attr.MaxInt = int64(max)
				}
			}
		} else if helpers.Contains([]string{"decimal8", "decimal16", "decimal32", "decimal64"}, leaf.Type.Kind.String()) {
			if attr.Type == "" {
				attr.Type = "Float64"
			}
			if leaf.Type.Range != nil {
				if attr.MinFloat == 0 {
					attr.MinFloat = float64(leaf.Type.Range[0].Min.Value)
					if leaf.Type.Range[0].Min.Negative {
						attr.MinFloat = -attr.MinFloat
					}
				}
				if attr.MaxFloat == 0 {
					attr.MaxFloat = float64(leaf.Type.Range[0].Max.Value)
				}
			}
		} else if helpers.Contains([]string{"boolean", "empty"}, leaf.Type.Kind.String()) {
			if leaf.Type.Kind.String() == "boolean" {
				if attr.TypeYangBool == "" {
					attr.TypeYangBool = "boolean"
				}
			} else if leaf.Type.Kind.String() == "empty" {
				if attr.TypeYangBool == "" {
					attr.TypeYangBool = "empty"
				}
			}
			if attr.Type == "" {
				attr.Type = "Bool"
			}
		} else if helpers.Contains([]string{"enumeration"}, leaf.Type.Kind.String()) {
			if attr.Type == "" {
				attr.Type = "String"
			}
			if len(attr.EnumValues) == 0 {
				attr.EnumValues = leaf.Type.Enum.Names()
			}
		} else {
			panic(fmt.Sprintf("Unknown leaf type, attribute: %s, type: %s", attr.YangName, leaf.Type.Kind.String()))
		}
	}
	if _, ok := leaf.Extra["presence"]; ok {
		if attr.TypeYangBool == "" {
			attr.TypeYangBool = "presence"
		}
		if attr.Type == "" {
			attr.Type = "Bool"
		}
	}
	if attr.XPath == "" {
		attr.XPath = attr.YangName
	}
	if attr.TfName == "" {
		tfName := strings.ReplaceAll(ToYangShortName(attr.XPath), "-", "_")
		tfName = strings.ReplaceAll(tfName, "/", "_")
		attr.TfName = tfName
	}
	if attr.Description == "" {
		attr.Description = strings.ReplaceAll(leaf.Description, "\n", " ")
	}
	if !attr.Mandatory && attr.DefaultValue == "" && !attr.Optional {
		foundChoice := false
		parent := leaf.Parent
		for parent != nil {
			if parent.IsChoice() {
				foundChoice = true
				break
			}
			parent = parent.Parent
		}
		if !foundChoice {
			attr.Mandatory = leaf.Mandatory.Value()
		}
	}
}

func augmentConfig(config *YamlConfig, yangModules *yang.Modules) {
	path := ""
	if config.AugmentPath != "" {
		path = config.AugmentPath
	} else {
		path = config.Path
	}

	path = strings.TrimPrefix(path, "/")
	module := strings.Split(path, ":")[0]
	e, errors := yangModules.GetModule(module)
	if len(errors) > 0 {
		fmt.Printf("YANG parser error(s): %+v\n\n", errors)
		return
	}

	p := path[len(module)+1:]
	e = resolvePath(e, p)

	addKeys(e, config)

	for ia := range config.Attributes {
		if config.Attributes[ia].Id || config.Attributes[ia].Reference || config.Attributes[ia].NoAugmentConfig {
			continue
		}
		// Skip YANG parsing for attributes without yang_name - these are custom attributes
		// that must have tf_name, type, and description explicitly provided in the YAML.
		// They will be included in schema/struct generation but skipped in body building/reading.
		if config.Attributes[ia].YangName == "" {
			continue
		}
		parseAttribute(e, &config.Attributes[ia])
		if config.Attributes[ia].Type == "List" || config.Attributes[ia].Type == "Set" {
			el := resolvePath(e, config.Attributes[ia].YangName)
			for iaa := range config.Attributes[ia].Attributes {
				if config.Attributes[ia].Attributes[iaa].NoAugmentConfig {
					continue
				}
				// Skip YANG parsing for nested attributes without yang_name
				if config.Attributes[ia].Attributes[iaa].YangName == "" {
					continue
				}
				parseAttribute(el, &config.Attributes[ia].Attributes[iaa])
				if config.Attributes[ia].Attributes[iaa].Type == "List" || config.Attributes[ia].Attributes[iaa].Type == "Set" {
					ell := resolvePath(el, config.Attributes[ia].Attributes[iaa].YangName)
					for iaaa := range config.Attributes[ia].Attributes[iaa].Attributes {
						if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].NoAugmentConfig {
							continue
						}
						// Skip YANG parsing for deeply nested attributes without yang_name
						if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].YangName == "" {
							continue
						}
						parseAttribute(ell, &config.Attributes[ia].Attributes[iaa].Attributes[iaaa])
						if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Type == "List" || config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Type == "Set" {
							elll := resolvePath(ell, config.Attributes[ia].Attributes[iaa].Attributes[iaaa].YangName)
							for iaaaa := range config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes {
								if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].NoAugmentConfig {
									continue
								}
								if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].YangName == "" {
									continue
								}
								parseAttribute(elll, &config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa])
							}
						}
					}
				}
			}
		}
	}

	if config.DsDescription == "" {
		config.DsDescription = fmt.Sprintf("This data source can read the %s configuration.", config.Name)
	}
	if config.ResDescription == "" {
		config.ResDescription = fmt.Sprintf("This resource can manage the %s configuration.", config.Name)
	}
}

func getTemplateSection(content, name string) string {
	scanner := bufio.NewScanner(strings.NewReader(content))
	result := ""
	foundSection := false
	beginRegex := regexp.MustCompile(`\/\/template:begin\s` + name + `$`)
	endRegex := regexp.MustCompile(`\/\/template:end\s` + name + `$`)
	for scanner.Scan() {
		line := scanner.Text()
		if !foundSection {
			match := beginRegex.MatchString(line)
			if match {
				foundSection = true
				result += line + "\n"
			}
		} else {
			result += line + "\n"
			match := endRegex.MatchString(line)
			if match {
				foundSection = false
			}
		}
	}
	return result
}

func renderTemplate(templatePath, outputPath string, config interface{}) {
	file, err := os.Open(templatePath)
	if err != nil {
		log.Fatalf("Error opening template: %v", err)
	}
	defer file.Close()

	// skip first line with 'build-ignore' directive for go files
	scanner := bufio.NewScanner(file)
	if strings.HasSuffix(templatePath, ".go") {
		scanner.Scan()
	}
	var temp string
	for scanner.Scan() {
		temp = temp + scanner.Text() + "\n"
	}

	template, err := template.New(path.Base(templatePath)).Funcs(functions).Parse(temp)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	output := new(bytes.Buffer)
	err = template.Execute(output, config)
	if err != nil {
		log.Fatalf("Error executing template for %s: %v", outputPath, err)
	}

	outputFile := filepath.Join(outputPath)
	existingFile, err := os.Open(outputPath)
	if err != nil {
		os.MkdirAll(filepath.Dir(outputFile), 0755)
	} else if strings.HasSuffix(templatePath, ".go") {
		existingScanner := bufio.NewScanner(existingFile)
		var newContent string
		currentSectionName := ""
		beginRegex := regexp.MustCompile(`\/\/template:begin\s(.*?)$`)
		endRegex := regexp.MustCompile(`\/\/template:end\s(.*?)$`)
		for existingScanner.Scan() {
			line := existingScanner.Text()
			if currentSectionName == "" {
				matches := beginRegex.FindStringSubmatch(line)
				if len(matches) > 1 && matches[1] != "" {
					currentSectionName = matches[1]
				} else {
					newContent += line + "\n"
				}
			} else {
				matches := endRegex.FindStringSubmatch(line)
				if len(matches) > 1 && matches[1] == currentSectionName {
					currentSectionName = ""
					newSection := getTemplateSection(string(output.Bytes()), matches[1])
					newContent += newSection
				}
			}
		}
		output = bytes.NewBufferString(newContent)
	}
	// write to output file
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	f.Write(output.Bytes())
}

func main() {
	var writeFlag bool
	flag.BoolVar(&writeFlag, "w", false, "Write full definitions")
	flag.Parse()

	resourceName := ""
	if len(os.Args) == 2 {
		resourceName = os.Args[1]
	}

	// Load configs
	var configs []YamlConfig
	files, _ := os.ReadDir(definitionsPath)

	for _, filename := range files {
		path := filepath.Join(definitionsPath, filename.Name())
		bytes, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("Error reading file %q: %v", path, err)
		}

		config := YamlConfig{}
		err = yaml.Unmarshal(bytes, &config)
		if err != nil {
			log.Fatalf("Error parsing %q: %v", path, err)
		}
		configs = append(configs, config)
	}

	items, _ := os.ReadDir(modelsPath)

	yangModules := yang.NewModules()

	// Iterate over yang models
	for _, item := range items {
		if filepath.Ext(item.Name()) != ".yang" {
			continue
		}

		fn := filepath.Join(modelsPath, item.Name())
		if err := yangModules.Read(fn); err != nil {
			log.Fatalf("yang parser: %v", err)
		}
	}

	for i := range configs {
		initBulkNames(&configs[i])
	}

	for i := range configs {
		if resourceName != "" && configs[i].Name != resourceName {
			continue
		}
		// Augment config by yang models
		if !configs[i].NoAugmentConfig {
			augmentConfig(&configs[i], yangModules)
		}
		validateBulkConfig(configs[i])

		fmt.Printf("Augmented %d/%d: %v\n", i+1, len(configs), configs[i].Name)

		if writeFlag {
			// Write full definitions
			yamlFile, err := yaml.Marshal(&configs[i])
			if err != nil {
				log.Fatalf("Error marshalling yaml: %v", err)
			}

			outputFile := filepath.Join(fullDefinitionsPath, SnakeCase(configs[i].Name)+".yaml")
			err = os.WriteFile(outputFile, yamlFile, 0644)
			if err != nil {
				log.Fatalf("Error writing YAML file: %v", err)
			}
		} else {
			// Iterate over templates and render files
			for _, t := range templates {
				if skipTemplate(configs[i], t.path) {
					continue
				}
				name := configs[i].Name
				if isBulkTemplate(t.path) {
					name = configs[i].BulkName
				}
				renderTemplate(t.path, t.prefix+SnakeCase(name)+t.suffix, configs[i])
			}
		}
	}

	if !writeFlag {
		// render provider.go
		renderTemplate(providerTemplate, providerLocation, configs)

		changelog, err := os.ReadFile(changelogOriginal)
		if err != nil {
			log.Fatalf("Error reading changelog: %v", err)
		}
		renderTemplate(changelogTemplate, changelogLocation, string(changelog))
	}
}
