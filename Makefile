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

check:
	@go run $(licenser) verify
	@go run $(linter) run
	@go run $(releaser) check
	@go run $(wire) check ./...

fmt:
	@gofmt -w -s .
	@goimports -w .
	@go vet ./...
	@go mod tidy
	@go run $(licenser) apply -r "ClavinJune/bjora" 2> /dev/null

test:
	@go test -coverprofile=coverage -covermode=count ./...
	@go tool cover -html=coverage

wire:
	@go run $(wire) ./...
