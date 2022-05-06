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

package pkg_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/clavinjune/bjora-project-golang/pkg/enum"

	"github.com/bwmarrin/snowflake"
	"github.com/clavinjune/bjora-project-golang/internal/util"
	"github.com/clavinjune/bjora-project-golang/pkg"
	"github.com/stretchr/testify/require"
)

func TestUserSpec_ToEntity(t *testing.T) {
	t.Parallel()
	now := time.Now()

	s := &pkg.UserSpec{
		ID:             snowflake.ID(1),
		Username:       "TestUserSpec_ToEntity",
		Email:          "TestUserSpec_ToEntity",
		Password:       "TestUserSpec_ToEntity",
		Gender:         enum.GenderMale,
		Address:        "TestUserSpec_ToEntity",
		Birthday:       now,
		CreatedAt:      now,
		CreatedBy:      "TestUserSpec_ToEntity",
		LastModifiedAt: now,
		LastModifiedBy: "TestUserSpec_ToEntity",
		IsActive:       true,
	}

	e := s.ToEntity()

	r := require.New(t)
	helperCompareUserSpecEntity(t, r, e, s)
}

func TestUserSpecFromEntity(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		t.Parallel()
		str := sql.NullString{
			String: "TestUserSpecFromEntity",
			Valid:  true,
		}

		e := &pkg.UserEntity{
			Entity: pkg.NewEntity(),
			ID: sql.NullInt64{
				Int64: int64(1),
				Valid: true,
			},
			Username: str,
			Email:    str,
			Password: str,
			Gender: sql.NullString{
				String: enum.GenderMale.String(),
				Valid:  true,
			},
			Address: str,
			Birthday: sql.NullString{
				String: "2006-12-31",
				Valid:  true,
			},
		}

		r := require.New(t)
		s, err := pkg.UserSpecFromEntity(e)

		r.NoError(err)
		helperCompareUserSpecEntity(t, r, e, s)
	})

	t.Run("birthday error", func(t *testing.T) {
		t.Parallel()
		e := &pkg.UserEntity{
			Birthday: sql.NullString{
				String: "2006-13-31",
				Valid:  true,
			},
		}

		r := require.New(t)
		_, err := pkg.UserSpecFromEntity(e)
		r.Error(err)
	})
}

func helperCompareUserSpecEntity(t *testing.T, r *require.Assertions, e *pkg.UserEntity, s *pkg.UserSpec) {
	t.Helper()

	r.True(e.ID.Valid)
	r.Equal(s.ID.Int64(), e.ID.Int64)

	r.True(e.Username.Valid)
	r.Equal(s.Username, e.Username.String)

	r.True(e.Password.Valid)
	r.Equal(s.Password, e.Password.String)

	r.True(e.Gender.Valid)
	r.Equal(s.Gender.String(), e.Gender.String)

	r.True(e.Address.Valid)
	r.Equal(s.Address, e.Address.String)

	r.True(e.Birthday.Valid)
	r.Equal(util.BirthdayFromTime(s.Birthday), e.Birthday.String)

	r.True(e.CreatedAt.Valid)
	r.Equal(s.CreatedAt, e.CreatedAt.Time)

	r.True(e.CreatedBy.Valid)
	r.Equal(s.CreatedBy, e.CreatedBy.String)

	r.True(e.LastModifiedAt.Valid)
	r.Equal(s.LastModifiedAt, e.LastModifiedAt.Time)

	r.True(e.LastModifiedBy.Valid)
	r.Equal(s.LastModifiedBy, e.LastModifiedBy.String)

	r.True(e.IsActive.Valid)
	r.Equal(s.IsActive, e.IsActive.Bool)
}
