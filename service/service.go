// Copyright 2023 ainsley.dev. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"github.com/ainsleydev/hupi/handler"
	"github.com/ainsleydev/hupi/hugo"
	"github.com/ainsleydev/hupi/strapi"
	"strings"
)

type App struct {
	Hugo   hugo.Hugo
	Server handler.Server
	Strapi *strapi.Client
}

type DevelopConfig struct {
	HugoBuildDirectory string
	HugoPort           string
	StrapiEnable       bool
	StrapiNoBuild      bool
	StrapiWatchAdmin   bool
}

func (a App) Develop() error {
	return a.Hugo.Server(nil)
}

var CliCommands = []string{
	"hugo-port",
	"hugo-build-directory",
	"strapi-enable",
	"strapi-no-build",
	"strapi-watch-admin",
}

func getHugoArgs(args []string) []string {
	var hugoArgs []string
Outer:
	for _, arg := range args {
		for _, cliArg := range CliCommands {
			if strings.Contains(arg, cliArg) {
				continue Outer
			}
			hugoArgs = append(hugoArgs, arg)
		}
	}
	return hugoArgs
}
