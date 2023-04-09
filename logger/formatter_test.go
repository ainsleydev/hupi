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

package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLocalFormatter_Format(t *testing.T) {
	const prefix = "prefix"
	// now := time.Now()

	tt := map[string]struct {
		input logrus.Entry
		want  string
	}{
		"Debug": {
			logrus.Entry{
				Level: logrus.DebugLevel,
			},
			"DEBUG",
		},
		"Warn": {
			logrus.Entry{
				Level: logrus.WarnLevel,
			},
			"WARN",
		},
		"Error": {
			logrus.Entry{
				Level: logrus.ErrorLevel,
			},
			"ERROR",
		},
		"Default": {
			logrus.Entry{
				Level: logrus.InfoLevel,
			},
			"INFO",
		},
		"Fields": {
			logrus.Entry{
				Data: logrus.Fields{
					"fields": logrus.Fields{
						"key": "value",
					},
				},
			},
			"value",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			l := localFormatter{Prefix: prefix}
			got, err := l.Format(&test.input)
			assert.NoError(t, err)
			assert.Contains(t, string(got), test.want)
		})
	}
}
