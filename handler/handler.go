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
	"github.com/ainsleydev/hupi/logger"
	"github.com/ainsleydev/hupi/strapi"
	"github.com/labstack/echo/v4"
)

// HandleStrapiWebhook handles the incoming webhook from Strapi and rebuilds
// Hugo when it encounters any change within Strapi.
func (s Server) HandleStrapiWebhook(ctx echo.Context) error {
	const op = "Server.HandleStrapiWebhook"

	logger.Debug("Received Strapi request")

	var entry = strapi.Entry{}
	err := ctx.Bind(&entry)
	if err != nil {
		return errors.NewInternal(err, "Failed to bind to response body", op)
	}

	logger.Info("Rebuilding Hugo...")

	err = s.Hugo.Rebuild()
	if err != nil {
		return err
	}

	logger.Info("Rebuild complete")

	return nil
}
