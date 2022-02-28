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

	"github.com/stretchr/testify/assert"

	"github.com/clavinjune/bjora-project-golang/pkg"
	"github.com/clavinjune/bjora-project-golang/pkg/mocks"
	"github.com/clavinjune/bjora-project-golang/user"
	"github.com/stretchr/testify/mock"
)

func TestService_FetchByEmail(t *testing.T) {
	m := new(mocks.UserRepository)
	m.EXPECT().
		FetchByEmail(mock.Anything, "example").
		Return(&pkg.UserEntity{ID: "1"}, nil)
	m.EXPECT().
		FetchByEmail(mock.Anything, "notexists").
		Return(nil, sql.ErrNoRows)
	m.EXPECT().
		FetchByEmail(mock.Anything, mock.Anything).
		Return(nil, nil)

	svc := user.ProvideService(m)

	tt := []struct {
		Name  string
		Email string
		Error error
	}{
		{
			Name:  "with not exists email",
			Email: "notexists",
			Error: sql.ErrNoRows,
		},
		{
			Name:  "with exists email",
			Email: "example",
			Error: nil,
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			_, err := svc.FetchByEmail(context.Background(), tc.Email)
			assert.ErrorIs(t, err, tc.Error)
		})
	}
}
