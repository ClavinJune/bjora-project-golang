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

package user

import (
	"context"
	"fmt"

	"github.com/clavinjune/bjora-project-golang/pkg"
)

type service struct {
	repo pkg.UserRepository
}

func (s service) Store(ctx context.Context, user *pkg.User) (*pkg.User, error) {
	panic("implement me")
}

func (s service) FetchByEmail(ctx context.Context, email string) (*pkg.User, error) {
	e, err := s.repo.FetchByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("svc.FetchByEmail: %w", err)
	}

	return &pkg.User{
		ID:                e.ID,
		Username:          e.Username,
		Email:             e.Email,
		Gender:            e.Gender,
		Address:           e.Address,
		ProfilePictureURL: e.ProfilePictureURL,
		Birthday:          e.Birthday,
	}, nil
}
