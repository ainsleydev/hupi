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
	"encoding/json"
	"github.com/ainsleyclark/errors"
	"github.com/ainsleydev/hupi/logger"
	"github.com/ainsleydev/hupi/strapi"
	"io"
	"net/http"
)

// Handle TODO - handles the request.
func (s Server) Handle(w http.ResponseWriter, r *http.Request) {
	const op = "TODO.Handle"

	defer r.Body.Close()
	buf, err := io.ReadAll(r.Body)
	if err != nil {
		logger.WithError(errors.NewInternal(err, "Failed to read response body", op)).Error()
		return
	}

	var entry = strapi.Entry{}
	err = json.Unmarshal(buf, &entry)
	if err != nil {
		logger.WithError(errors.NewInternal(err, "Failed to unmarshal response body", op)).Error()
		return
	}

	logger.WithField("body", string(buf)).Trace("Received request")
	logger.Info("Rebuilding Hugo...")

	err = s.hugo.Rebuild()
	if err != nil {
		logger.WithError(err).Error()
	}

	logger.Info("Rebuild complete")
}
