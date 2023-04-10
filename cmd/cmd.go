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

package cmd

import (
	"github.com/ainsleydev/hupi/logger"
	"github.com/ainsleydev/hupi/version"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func Run() {
	app := &cli.App{
		Name:                 "hupi",
		Usage:                "A cross platform CLI command line tool for Hugo and Strapi",
		DefaultCommand:       "develop",
		Version:              version.Version,
		Flags:                nil,
		EnableBashCompletion: true,
		Before: func(ctx *cli.Context) error {
			logger.Bootstrap("hupi")
			return nil
		},
		Authors: []*cli.Author{
			{
				Name:  "ainsley.dev",
				Email: "hello@ainsley.dev",
			},
		},
		Commands: []*cli.Command{
			versionCommand,
			developCommand,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
