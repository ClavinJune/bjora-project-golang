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

package util_test

import (
	"database/sql"
	"encoding/json"
	"strings"
	"testing"

	"github.com/clavinjune/bjora-project-golang/internal/util"
	"github.com/stretchr/testify/require"
)

type expectation struct {
	FuncName string
	FileLine string
	Message  string
	Caused   any
}

func TestNewError(t *testing.T) {
	t.Parallel()
	actualErr := util.WrapError(sql.ErrNoRows)

	r := require.New(t)

	var err *util.Error
	r.ErrorAs(actualErr, &err)
	r.ErrorIs(actualErr, err)
	r.ErrorIs(actualErr, sql.ErrNoRows)
	r.NotErrorIs(actualErr, sql.ErrTxDone)
}

func TestNewErrorWithMsg(t *testing.T) {
	t.Parallel()
	actualErr := util.WrapErrorWithMsg(sql.ErrNoRows, "foobar")

	r := require.New(t)

	var err *util.Error
	r.ErrorAs(actualErr, &err)
	r.ErrorIs(actualErr, err)
	r.ErrorIs(actualErr, sql.ErrNoRows)
	r.NotErrorIs(actualErr, sql.ErrTxDone)
}

func TestNewErrorFromMsg(t *testing.T) {
	t.Parallel()
	actualErr := util.NewErrorFromMsg("foobar")

	r := require.New(t)

	var err *util.Error
	r.ErrorAs(actualErr, &err)
	r.ErrorIs(actualErr, err)
	r.NotErrorIs(actualErr, sql.ErrTxDone)
}

func TestError_Error(t *testing.T) {
	tt := []struct {
		name     string
		in       error
		expected expectation
	}{
		{
			name: "without message",
			in:   util.WrapError(sql.ErrNoRows),
			expected: expectation{
				FuncName: "util_test.TestError_Error",
				FileLine: "error_test.go:",
				Caused:   "sql: no rows in result set",
			},
		},
		{
			name: "with message",
			in:   util.WrapErrorWithMsg(sql.ErrNoRows, "foobar"),
			expected: expectation{
				FuncName: "util_test.TestError_Error",
				FileLine: "error_test.go:",
				Message:  "foobar",
				Caused:   "sql: no rows in result set",
			},
		},
		{
			name: "from message",
			in:   util.NewErrorFromMsg("foobar"),
			expected: expectation{
				FuncName: "util_test.TestError_Error",
				FileLine: "error_test.go:",
				Caused:   "foobar",
			},
		},
		{
			name: "nested",
			in:   util.WrapErrorWithMsg(util.WrapError(sql.ErrNoRows), "foobar"),
			expected: expectation{
				FuncName: "util_test.TestError_Error",
				FileLine: "error_test.go:",
				Message:  "foobar",
				Caused: expectation{
					FuncName: "util_test.TestError_Error",
					FileLine: "error_test.go:",
					Caused:   "sql: no rows in result set",
				},
			},
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var actual expectation

			r := require.New(t)
			err := json.Unmarshal([]byte(tc.in.Error()), &actual)
			r.NoError(err)

			helperCompareCaused(t, r, tc.expected, actual)
		})
	}
}

func helperCompareCaused(t *testing.T, r *require.Assertions, expected, actual expectation) {
	t.Helper()

	r.Equal(expected.FuncName, actual.FuncName)
	r.True(strings.HasPrefix(actual.FileLine, expected.FileLine))
	r.Equal(expected.Message, actual.Message)

	if _, ok := expected.Caused.(string); ok {
		r.Equal(expected.Caused, actual.Caused)
		return
	}

	b, err := json.Marshal(actual.Caused.(map[string]any)["caused"])
	r.NoError(err)
	var e expectation
	r.NoError(json.Unmarshal(b, &e))
	helperCompareCaused(t, r, expected.Caused.(expectation), e)
}
