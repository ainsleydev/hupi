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

type (
	Develop struct {
		Hugo   hugo.Hugo
		Server handler.Server
		Strapi *strapi.Client
		Config DevelopConfig
	}
	DevelopConfig struct {
		Args               []string
		HugoPort           string
		HugoBuildDirectory string
		StrapiEnable       bool
		StrapiNoBuild      bool
		StrapiWatchAdmin   bool
	}
)

// NewDevelop creates a new develop type that is used for
// dev environments.
func NewDevelop(config DevelopConfig) *Develop {
	h := hugo.Client{
		BuildDirectory: config.HugoBuildDirectory,
	}
	return &Develop{
		Hugo: h,
		Server: handler.Server{
			Addr: config.HugoPort,
		},
		Strapi: nil,
		Config: config,
	}
}

func (d Develop) Develop() error {
	go func() {
		d.Server.ListenAndServe()
	}()
	if d.Config.StrapiEnable {
		// Enable strapi
	}
	return d.Hugo.Server(getHugoArgs(d.Config.Args))
}

var CliCommands = []string{
	"hugoPort",
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
