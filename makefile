VERSION = $(shell git branch --show-current)
DATABASE_CONNECT="postgres://postgres:postgres@localhost:5432/stone?sslmode=disable"
MIGRATION_SOURCE="file://db/migrations"

NAME = initial_schemas


mkfile_path:=$(word $(words $(MAKEFILE_LIST)),$(MAKEFILE_LIST))
mkfile_dir:=$(shell cd $(shell dirname $(mkfile_path)); pwd)

# commands to start project

.PHONY: run-docker
run-docker:
	rm -f .ssh
	ln -s "${HOME}/.ssh" .ssh
	docker-compose build
	docker-compose up

.PHONY: dev
dev:
	VERSION=$(VERSION) go run main.go

# commands to test project
.PHONY: test
test:
	chmod 777 $(mkfile_dir)/.gocache || true
	docker run --rm --name go-test \
		-v $(mkfile_dir):/opt/app \
		-v $(mkfile_dir)/.gocache:/go \
		-w /opt/app \
		golang:latest \
		go test -count=1 -cover -failfast -coverprofile=coverage.out ./...
	docker image rm golang

.PHONY: test-cover
test-cover: test
	go tool cover -html=coverage.out

# comando to generante docs
.PHONY: docs
docs:
	go install github.com/swaggo/swag/cmd/swag@v1.6.7
	go get github.com/swaggo/echo-swagger@v1.3.0
	swag init --parseDependency --parseInternal

# command to generate mocks
.PHONY: mocks
mocks: 
	rm -rf ./mocks
	mkdir mocks

	# application
	mockgen -source=./app/account/account.go -destination=./mocks/account_app_mock.go -package=mocks -mock_names=App=MockAccountApp
	mockgen -source=./app/login/login.go -destination=./mocks/login_app_mock.go -package=mocks -mock_names=App=MockLoginApp
	mockgen -source=./app/transfer/transfer.go -destination=./mocks/transfer_app_mock.go -package=mocks -mock_names=App=MockTransferApp

	# stores
	mockgen -source=./store/account/account.go -destination=./mocks/account_mock.go -package=mocks -mock_names=Store=MockAccountStore
	mockgen -source=./store/transfer/transfer.go -destination=./mocks/transfer_mock.go -package=mocks -mock_names=Store=MockTransferStore

	# utils
	mockgen -source=./utils/session/session.go -destination=./mocks/session_mock.go -package=mocks -mock_names=Store=MockSessionStore
	mockgen -source=./utils/password/password.go -destination=./mocks/password_mock.go -package=mocks -mock_names=Store=MockPasswordStore

# command to generate migration
.PHONY: migration-create
migration-create:
	migrate create -ext sql -dir db/migrations -seq $(NAME)

.PHONY: migration-down
migration-down:
	migrate -source $(MIGRATION_SOURCE) -database $(DATABASE_CONNECT) --verbose down 1
