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
	"database/sql"
	"sync"

	"github.com/clavinjune/bjora-project-golang/pkg"

	"github.com/google/wire"
)

var (
	repoInst *postgres
	repoOnce sync.Once

	svcInst *service
	svcOnce sync.Once

	hdlInst *handler
	hdlOnce sync.Once

	// ProviderSet provides user's providers
	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideRepository,
		ProvideService,
		ProvideHandler,

		wire.Bind(new(pkg.UserRepository), new(*postgres)),
		wire.Bind(new(pkg.UserService), new(*service)),
		wire.Bind(new(pkg.UserHandler), new(*handler)),
	)
)

// ProvideRepository provides user postgres
func ProvideRepository(db *sql.DB) *postgres {
	repoOnce.Do(func() {
		repoInst = &postgres{
			db: db,
		}
	})

	return repoInst
}

// ProvideService provides user service
func ProvideService(repo pkg.UserRepository) *service {
	svcOnce.Do(func() {
		svcInst = &service{
			repo: repo,
		}
	})

	return svcInst
}

// ProvideHandler provides user HTTP handlerutil
func ProvideHandler(svc pkg.UserService) *handler {
	hdlOnce.Do(func() {
		hdlInst = &handler{
			svc: svc,
		}
	})

	return hdlInst
}
