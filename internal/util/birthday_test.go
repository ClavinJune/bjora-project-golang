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
	"testing"
	"time"

	"github.com/clavinjune/bjora-project-golang/internal/util"

	"github.com/stretchr/testify/require"
)

func TestBirthdayFromTime(t *testing.T) {
	tt := []struct {
		name     string
		in       time.Time
		expected string
	}{
		{
			name:     "zero",
			in:       time.Time{},
			expected: "0001-01-01",
		},
		{
			name:     "not zero",
			in:       time.Date(2006, time.February, 1, 0, 0, 0, 0, time.UTC),
			expected: "2006-02-01",
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			actual := util.BirthdayFromTime(tc.in)
			r.Equal(tc.expected, actual)
		})
	}
}

func TestBirthdayFromStr(t *testing.T) {
	tt := []struct {
		name        string
		in          string
		expectTime  time.Time
		isExpectErr bool
	}{
		{
			name:        "empty string",
			in:          "",
			expectTime:  time.Time{},
			isExpectErr: true,
		},
		{
			name:        "valid layout",
			in:          "2006-02-01",
			expectTime:  time.Date(2006, time.February, 1, 0, 0, 0, 0, time.UTC),
			isExpectErr: false,
		},
		{
			name:        "invalid layout",
			in:          "20060201",
			expectTime:  time.Time{},
			isExpectErr: true,
		},
		{
			name:        "invalid date",
			in:          "2006-13-01",
			expectTime:  time.Time{},
			isExpectErr: true,
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			actualTime, actualErr := util.BirthdayFromStr(tc.in)
			r.Equal(tc.expectTime, actualTime)

			if tc.isExpectErr {
				r.Error(actualErr)

				{
					var err *time.ParseError
					r.ErrorAs(actualErr, &err)
					r.ErrorIs(actualErr, err)
				}

				{
					var err *util.Error
					r.ErrorAs(actualErr, &err)
					r.ErrorIs(actualErr, err)
				}
			} else {
				r.NoError(actualErr)
			}
		})
	}
}
