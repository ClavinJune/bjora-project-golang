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
	"os"
	"testing"
	"time"

	"github.com/clavinjune/bjora-project-golang/pkg"
	"github.com/stretchr/testify/require"
)

func TestNewEntity(t *testing.T) {
	t.Parallel()
	appName := os.Getenv("APP_NAME")
	r := require.New(t)
	e := pkg.NewEntity()

	r.True(e.CreatedAt.Valid)
	r.False(e.CreatedAt.Time.IsZero())
	r.True(e.CreatedBy.Valid)
	r.Equal(appName, e.CreatedBy.String)

	r.True(e.LastModifiedAt.Valid)
	r.False(e.LastModifiedAt.Time.IsZero())
	r.True(e.LastModifiedBy.Valid)
	r.Equal(appName, e.LastModifiedBy.String)

	r.True(e.IsActive.Valid)
	r.True(e.IsActive.Bool)
}

func TestEntity_Update(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	e := pkg.NewEntity()
	time.Sleep(time.Millisecond)

	prevTime := e.LastModifiedAt.Time
	by := "TestEntity_Update"

	e.Update(by)

	r.True(e.LastModifiedAt.Valid)
	r.NotEqual(prevTime, e.LastModifiedAt.Time)
	r.Less(prevTime, e.LastModifiedAt.Time)

	r.True(e.LastModifiedBy.Valid)
	r.Equal(by, e.LastModifiedBy.String)
}

func TestEntity_Delete(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	e := pkg.NewEntity()
	time.Sleep(time.Millisecond)

	prevTime := e.LastModifiedAt.Time
	by := "TestEntity_Delete"

	e.Delete(by)

	r.True(e.LastModifiedBy.Valid)
	r.NotEqual(prevTime, e.LastModifiedAt.Time)
	r.Less(prevTime, e.LastModifiedAt.Time)

	r.True(e.LastModifiedBy.Valid)
	r.Equal(by, e.LastModifiedBy.String)

	r.True(e.IsActive.Valid)
	r.False(e.IsActive.Bool)
}
