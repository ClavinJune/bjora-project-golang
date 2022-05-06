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

package user_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	log.Println("setup test environment variable")
	err := godotenv.Overload("../../.env.test")
	if err != nil {
		panic(err)
	}

	log.Println("migrate up test database")
	migrator, err := migrate.New(
		"file://../../blueprint/db-migration",
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("POSTGRES_USERNAME"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DATABASE"),
		),
	)
	if err != nil {
		panic(err)
	}

	_ = migrator.Down()
	if err := migrator.Up(); err != nil {
		panic(err)
	}
	defer func() {
		log.Println("migrate down test database")
		_ = migrator.Down()
	}()
	m.Run()
}
