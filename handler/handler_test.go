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
	"github.com/ainsleydev/hupi/internal/mocks"
	"github.com/ainsleydev/hupi/logger"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer_Handle(t *testing.T) {
	logger.SetOutput(io.Discard)

	tt := map[string]struct {
		payload any
		mock    func(m *mocks.Hugo)
		want    any
	}{
		"Bind Error": {
			payload: "wrong",
			want:    "Failed to bind to response body",
		},
		"Rebuild Error": {
			mock: func(m *mocks.Hugo) {
				m.On("Rebuild").Return(errors.New("rebuild error"))
			},
			want: "rebuild error",
		},
		"OK": {
			mock: func(m *mocks.Hugo) {
				m.On("Rebuild").Return(nil)
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(nil)
			defer server.Close()

			hugo := mocks.NewHugo(t)
			if test.mock != nil {
				test.mock(hugo)
			}

			s := Server{
				Hugo: hugo,
			}

			b, err := json.Marshal(test.payload)
			require.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/rebuild", strings.NewReader(string(b)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := echo.New().NewContext(req, rec)

			err = s.HandleStrapiWebhook(c)
			if err != nil {
				assert.ErrorContains(t, err, test.want.(string))
				return
			}
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}
