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
	"strings"
	"time"
)

// Entity provides metadata columns
type Entity struct {
	CreatedAt      sql.NullTime   `db:"created_at"`
	CreatedBy      sql.NullString `db:"created_by"`
	LastModifiedAt sql.NullTime   `db:"last_modified_at"`
	LastModifiedBy sql.NullString `db:"last_modified_by"`
	IsActive       sql.NullBool   `db:"is_active"`
}

// Update updates LastModifiedAt and LastModifiedBy
func (e *Entity) Update(by string) {
	t, s := createMetadataNow(by)
	e.LastModifiedAt = t
	e.LastModifiedBy = s
}

// Delete updates LastModifiedAt, LastModifiedBy, and IsActive
func (e *Entity) Delete(by string) {
	e.Update(by)
	e.IsActive = sql.NullBool{
		Bool:  false,
		Valid: true,
	}
}

// NewEntity creates new metadata
func NewEntity() *Entity {
	t, s := createMetadataNow(os.Getenv("APP_NAME"))

	return &Entity{
		CreatedAt:      t,
		CreatedBy:      s,
		LastModifiedAt: t,
		LastModifiedBy: s,
		IsActive: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
	}
}

// createMetadataNow produces valid sql.NullTime(now) and sql.NullString(by)
func createMetadataNow(by string) (sql.NullTime, sql.NullString) {
	t := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	s := sql.NullString{
		String: strings.TrimSpace(by),
		Valid:  true,
	}

	return t, s
}
