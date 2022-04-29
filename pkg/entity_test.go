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
	"os"
	"testing"
	"time"

	"github.com/clavinjune/bjora-project-golang/pkg"
	"github.com/stretchr/testify/require"
)

func TestNewEntity(t *testing.T) {
	appName := os.Getenv("APP_NAME")
	r := require.New(t)
	e := pkg.NewEntity()

	r.True(e.CreatedAt.Valid)
	r.False(e.CreatedAt.Time.IsZero())
	r.True(e.CreatedBy.Valid)
	r.Equal(appName, e.CreatedBy.String)

	r.True(e.LastModifiedAt.Valid)
	r.False(e.LastModifiedAt.Time.IsZero())
	r.True(e.LastModifiedBy.Valid)
	r.Equal(appName, e.LastModifiedBy.String)

	r.False(e.DeletedAt.Valid)
	r.True(e.DeletedAt.Time.IsZero())
	r.False(e.DeletedBy.Valid)
	r.Empty(e.DeletedBy.String)
	r.False(e.IsDeleted())
}

func TestEntity_IsDeleted(t *testing.T) {
	tt := []struct {
		name      string
		deletedAt sql.NullTime
		deletedBy sql.NullString
		want      bool
	}{
		{
			name:      "deletedAt empty deletedBy empty",
			deletedAt: sql.NullTime{},
			deletedBy: sql.NullString{},
			want:      false,
		},
		{
			name: "deletedAt time not zero but not valid",
			deletedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: false,
			},
			deletedBy: sql.NullString{},
			want:      false,
		},
		{
			name: "deletedAt time zero but valid",
			deletedAt: sql.NullTime{
				Time:  time.Time{},
				Valid: true,
			},
			deletedBy: sql.NullString{},
			want:      false,
		},
		{
			name: "deletedAt time not zero and valid but deletedBy empty",
			deletedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			deletedBy: sql.NullString{},
			want:      false,
		},
		{
			name: "deletedBy string not empty but not valid",
			deletedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			deletedBy: sql.NullString{
				String: "testing",
				Valid:  false,
			},
			want: false,
		},
		{
			name: "deletedBy string empty but valid",
			deletedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			deletedBy: sql.NullString{
				String: "",
				Valid:  true,
			},
			want: false,
		},
		{
			name:      "deletedBy string not empty and valid but deletedAt empty",
			deletedAt: sql.NullTime{},
			deletedBy: sql.NullString{
				String: "testing",
				Valid:  true,
			},
			want: false,
		},
		{
			name: "deletedBy not empty and deletedAt not empty",
			deletedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			deletedBy: sql.NullString{
				String: "testing",
				Valid:  true,
			},
			want: true,
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			e := pkg.NewEntity()
			e.DeletedAt = tc.deletedAt
			e.DeletedBy = tc.deletedBy

			r.Equal(tc.want, e.IsDeleted())
		})
	}
}
