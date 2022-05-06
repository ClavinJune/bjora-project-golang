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
	"fmt"
	"os"
	"testing"

	"github.com/clavinjune/bjora-project-golang/internal/user"
	"github.com/clavinjune/bjora-project-golang/internal/util"
	"github.com/clavinjune/bjora-project-golang/pkg"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestRepository_Store(t *testing.T) {
	r := require.New(t)
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
	))
	r.NoError(err)

	id := util.ProvideSnowflake().Generate()
	sqlns := sql.NullString{
		String: id.String(),
		Valid:  true,
	}
	e := &pkg.UserEntity{
		Entity: pkg.NewEntity(),
		ID: sql.NullInt64{
			Int64: id.Int64(),
			Valid: true,
		},
		Username: sqlns,
		Email:    sqlns,
		Password: sqlns,
		Gender:   sqlns,
		Address:  sqlns,
		Birthday: sqlns,
	}
	result, err := user.ProvideRepository(db).Store(context.Background(), e)
	r.NoError(err)
	result.Entity = e.Entity
	r.Equal(e, result)
}
