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

package user_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/clavinjune/bjora-project-golang/internal/user"
	"github.com/clavinjune/bjora-project-golang/internal/util"
	"github.com/clavinjune/bjora-project-golang/pkg/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/clavinjune/bjora-project-golang/pkg"
)

func helperCreateEntity(t *testing.T) *pkg.UserEntity {
	t.Helper()

	ns := sql.NullString{
		String: "helperCreateEntity",
		Valid:  true,
	}

	return &pkg.UserEntity{
		Entity: pkg.NewEntity(),
		ID: sql.NullInt64{
			Int64: 1,
			Valid: true,
		},
		Username: ns,
		Email:    ns,
		Password: ns,
		Gender:   ns,
		Address:  ns,
		Birthday: sql.NullString{
			String: "2006-12-31",
			Valid:  true,
		},
	}
}

func TestService_Store(t *testing.T) {
	snowNode := util.ProvideSnowflake()

	t.Run("error from repository", func(t *testing.T) {
		r := require.New(t)
		repo := new(mocks.UserRepository)
		defer repo.AssertExpectations(t)

		expectedError := util.NewErrorFromMsg("mocked error from repo")

		repo.EXPECT().
			Store(mock.Anything, mock.Anything).
			Return(nil, expectedError)

		stored, err := user.ProvideService(repo, snowNode).
			Store(context.Background(), &pkg.UserSpec{})

		r.Error(err)
		r.ErrorIs(err, expectedError)
		{
			var targetErr *util.Error
			r.ErrorAs(err, &targetErr)
		}
		r.Nil(stored)
	})

	t.Run("error from parsing birthday", func(t *testing.T) {
		r := require.New(t)
		repo := new(mocks.UserRepository)
		defer repo.AssertExpectations(t)

		entity := helperCreateEntity(t)
		entity.Birthday.String = "2001123-2312313123-12"

		repo.EXPECT().
			Store(mock.Anything, mock.Anything).
			Return(entity, nil)

		stored, err := user.ProvideService(repo, snowNode).
			Store(context.Background(), &pkg.UserSpec{})

		r.Error(err)
		{
			var targetErr *util.Error
			r.ErrorAs(err, &targetErr)
		}
		{
			var targetErr *time.ParseError
			r.ErrorAs(err, &targetErr)
		}
		r.Nil(stored)
	})

	t.Run("no error", func(t *testing.T) {
		r := require.New(t)
		repo := new(mocks.UserRepository)
		defer repo.AssertExpectations(t)

		entity := helperCreateEntity(t)
		repo.EXPECT().
			Store(mock.Anything, mock.Anything).
			Return(entity, nil)

		stored, err := user.ProvideService(repo, snowNode).
			Store(context.Background(), &pkg.UserSpec{})

		r.NoError(err)
		s, err := pkg.UserSpecFromEntity(entity)
		r.NoError(err)
		r.Equal(s, stored)
	})
}
