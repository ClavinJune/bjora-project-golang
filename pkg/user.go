// Copyright 2022 ClavinJune/bjora
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pkg

import "time"

type (
	// User defines publicly exposed user attributes
	User struct {
		ID             string    `json:"id"`
		Username       string    `json:"username"`
		Email          string    `json:"email"`
		Gender         Gender    `json:"gender"`
		Address        string    `json:"address"`
		ProfilePicture string    `json:"profile_picture"`
		Birthday       time.Time `json:"birthday"`
	}

	// UserEntity defines database model
	UserEntity struct {
		ID             string
		Username       string
		Email          string
		Password       string
		Gender         string
		Address        string
		ProfilePicture string
		Birthday       string
	}
)
