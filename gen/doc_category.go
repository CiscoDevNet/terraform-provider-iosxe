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
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	definitionsPath = "./gen/definitions/"
)

type YamlConfig struct {
	Name        string `yaml:"name"`
	DocCategory string `yaml:"doc_category"`
}

var docPaths = []string{"./docs/data-sources/", "./docs/resources/"}

func SnakeCase(s string) string {
	var g []string

	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.ToLower(value))
	}
	return strings.Join(g, "_")
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

	for i := range configs {
		for _, path := range docPaths {
			filename := path + SnakeCase(configs[i].Name) + ".md"
			content, err := ioutil.ReadFile(filename)
			if err != nil {
				log.Fatalf("Error opening documentation: %v", err)
			}

			s := string(content)
			s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "`+configs[i].DocCategory+`"`)

			ioutil.WriteFile(filename, []byte(s), 0644)
		}
	}

	// update iosxe_restconf resource and data source
	for _, path := range docPaths {
		filename := path + "restconf.md"
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalf("Error opening documentation: %v", err)
		}

		s := string(content)
		s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "General"`)

		ioutil.WriteFile(filename, []byte(s), 0644)
	}
}
