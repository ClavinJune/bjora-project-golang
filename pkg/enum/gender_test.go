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

package enum_test

import (
	"testing"

	"github.com/clavinjune/bjora-project-golang/pkg/enum"

	"github.com/stretchr/testify/require"
)

func TestGenderFrom(t *testing.T) {
	tt := []struct {
		name     string
		in       string
		expected enum.Gender
	}{
		{
			name:     "male",
			in:       " male ",
			expected: enum.GenderMale,
		},
		{
			name:     "female",
			in:       " female ",
			expected: enum.GenderFemale,
		},
		{
			name:     "undefined",
			in:       "asdb",
			expected: enum.GenderUndefined,
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			actual := enum.GenderFrom(tc.in)

			r.Equal(tc.expected, actual)
		})
	}
}

func TestGender_String(t *testing.T) {
	tt := []struct {
		name     string
		in       enum.Gender
		expected string
	}{
		{
			name:     "male",
			in:       enum.GenderMale,
			expected: "Male",
		},
		{
			name:     "female",
			in:       enum.GenderFemale,
			expected: "Female",
		},
		{
			name:     "undefined",
			in:       enum.GenderUndefined,
			expected: "Undefined",
		},
		{
			name:     "random",
			in:       enum.Gender(16),
			expected: "Gender(16)",
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			actual := tc.in.String()

			r.Equal(tc.expected, actual)
		})
	}
}

func FuzzGenderFrom(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		t.Parallel()
		out := enum.GenderFrom(s).String()
		require.New(t).Equal(enum.GenderUndefined.String(), out)
	})
}
