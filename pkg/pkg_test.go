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
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	log.Println("setup test environment variable")
	err := godotenv.Overload("../.env.test")
	if err != nil {
		panic(err)
	}
	m.Run()
}
