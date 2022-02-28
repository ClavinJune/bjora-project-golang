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

package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/clavinjune/bjora-project-golang/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

const (
	timeout = time.Minute
	dbConn  = 5
)

func createApp() *fiber.App {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
		CaseSensitive: true,
		Immutable:     true,
		ReadTimeout:   timeout,
		WriteTimeout:  timeout,
		IdleTimeout:   timeout,
		AppName:       "bjora-api",
	})

	app.Use(requestid.New(), logger.New())

	return app
}

func createChannel() (chan os.Signal, func()) {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	return stopCh, func() {
		close(stopCh)
	}
}

func startApp(app *fiber.App) {
	if err := app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Println("app forced to shutdown")
	} else {
		log.Println("app shutdowned gracefully")
	}
}

func shutdownApp(app *fiber.App) {
	if err := app.Shutdown(); err != nil {
		log.Println("app.Shutdown() error", err.Error())
	} else {
		log.Println("app.Shutdown()")
	}
}

func connectDb() (*sql.DB, func()) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
		"disable",
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	db.SetConnMaxIdleTime(timeout)
	db.SetConnMaxLifetime(timeout)
	db.SetMaxIdleConns(dbConn)
	db.SetMaxOpenConns(dbConn)

	return db, func() {
		_ = db.Close()
	}
}

func main() {
	app := createApp()
	ch, closeCh := createChannel()
	defer closeCh()

	db, closeDb := connectDb()
	defer closeDb()

	user.Wire(db).ApplyRoute(app.Group("/user"))

	go startApp(app)
	<-ch
	shutdownApp(app)
}
