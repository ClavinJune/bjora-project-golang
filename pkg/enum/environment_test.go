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

func TestEnvironmentFrom(t *testing.T) {
	tt := []struct {
		name     string
		in       string
		expected enum.Environment
	}{
		{
			name:     "test",
			in:       " test ",
			expected: enum.EnvironmentTest,
		},
		{
			name:     "dev",
			in:       " dev ",
			expected: enum.EnvironmentDev,
		},
		{
			name:     "stg",
			in:       " stg ",
			expected: enum.EnvironmentStg,
		},
		{
			name:     "prod",
			in:       " prod ",
			expected: enum.EnvironmentProd,
		},
		{
			name:     "undefined",
			in:       "asdb",
			expected: enum.EnvironmentUndefined,
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			actual := enum.EnvironmentFrom(tc.in)

			r.Equal(tc.expected, actual)
		})
	}
}

func TestEnvironment_String(t *testing.T) {
	tt := []struct {
		name     string
		in       enum.Environment
		expected string
	}{
		{
			name:     "test",
			in:       enum.EnvironmentTest,
			expected: "Test",
		},
		{
			name:     "dev",
			in:       enum.EnvironmentDev,
			expected: "Dev",
		},
		{
			name:     "stg",
			in:       enum.EnvironmentStg,
			expected: "Stg",
		},
		{
			name:     "prod",
			in:       enum.EnvironmentProd,
			expected: "Prod",
		},
		{
			name:     "undefined",
			in:       enum.EnvironmentUndefined,
			expected: "Undefined",
		},
		{
			name:     "random",
			in:       enum.Environment(16),
			expected: "Environment(16)",
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

func FuzzEnvironmentFrom(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		t.Parallel()
		out := enum.EnvironmentFrom(s).String()
		require.New(t).Equal(enum.EnvironmentUndefined.String(), out)
	})
}
