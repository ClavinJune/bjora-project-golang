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

	"github.com/bwmarrin/snowflake"
	"github.com/clavinjune/bjora-project-golang/internal/util"
	"github.com/stretchr/testify/require"
)

func TestProvideSnowflake(t *testing.T) {
	t.Parallel()

	gen1 := util.ProvideSnowflake()
	gen2 := util.ProvideSnowflake()

	ch := make(chan snowflake.ID, 2)

	goFunc := func(g *snowflake.Node, c chan<- snowflake.ID) {
		c <- g.Generate()
	}

	go goFunc(gen1, ch)
	go goFunc(gen2, ch)

	ids := [2]snowflake.ID{
		<-ch,
		<-ch,
	}

	r := require.New(t)
	r.NotEqual(ids[0], ids[1])
}
