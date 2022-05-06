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
	"database/sql"
	"time"

	"github.com/clavinjune/bjora-project-golang/pkg/enum"

	"github.com/clavinjune/bjora-project-golang/internal/util"

	"github.com/bwmarrin/snowflake"
)

type (
	// UserEntity is used to store user data to database
	UserEntity struct {
		*Entity
		ID       sql.NullInt64  `db:"id"`
		Username sql.NullString `db:"username"`
		Email    sql.NullString `db:"email"`
		Password sql.NullString `db:"password"`
		Gender   sql.NullString `db:"gender"`
		Address  sql.NullString `db:"address"`
		Birthday sql.NullString `db:"birthday"`
	}

	// UserSpec is a specification for User
	UserSpec struct {
		ID             snowflake.ID
		Username       string
		Email          string
		Password       string
		Gender         enum.Gender
		Address        string
		Birthday       time.Time
		CreatedAt      time.Time
		CreatedBy      string
		LastModifiedAt time.Time
		LastModifiedBy string
		IsActive       bool
	}

	// UserRepository is an interface to interact with repository
	UserRepository interface {
		Store(ctx context.Context, e *UserEntity) (*UserEntity, error)
	}

	// UserService is an interface to handle business logic
	UserService interface {
		Store(ctx context.Context, u *UserSpec) (*UserSpec, error)
	}
)

// ToEntity converts UserSpec into *UserEntity
func (u *UserSpec) ToEntity() *UserEntity {
	return &UserEntity{
		Entity: &Entity{
			CreatedAt: sql.NullTime{
				Time:  u.CreatedAt,
				Valid: true,
			},
			CreatedBy: sql.NullString{
				String: u.CreatedBy,
				Valid:  true,
			},
			LastModifiedAt: sql.NullTime{
				Time:  u.LastModifiedAt,
				Valid: true,
			},
			LastModifiedBy: sql.NullString{
				String: u.LastModifiedBy,
				Valid:  true,
			},
			IsActive: sql.NullBool{
				Bool:  u.IsActive,
				Valid: true,
			},
		},
		ID: sql.NullInt64{
			Int64: u.ID.Int64(),
			Valid: true,
		},
		Username: sql.NullString{
			String: u.Username,
			Valid:  true,
		},
		Email: sql.NullString{
			String: u.Email,
			Valid:  true,
		},
		Password: sql.NullString{
			String: u.Password,
			Valid:  true,
		},
		Gender: sql.NullString{
			String: u.Gender.String(),
			Valid:  true,
		},
		Address: sql.NullString{
			String: u.Address,
			Valid:  true,
		},
		Birthday: sql.NullString{
			String: util.BirthdayFromTime(u.Birthday),
			Valid:  true,
		},
	}
}

// UserSpecFromEntity converts *UserEntity into *UserSpec
func UserSpecFromEntity(e *UserEntity) (*UserSpec, error) {
	if e == nil {
		return nil, util.NewErrorFromMsg("user entity is nil")
	}

	birthday, err := util.BirthdayFromStr(e.Birthday.String)
	if err != nil {
		return nil, util.WrapError(err)
	}

	return &UserSpec{
		ID:             snowflake.ID(e.ID.Int64),
		Username:       e.Username.String,
		Email:          e.Email.String,
		Password:       e.Password.String,
		Gender:         enum.GenderFrom(e.Gender.String),
		Address:        e.Address.String,
		Birthday:       birthday,
		CreatedAt:      e.CreatedAt.Time,
		CreatedBy:      e.CreatedBy.String,
		LastModifiedAt: e.LastModifiedAt.Time,
		LastModifiedBy: e.LastModifiedBy.String,
		IsActive:       e.IsActive.Bool,
	}, nil
}

// UserSpecFromEntities converts []*UserEntity into []*UserSpec
func UserSpecFromEntities(entities []*UserEntity) ([]*UserSpec, error) {
	n := len(entities)
	specs := make([]*UserSpec, n)
	for i := range entities {
		e := entities[i]
		spec, err := UserSpecFromEntity(e)
		if err != nil {
			return nil, util.WrapError(err)
		}

		specs[i] = spec
	}

	return specs, nil
}
