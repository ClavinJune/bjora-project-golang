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
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/clavinjune/bjora-project-golang/pkg"
	"github.com/clavinjune/bjora-project-golang/pkg/mocks"
	"github.com/clavinjune/bjora-project-golang/user"
	"github.com/stretchr/testify/mock"
)

func TestHandler_FetchByEmail(t *testing.T) {
	m := new(mocks.UserService)
	m.EXPECT().
		FetchByEmail(mock.Anything, "example").
		Return(&pkg.User{ID: "1"}, nil)
	m.EXPECT().
		FetchByEmail(mock.Anything, "notexists").
		Return(nil, sql.ErrNoRows)
	m.EXPECT().
		FetchByEmail(mock.Anything, mock.Anything).
		Return(nil, nil)

	app := fiber.New()
	user.ProvideHandler(m).ApplyRoute(app.Group("/user"))

	tt := []struct {
		Name       string
		Endpoint   string
		ExpectCode int
	}{
		{
			Name:       "without email query",
			Endpoint:   "/user/",
			ExpectCode: http.StatusNotFound,
		},
		{
			Name:       "with not exists email",
			Endpoint:   "/user/?email=notexists",
			ExpectCode: http.StatusNotFound,
		},
		{
			Name:       "with exists email",
			Endpoint:   "/user/?email=example",
			ExpectCode: http.StatusOK,
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			resp, err := app.Test(httptest.NewRequest(http.MethodGet, tc.Endpoint, nil))
			assert.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, tc.ExpectCode, resp.StatusCode)
		})
	}
}
