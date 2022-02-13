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

	"github.com/clavinjune/bjora-project-golang/pkg"
)

type repository struct {
	db *sql.DB
}

func (r *repository) Store(ctx context.Context, entity *pkg.UserEntity) (*pkg.UserEntity, error) {
	panic("implement me")
}

func (r *repository) FetchByEmail(ctx context.Context, email string) (*pkg.UserEntity, error) {
	panic("implement me")
}
