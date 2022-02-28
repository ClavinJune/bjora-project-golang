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
	"database/sql"
	"fmt"
	"time"

	"github.com/clavinjune/bjora-project-golang/pkg/closerutil"

	"github.com/clavinjune/bjora-project-golang/pkg"
)

const (
	queryFetch        string = `SELECT * FROM bjora.users`
	queryFetchByEmail string = queryFetch + ` WHERE email = $1 and is_active = true`
	queryStore        string = `INSERT INTO bjora.users
(id, username, email, password, gender, address, profile_picture_url, birthday, role)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *`
)

type postgres struct {
	db *sql.DB
}

func (p *postgres) Fetch(ctx context.Context) ([]*pkg.UserEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := p.db.QueryContext(ctx, queryFetch)

	if err != nil {
		return nil, err
	}

	defer closerutil.Close(rows)

	return p.scanUsers(rows)
}

func (p *postgres) FetchByEmail(ctx context.Context, email string) (*pkg.UserEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := p.db.QueryContext(ctx, queryFetchByEmail, email)

	if err != nil {
		return nil, err
	}

	defer closerutil.Close(rows)
	if rows.Next() {
		return p.scanUser(rows)
	}

	return nil, sql.ErrNoRows
}

func (p *postgres) Store(ctx context.Context, entity *pkg.UserEntity) (*pkg.UserEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := p.db.QueryContext(ctx, queryStore,
		entity.ID, entity.Username, entity.Email, entity.Password, entity.Gender, entity.Address,
		entity.ProfilePictureURL, entity.Birthday, entity.Role)

	if err != nil {
		return nil, err
	}

	defer closerutil.Close(rows)
	if rows.Next() {
		return p.scanUser(rows)
	}

	return nil, fmt.Errorf("%w: failed to store entity", sql.ErrNoRows)
}

func (p *postgres) scanUsers(rows *sql.Rows) ([]*pkg.UserEntity, error) {
	result := make([]*pkg.UserEntity, 0)

	for rows.Next() {
		e, err := p.scanUser(rows)
		if err != nil {
			return nil, err
		}

		result = append(result, e)
	}

	return result, nil
}

func (postgres) scanUser(rows *sql.Rows) (*pkg.UserEntity, error) {
	var e pkg.UserEntity

	err := rows.Scan(&e.ID, &e.Username, &e.Email, &e.Password,
		&e.Gender, &e.Address, &e.ProfilePictureURL, &e.Birthday, &e.Role)

	if err != nil {
		return nil, err
	}

	return &e, nil
}
