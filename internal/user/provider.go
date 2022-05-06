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
	"os"
	"sync"

	"github.com/jmoiron/sqlx"

	"github.com/bwmarrin/snowflake"

	"github.com/clavinjune/bjora-project-golang/pkg/enum"

	"github.com/clavinjune/bjora-project-golang/pkg"
	"github.com/google/wire"
)

var (
	repoInst *repository
	repoOnce sync.Once

	svcInst *service
	svcOnce sync.Once

	// ProviderSet provides user's providers
	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideRepository, ProvideService,
		wire.Bind(new(pkg.UserRepository), new(*repository)),
		wire.Bind(new(pkg.UserService), new(*service)),
	)
)

func newRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

// ProvideRepository produces *user.repository
func ProvideRepository(db *sqlx.DB) *repository {
	repoOnce.Do(func() {
		repoInst = newRepository(db)
	})

	return repoInst
}

func newService(repo pkg.UserRepository, snowNode *snowflake.Node) *service {
	return &service{
		repo:     repo,
		snowNode: snowNode,
	}
}

// ProvideService produces *user.service
func ProvideService(repo pkg.UserRepository, snowNode *snowflake.Node) *service {
	svcOnce.Do(func() {
		svcInst = newService(repo, snowNode)
	})

	// always create new service on test
	if enum.EnvironmentFrom(os.Getenv("APP_ENV")) == enum.EnvironmentTest {
		return newService(repo, snowNode)
	}
	return svcInst
}
