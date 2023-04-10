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
	"github.com/ainsleydev/hupi/version"
	"github.com/urfave/cli/v2"
)

// versionCommand is the command to print out the version
// number of the application.
var versionCommand = &cli.Command{
	Name: "version",
	Aliases: []string{
		"v",
	},
	Action: func(ctx *cli.Context) error {
		fmt.Printf("v%s\n", version.Version)
		return nil
	},
}
