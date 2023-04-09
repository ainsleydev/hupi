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

package strapi

import "time"

type (
	Event struct {
		Event     string    `json:"event"`
		CreatedAt time.Time `json:"createdAt"`
		Model     string    `json:"model"`
		Entry     Entry     `json:"entry,omitempty"`
		Media     Media     `json:"media,omitempty"`
	}
	Entry struct {
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}
	Media struct {
		Media struct {
			ID        int       `json:"id"`
			Name      string    `json:"name"`
			Hash      string    `json:"hash"`
			Sha256    string    `json:"sha256"`
			Ext       string    `json:"ext"`
			Mime      string    `json:"mime"`
			Size      float64   `json:"size"`
			URL       string    `json:"url"`
			CreatedAt time.Time `json:"createdAt"`
			UpdatedAt time.Time `json:"updatedAt"`
		} `json:"media"`
	}
)

func (e Event) Fields() map[string]any {
	return map[string]any{
		"event": e.Event,
		"model": e.Model,
	}
}
