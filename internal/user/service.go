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
	"context"

	"github.com/bwmarrin/snowflake"

	"github.com/clavinjune/bjora-project-golang/internal/util"

	"github.com/clavinjune/bjora-project-golang/pkg"
)

type service struct {
	repo     pkg.UserRepository
	snowNode *snowflake.Node
}

func (s service) Store(ctx context.Context, u *pkg.UserSpec) (*pkg.UserSpec, error) {
	u.ID = s.snowNode.Generate()
	e := u.ToEntity()
	e.Entity = pkg.NewEntity()

	stored, err := s.repo.Store(ctx, e)

	if err != nil {
		return nil, util.WrapError(err)
	}

	result, err := pkg.UserSpecFromEntity(stored)
	if err != nil {
		return nil, util.WrapError(err)
	}

	return result, nil
}
