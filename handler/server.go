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

package handler

import (
	"github.com/ainsleyclark/errors"
	"github.com/ainsleydev/hupi/hugo"
	"net/http"
)

// Server TODO - handles the webhook request.
type Server struct {
	hugo hugo.Hugo
}

func (s Server) ListenAndServe(port string) error {
	const op = "Server.ListenAndServe"
	http.HandleFunc("/rebuild", s.Handle)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return errors.NewInternal(err, "Failed to start server", op)
	}
	return nil
}
