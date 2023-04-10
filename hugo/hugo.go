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
	"bufio"
	"fmt"
	"github.com/ainsleyclark/errors"
	"os/exec"
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

	cmd := exec.Command("hugo", append(args, "server")...)
	stdOut, _ := cmd.StdoutPipe()

	err := cmd.Start()
	if err != nil {
		return errors.NewInvalid(err, "Failed to start Hugo server", op)
	}

	scanner := bufio.NewScanner(stdOut)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	err = cmd.Wait()
	if err != nil {
		return errors.NewInvalid(err, "Failed to start Hugo server", op)
	}

	return nil
}

func (c Client) Build(args []string) error {
	const op = "Hugo.Build"

	cmd := exec.Command("hugo", args...)
	_, err := cmd.StdoutPipe()
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
		return errors.NewInternal(err, "Could not clean public folder", op)
	}

	cmd := exec.Command("hugo", "--cleanDestinationDir")
	cmd.Dir = c.BuildDirectory
	_, err = cmd.Output()
	if err != nil {
		return errors.NewInternal(err, "Failed to rebuild Hugo", op)
	}

	return nil
}

func (c Client) clean() error {
	// TODO, we need to pass the public folder dynamically here.
	cmd := exec.Command("rm", "-rf", "./prebuild/public")
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}
