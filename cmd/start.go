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
	"fmt"
	"github.com/ainsleydev/hupi/handler"
	"github.com/ainsleydev/hupi/hugo"
	"github.com/ainsleydev/hupi/service"
	"github.com/ainsleydev/hupi/strapi"
	"github.com/urfave/cli/v2"
	"strings"
)

// developCommand is the command to start the Hugo server and
// watch for changes via the Strapi webhook.
var developCommand = &cli.Command{
	Name: "develop",
	Aliases: []string{
		"dev",
	},
	SkipFlagParsing: true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "hugo-port",
			Value:    "5000",
			Usage:    "Port on which the Hugo server will listen on.",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "hugo-build-directory",
			Value:    "prebuild",
			Usage:    "Determines the directory in which Hupi will build content.",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "strapi-enable",
			Value:    false,
			Usage:    "When enabled, Hupi will run Strapi develop.",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "strapi-no-build",
			Value:    false,
			Usage:    "Starts Strapi with autoReload enabled and skip the administration panel build process.",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "strapi-watch-admin",
			Value:    false,
			Usage:    "Starts Strapi with autoReload enabled and the front-end development server. It allows you to customize the administration panel.",
			Required: false,
		},
	},
	Action: func(ctx *cli.Context) error {
		enableStrapi := ctx.Bool("strapi-enable")
		if enableStrapi {
			_ = strapi.Client{
				NoBuild:    ctx.Bool("strapi-no-build"),
				WatchAdmin: ctx.Bool("strapi-watch-admin"),
			}
		}

		hu := hugo.Client{
			BuildDirectory: ctx.String("hugo-build-directory"),
		}

		handle := handler.Server{}
		go func() {
			err := handle.ListenAndServe(ctx.String("hugo-port"))
			if err != nil {
				return
			}
		}()

		var hugoArgs []string
		args := ctx.Args().Slice()
	Outer:
		for _, arg := range args {
			for _, cliArg := range service.CliCommands {
				if strings.Contains(arg, cliArg) {
					continue Outer
				}
				hugoArgs = append(hugoArgs, arg)
			}
		}

		err := hu.Server(hugoArgs)
		if err != nil {
			fmt.Println("gdg", err.Error())
			return err
		}

		return nil
	},
}
