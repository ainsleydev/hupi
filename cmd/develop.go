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
	"github.com/ainsleydev/hupi/service"
	"github.com/urfave/cli/v2"
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
			Name:     "hugoPort",
			Usage:    "Port on which the Hugo server will listen on.",
			Value:    "5001",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "hugoBuildDirectory",
			Value:    "prebuild",
			Usage:    "Determines the directory in which Hupi will build content.",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "strapiEnable",
			Value:    false,
			Usage:    "When enabled, Hupi will run Strapi develop.",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "strapiNoBuild",
			Value:    false,
			Usage:    "Starts Strapi with autoReload enabled and skip the administration panel build process.",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "strapiWatchAdmin",
			Value:    false,
			Usage:    "Starts Strapi with autoReload enabled and the front-end development server. It allows you to customize the administration panel.",
			Required: false,
		},
	},
	Action: func(ctx *cli.Context) error {
		service.NewDevelop(service.DevelopConfig{
			Args:               ctx.Args().Slice(),
			HugoPort:           ctx.String("hugoPort"),
			HugoBuildDirectory: ctx.String("hugoBuildDirectory"),
			StrapiEnable:       ctx.Bool("strapiEnable"),
			StrapiNoBuild:      ctx.Bool("strapiNoBuild"),
			StrapiWatchAdmin:   ctx.Bool("strapiWatchAdmin"),
		}).Develop()
		return nil
	},
}
