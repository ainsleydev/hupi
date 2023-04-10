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

package hugo

import (
	"github.com/ainsleyclark/errors"
	"github.com/ainsleydev/hupi/execute"
	"path/filepath"
)

type (
	// Hugo - TODO
	Hugo interface {
		Server(args []string) error
		Build(args []string) error
		Rebuild() error
	}
	Client struct {
		BuildDirectory string
	}
)

func (c Client) Server(args []string) error {
	const op = "Hugo.Server"

	out := &hugoWriter{}
	err := execute.Task{
		Command: "hugo",
		Args:    append(args, "server"),
		StdOut:  out,
		StdErr:  out,
	}.Run()
	if err != nil {
		return errors.NewInternal(err, "Failed to start Hugo server", op)
	}

	return nil
}

func (c Client) Build(args []string) error {
	const op = "Hugo.Build"

	err := execute.Task{
		Command: "hugo",
		Args:    args,
	}.Run()
	if err != nil {
		return errors.NewInternal(err, "Failed build Hugo", op)
	}

	return nil
}

// Rebuild rebuilds the Hugo site.
func (c Client) Rebuild() error {
	const op = "Hugo.Rebuild"

	err := c.clean()
	if err != nil {
		return errors.NewInternal(err, "Failed to clean public folder", op)
	}

	err = execute.Task{
		Command: "hugo",
		Dir:     c.BuildDirectory,
		Args: []string{
			"--cleanDestinationDir",
		},
		StdOut: nil,
		StdErr: nil,
	}.Run()
	if err != nil {
		return errors.NewInternal(err, "Failed to rebuild Hugo", op)
	}

	return nil
}

func (c Client) clean() error {
	return execute.Task{
		Command: "rm",
		Args: []string{
			"rf",
			filepath.Join(c.BuildDirectory, "public"),
		},
		StdOut: nil,
		StdErr: nil,
	}.Run()
}
