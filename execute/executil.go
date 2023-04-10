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

package execute

import (
	"io"
	"os/exec"
)

type (
	Runner interface {
		Run() error
	}
	Task struct {
		Command string
		Args    []string
		Dir     string
		StdOut  io.Writer
		StdErr  io.Writer
	}
)

func (t Task) Run() error {
	cmd := exec.Command(t.Command, t.Args...)
	cmd.Dir = t.Dir

	r, w := io.Pipe()
	cmd.Stdin = r
	cmd.Stdout = io.MultiWriter(w, t.StdOut)
	cmd.Stderr = io.MultiWriter(w, t.StdErr)

	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}

func CommandExists(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}
