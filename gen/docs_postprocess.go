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
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

var generalResources = []string{"save_config", "cli", "commit", "yang"}

const (
	definitionsPath = "./gen/definitions/"
)

type YamlConfig struct {
	Name         string `yaml:"name"`
	BulkName     string `yaml:"bulk_name"`
	BulkResource bool   `yaml:"bulk_resource"`
	DocCategory  string `yaml:"doc_category"`
}

var docPaths = []string{"./docs/data-sources/", "./docs/resources/"}
var generalDocPaths = []string{"./docs/data-sources/", "./docs/resources/", "./docs/actions/"}

func SnakeCase(s string) string {
	var g []string

	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.ToLower(value))
	}
	return strings.Join(g, "_")
}

// pluralize returns the plural form of the last word of s. It must stay in sync with the
// identically named function in gen/generator.go, which derives the default bulk resource name.
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

// pageTitleRegex matches the default tfplugindocs page title, e.g. "iosxe_bgp Resource - terraform-provider-iosxe"
var pageTitleRegex = regexp.MustCompile(`(?m)^page_title: "(\S+) (Resource|Data Source|Action) - terraform-provider-iosxe"$`)

// cleanFrontmatter shortens the page title to e.g. "iosxe_bgp (Resource)" and reduces the
// description to its first line. Both are only used for page titles and meta descriptions,
// e.g. by the OpenTofu registry.
func cleanFrontmatter(s string) string {
	s = pageTitleRegex.ReplaceAllString(s, `page_title: "$1 ($2)"`)

	lines := strings.Split(s, "\n")
	out := make([]string, 0, len(lines))
	inDescription, keptLine := false, false
	for i, line := range lines {
		if i > 0 && line == "---" {
			// end of frontmatter, keep the rest unchanged
			out = append(out, lines[i:]...)
			break
		}
		if line == "description: |-" {
			inDescription, keptLine = true, false
			out = append(out, line)
			continue
		}
		if inDescription && strings.HasPrefix(line, "  ") {
			if !keptLine {
				out = append(out, line)
				keptLine = true
			}
			continue
		}
		inDescription = false
		out = append(out, line)
	}
	return strings.Join(out, "\n")
}

func main() {
	items, _ := os.ReadDir(definitionsPath)
	configs := make([]YamlConfig, len(items))

	// Load configs
	for i, filename := range items {
		yamlFile, err := os.ReadFile(filepath.Join(definitionsPath, filename.Name()))
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

	for i := range configs {
		for _, path := range docPaths {
			filename := path + SnakeCase(configs[i].Name) + ".md"
			content, err := os.ReadFile(filename)
			if err != nil {
				log.Fatalf("Error opening documentation: %v", err)
			}

			s := string(content)
			s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "`+configs[i].DocCategory+`"`)

			os.WriteFile(filename, []byte(s), 0644)
		}

		// Bulk resources only have a resource, no data source
		if !configs[i].BulkResource {
			continue
		}
		bulkName := configs[i].BulkName
		if bulkName == "" {
			bulkName = pluralize(configs[i].Name)
		}
		filename := "./docs/resources/" + SnakeCase(bulkName) + ".md"
		content, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalf("Error opening documentation: %v", err)
		}
		s := strings.ReplaceAll(string(content), `subcategory: ""`, `subcategory: "`+configs[i].DocCategory+`"`)
		os.WriteFile(filename, []byte(s), 0644)
	}

	// Update general resources with "General" subcategory
	for _, resource := range generalResources {
		for _, path := range generalDocPaths {
			filename := fmt.Sprintf("%s%s.md", path, resource)
			content, err := os.ReadFile(filename)
			if err != nil {
				// Skip if file doesn't exist (e.g., data source may not exist for all resources)
				if os.IsNotExist(err) {
					continue
				}
				log.Fatalf("Error opening documentation: %v", err)
			}
			s := string(content)
			s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "General"`)
			os.WriteFile(filename, []byte(s), 0644)
		}
	}

	// Clean up frontmatter of all docs
	for _, path := range generalDocPaths {
		files, _ := filepath.Glob(filepath.Join(path, "*.md"))
		for _, filename := range files {
			content, err := os.ReadFile(filename)
			if err != nil {
				log.Fatalf("Error reading file: %v", err)
			}
			os.WriteFile(filename, []byte(cleanFrontmatter(string(content))), 0644)
		}
	}
}
