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
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clavinjune/bjora-project-golang/pkg"
)

func TestGenderFrom(t *testing.T) {
	tt := []struct {
		name string
		in   string
		want pkg.Gender
	}{
		{
			name: "male",
			in:   " male ",
			want: pkg.GenderMale,
		},
		{
			name: "female",
			in:   " female ",
			want: pkg.GenderFemale,
		},
		{
			name: "undefined",
			in:   "asdb",
			want: pkg.GenderUndefined,
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			got := pkg.GenderFrom(tc.in)

			r.Equal(tc.want, got)
		})
	}
}

func TestGender_String(t *testing.T) {
	tt := []struct {
		name string
		in   pkg.Gender
		want string
	}{
		{
			name: "male",
			in:   pkg.GenderMale,
			want: "Male",
		},
		{
			name: "female",
			in:   pkg.GenderFemale,
			want: "Female",
		},
		{
			name: "undefined",
			in:   pkg.GenderUndefined,
			want: "Undefined",
		},
		{
			name: "random",
			in:   pkg.Gender(16),
			want: "Gender(16)",
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			got := tc.in.String()

			r.Equal(tc.want, got)
		})
	}
}
