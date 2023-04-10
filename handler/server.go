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
	"github.com/ainsleydev/hupi/hugo"
	"github.com/ainsleydev/hupi/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

// Server handles the webhook request from Strapi.
type Server struct {
	Addr string
	Hugo hugo.Hugo
}

// ListenAndServe listens on the port specified and then calls
// Serve with handler to handle requests on incoming connections.
func (s Server) ListenAndServe() {
	const op = "Server.ListenAndServe"

	// Bootstrap server.
	e := echo.New()
	e.Any("/rebuild", s.HandleStrapiWebhook)
	e.HideBanner = true
	e.HidePort = true
	e.Logger.SetLevel(log.OFF)

	// Start server.
	go func() {
		if err := e.Start(":" + s.Addr); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Error starting server: " + err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, os.Interrupt)
	//<-quit
	//if err := e.Shutdown(context.Background()); err != nil {
	//	logger.Fatal(errors.NewInternal(err, "Error shutting down server", op))
	//}
}
