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

	"github.com/clavinjune/bjora-project-golang/internal/util"

	"github.com/clavinjune/bjora-project-golang/pkg"
	"github.com/jmoiron/sqlx"
)

const (
	insertQuery string = `INSERT INTO bjora.users(id, username, email, password, gender, address, birthday)
VALUES(:id, :username, :email, :password, :gender, :address, :birthday)
RETURNING *;`
)

type repository struct {
	db *sqlx.DB
}

func (r repository) Store(ctx context.Context, e *pkg.UserEntity) (*pkg.UserEntity, error) {
	rows, err := r.db.NamedQueryContext(ctx, insertQuery, e)
	if err != nil {
		return nil, util.WrapError(err)
	}

	defer rows.Close()
	result, err := r.read(rows)
	if err != nil {
		return nil, util.WrapError(err)
	}

	return result, nil
}

func (repository) read(rows *sqlx.Rows) (*pkg.UserEntity, error) {
	if rows.Next() {
		var row pkg.UserEntity
		err := rows.StructScan(&row)
		if err != nil {
			return nil, util.WrapError(err)
		}

		return &row, nil
	}

	return nil, util.NewErrorFromMsg("rows are empty")
}
