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
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/openconfig/goyang/pkg/yang"
	"gopkg.in/yaml.v3"
)

const (
	definitionsPath   = "./gen/definitions/"
	modelsPath        = "./gen/models/"
	providerTemplate  = "./gen/templates/provider.go"
	providerLocation  = "./internal/provider/provider.go"
	changelogTemplate = "./gen/templates/changelog.md.tmpl"
	changelogLocation = "./templates/guides/changelog.md.tmpl"
	changelogOriginal = "./CHANGELOG.md"
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
}

type YamlConfig struct {
	Name                    string                `yaml:"name"`
	Path                    string                `yaml:"path"`
	AugmentPath             string                `yaml:"augment_path"`
	NoDelete                bool                  `yaml:"no_delete"`
	NoDeleteAttributes      bool                  `yaml:"no_delete_attributes"`
	DefaultDeleteAttributes bool                  `yaml:"default_delete_attributes"`
	TestTags                []string              `yaml:"test_tags"`
	SkipMinimumTest         bool                  `yaml:"skip_minimum_test"`
	NoAugmentConfig         bool                  `yaml:"no_augment_config"`
	DsDescription           string                `yaml:"ds_description"`
	ResDescription          string                `yaml:"res_description"`
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
	TypeYangBool    string                `yaml:"type_yang_bool"`
	Id              bool                  `yaml:"id"`
	Reference       bool                  `yaml:"reference"`
	Mandatory       bool                  `yaml:"mandatory"`
	Optional        bool                  `yaml:"optional"`
	WriteOnly       bool                  `yaml:"write_only"`
	ExcludeTest     bool                  `yaml:"exclude_test"`
	ExcludeExample  bool                  `yaml:"exclude_example"`
	Description     string                `yaml:"description"`
	Example         string                `yaml:"example"`
	EnumValues      []string              `yaml:"enum_values"`
	MinInt          int64                 `yaml:"min_int"`
	MaxInt          int64                 `yaml:"max_int"`
	MinFloat        float64               `yaml:"min_float"`
	MaxFloat        float64               `yaml:"max_float"`
	StringPatterns  []string              `yaml:"string_patterns"`
	StringMinLength int64                 `yaml:"string_min_length"`
	StringMaxLength int64                 `yaml:"string_max_length"`
	DefaultValue    string                `yaml:"default_value"`
	RequiresReplace bool                  `yaml:"requires_replace"`
	NoAugmentConfig bool                  `yaml:"no_augment_config"`
	DeleteParent    bool                  `yaml:"delete_parent"`
	NoDelete        bool                  `yaml:"no_delete"`
	TestTags        []string              `yaml:"test_tags"`
	Attributes      []YamlConfigAttribute `yaml:"attributes"`
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

// Templating helper function to get short YAMG name without prefix (xxx:abc -> abc)
func ToYangShortName(s string) string {
	if strings.Contains(s, ":") {
		s = strings.Split(s, ":")[1]
	}
	return s
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

// Templating helper function to convert YANG name to GO name
func ToJsonPath(yangPath, xPath string) string {
	if xPath != "" {
		return strings.ReplaceAll(xPath, "/", ".")
	}
	return strings.ReplaceAll(yangPath, "/", ".")
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

// Templating helper function to get example dn
func GetExamplePath(path string, attributes []YamlConfigAttribute) string {
	a := make([]interface{}, 0, len(attributes))
	for _, attr := range attributes {
		if attr.Id || attr.Reference {
			a = append(a, attr.Example)
		}
	}
	return fmt.Sprintf(path, a...)
}

// Templating helper function to identify last element of list
func IsLast(index int, len int) bool {
	return index+1 == len
}

// Templating helper function to remove last element of path
func RemoveLastPathElement(p string) string {
	return path.Dir(p)
}

// Templating helper function to get xpath if available
func GetXPath(yangPath, xPath string) string {
	if xPath != "" {
		return xPath
	}
	return yangPath
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// Map of templating functions
var functions = template.FuncMap{
	"toGoName":              ToGoName,
	"toJsonPath":            ToJsonPath,
	"camelCase":             CamelCase,
	"snakeCase":             SnakeCase,
	"hasId":                 HasId,
	"getExamplePath":        GetExamplePath,
	"isLast":                IsLast,
	"sprintf":               fmt.Sprintf,
	"removeLastPathElement": RemoveLastPathElement,
	"getXPath":              GetXPath,
	"contains":              contains,
}

func resolvePath(e *yang.Entry, path string) *yang.Entry {
	pathElements := strings.Split(path, "/")

	for _, pathElement := range pathElements {
		if len(pathElement) > 0 {
			// remove key
			if strings.Contains(pathElement, "=") {
				pathElement = pathElement[:strings.Index(pathElement, "=")]
			}
			// remove reference
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
			if contains([]string{"string", "union", "leafref"}, leaf.Type.Kind.String()) {
				attr.Type = "StringList"
			} else if contains([]string{"uint8", "uint16", "uint32", "uint64"}, leaf.Type.Kind.String()) {
				attr.Type = "Int64List"
			} else {
				panic(fmt.Sprintf("Unknown leaf-list type, attribute: %s, type: %s", attr.YangName, leaf.Type.Kind.String()))
			}
			// TODO parse union type
		} else if contains([]string{"string", "union", "leafref"}, leaf.Type.Kind.String()) {
			attr.Type = "String"
			if leaf.Type.Length != nil {
				attr.StringMinLength = int64(leaf.Type.Length[0].Min.Value)
				max := leaf.Type.Length[0].Max.Value
				// hack to not introduce unsigned types
				if max > math.MaxInt64 {
					max = math.MaxInt64
				}
				attr.StringMaxLength = int64(max)
			}
			if len(leaf.Type.Pattern) > 0 {
				attr.StringPatterns = leaf.Type.Pattern
			}
		} else if contains([]string{"uint8", "uint16", "uint32", "uint64"}, leaf.Type.Kind.String()) {
			attr.Type = "Int64"
			if leaf.Type.Range != nil {
				attr.MinInt = int64(leaf.Type.Range[0].Min.Value)
				max := leaf.Type.Range[0].Max.Value
				// hack to not introduce unsigned types
				if max > math.MaxInt64 {
					max = math.MaxInt64
				}
				attr.MaxInt = int64(max)
			}
		} else if contains([]string{"decimal8", "decimal16", "decimal32", "decimal64"}, leaf.Type.Kind.String()) {
			attr.Type = "Float64"
			if leaf.Type.Range != nil {
				attr.MinFloat = float64(leaf.Type.Range[0].Min.Value)
				attr.MaxFloat = float64(leaf.Type.Range[0].Max.Value)
			}
		} else if contains([]string{"boolean", "empty"}, leaf.Type.Kind.String()) {
			if leaf.Type.Kind.String() == "boolean" {
				attr.TypeYangBool = "boolean"
			} else if leaf.Type.Kind.String() == "empty" {
				attr.TypeYangBool = "empty"
			}
			attr.Type = "Bool"
		} else if contains([]string{"enumeration"}, leaf.Type.Kind.String()) {
			attr.Type = "String"
			attr.EnumValues = leaf.Type.Enum.Names()
		} else {
			panic(fmt.Sprintf("Unknown leaf type, attribute: %s, type: %s", attr.YangName, leaf.Type.Kind.String()))
		}
	}
	if _, ok := leaf.Extra["presence"]; ok {
		attr.TypeYangBool = "presence"
		attr.Type = "Bool"
	}
	if attr.TfName == "" {
		tfName := strings.ReplaceAll(ToYangShortName(attr.YangName), "-", "_")
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

func augmentConfig(config *YamlConfig, modelPaths []string) {
	path := ""
	if config.AugmentPath != "" {
		path = config.AugmentPath
	} else {
		path = config.Path
	}

	module := strings.Split(path, ":")[0]
	e, errors := yang.GetModule(module, modelPaths...)
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
		parseAttribute(e, &config.Attributes[ia])
		if config.Attributes[ia].Type == "List" {
			el := resolvePath(e, config.Attributes[ia].YangName)
			for iaa := range config.Attributes[ia].Attributes {
				if config.Attributes[ia].Attributes[iaa].NoAugmentConfig {
					continue
				}
				parseAttribute(el, &config.Attributes[ia].Attributes[iaa])
				if config.Attributes[ia].Attributes[iaa].Type == "List" {
					ell := resolvePath(el, config.Attributes[ia].Attributes[iaa].YangName)
					for iaaa := range config.Attributes[ia].Attributes[iaa].Attributes {
						if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].NoAugmentConfig {
							continue
						}
						parseAttribute(ell, &config.Attributes[ia].Attributes[iaa].Attributes[iaaa])
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

	// create output file
	outputFile := filepath.Join(outputPath)
	os.MkdirAll(filepath.Dir(outputFile), 0755)
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}

	output := new(bytes.Buffer)
	err = template.Execute(output, config)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
	f.Write(output.Bytes())
}

func main() {
	items, _ := ioutil.ReadDir(definitionsPath)
	configs := make([]YamlConfig, len(items))

	// Load configs
	for i, filename := range items {
		yamlFile, err := ioutil.ReadFile(filepath.Join(definitionsPath, filename.Name()))
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		config := YamlConfig{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalf("Error parsing yaml: %v", err)
		}
		configs[i] = config
	}

	items, _ = ioutil.ReadDir(modelsPath)
	modelPaths := make([]string, 0)

	// Iterate over yang models
	for _, item := range items {
		if filepath.Ext(item.Name()) == ".yang" {
			modelPaths = append(modelPaths, filepath.Join(modelsPath, item.Name()))
		}
	}

	for i := range configs {
		// Augment config by yang models
		if !configs[i].NoAugmentConfig {
			augmentConfig(&configs[i], modelPaths)
		}

		// Iterate over templates and render files
		for _, t := range templates {
			renderTemplate(t.path, t.prefix+SnakeCase(configs[i].Name)+t.suffix, configs[i])
		}
	}

	// render provider.go
	renderTemplate(providerTemplate, providerLocation, configs)

	changelog, err := ioutil.ReadFile(changelogOriginal)
	if err != nil {
		log.Fatalf("Error reading changelog: %v", err)
	}
	renderTemplate(changelogTemplate, changelogLocation, string(changelog))
}
