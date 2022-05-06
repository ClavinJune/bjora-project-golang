# Copyright 2022 ClavinJune/bjora
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

include tools.mk
include .env
export

check:
	@go run $(licenser) verify
	@go run $(linter) run
	@go run $(wire) check ./...

clean:
	@rm -rf result.json coverage.out
	@find . -type f -name "*_gen.go" -delete
	@find . -type d -name "mocks" -exec rm -rf {} +

db/connect:
	@PGPASSWORD="${POSTGRES_PASSWORD}" psql -h "${POSTGRES_HOST}" -U "${POSTGRES_USERNAME}" "${POSTGRES_DATABASE}"

db/connect/test:
	@PGPASSWORD="${POSTGRES_PASSWORD}_test" psql -p 5433 -h "${POSTGRES_HOST}" -U "${POSTGRES_USERNAME}_test" "${POSTGRES_DATABASE}_test"

db/migrate/create:
	@go run -tags "postgres" $(migrator) \
 	create -ext sql -dir blueprint/db-migration \
 	-format "200601021504" \
 	create_[change_this]_table

db/migrate/up:
	@go run -tags "postgres" $(migrator) \
	-source file://blueprint/db-migration \
	-database "postgres://${POSTGRES_USERNAME}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable" \
	-verbose up

db/migrate/down:
	@go run -tags "postgres" $(migrator) \
	-source file://blueprint/db-migration \
	-database "postgres://${POSTGRES_USERNAME}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable" \
	-verbose down

docker/compose/up:
	@docker compose up -d

docker/compose/up/test:
	@docker compose -f docker-compose-test.yml up -d

docker/compose/down:
	@docker compose stop && docker compose down

docker/compose/down/test:
	@docker compose -f docker-compose-test.yml stop && docker compose -f docker-compose-test.yml down

docker/volume/create:
	@docker volume create minio-data
	@docker volume create postgres-data

docker/volume/clean:
	@docker volume rm minio-data || true
	@docker volume rm postgres-data || true

fmt:
	@gofmt -w -s .
	@go run $(importer) -w .
	@go vet ./...
	@go mod tidy
	@go run $(licenser) apply -r "ClavinJune/bjora" 2> /dev/null

gen: clean tools/install/stringer
	@go generate ./...
	@$(MAKE) wire mock fmt

mock:
	@go run $(mocker) --all --with-expecter --output "./pkg/mocks"

test: docker/compose/up/test
	@sleep 1
	@go test -count=1 -v ./...
	@$(MAKE) docker/compose/down/test

test/fuzz:
	@go test -v -fuzz=FuzzGenderFrom -fuzztime=1s ./pkg/enum/
	@go test -v -fuzz=FuzzEnvironmentFrom -fuzztime=1s ./pkg/enum/

test/coverage:
	@go test -count=1 -v -json -coverprofile=coverage.out -covermode=count `go list ./... | grep -v mocks` > result.json
	@go tool cover -html=coverage.out

tools/install/stringer:
	@go install $(stringer)

wire:
	@go run $(wire) ./...
