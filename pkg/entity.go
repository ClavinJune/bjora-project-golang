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
	"database/sql"
	"os"
	"time"
)

var (
	defaultMetadataString = sql.NullString{
		String: os.Getenv("APP_NAME"),
		Valid:  true,
	}
)

// Entity provides metadata columns
type Entity struct {
	CreatedAt      sql.NullTime
	CreatedBy      sql.NullString
	LastModifiedAt sql.NullTime
	LastModifiedBy sql.NullString
	DeletedAt      sql.NullTime
	DeletedBy      sql.NullString
}

// IsDeleted check whether entity has been deleted or not
func (e *Entity) IsDeleted() bool {
	return e.DeletedAt.Valid &&
		!e.DeletedAt.Time.IsZero() &&
		e.DeletedBy.Valid &&
		e.DeletedBy.String != ""
}

// NewEntity creates new metadata
func NewEntity() *Entity {
	defaultMetadataTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	return &Entity{
		CreatedAt:      defaultMetadataTime,
		CreatedBy:      defaultMetadataString,
		LastModifiedAt: defaultMetadataTime,
		LastModifiedBy: defaultMetadataString,
		DeletedAt: sql.NullTime{
			Valid: false,
		},
		DeletedBy: sql.NullString{
			Valid: false,
		},
	}
}
