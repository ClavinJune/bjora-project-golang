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

import (
	"context"
	"net/http"
)

type (
	// User defines publicly exposed user attributes
	User struct {
		ID                string `json:"id"`
		Username          string `json:"username"`
		Email             string `json:"email"`
		Gender            Gender `json:"gender"`
		Address           string `json:"address"`
		ProfilePictureURL string `json:"profile_picture"`
		Birthday          string `json:"birthday"`
	}

	// UserEntity defines database model
	UserEntity struct {
		ID                string
		Username          string
		Email             string
		Password          string
		Gender            Gender
		Address           string
		ProfilePictureURL string
		Birthday          string
	}

	// UserRepository used for accessing storage
	UserRepository interface {
		Store(ctx context.Context, entity *UserEntity) (*UserEntity, error)
		FetchByEmail(ctx context.Context, email string) (*UserEntity, error)
	}

	// UserService used for communicating with repository
	UserService interface {
		Store(ctx context.Context, user *User) (*User, error)
		FetchByEmail(ctx context.Context, email string) (*User, error)
	}

	// UserHandler used for handling HTTP Request
	UserHandler interface {
		Store() http.HandlerFunc
		FetchByEmail() http.HandlerFunc
	}
)
