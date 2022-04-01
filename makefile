VERSION = $(shell git branch --show-current)
DATABASE_CONNECT="postgres://postgres:postgres@localhost:5432/stone?sslmode=disable"
MIGRATION_SOURCE="file://db/migrations"

NAME = initial_schemas

# commands to start project

.PHONY: dev
dev:
	rm -f .ssh
	ln -s "${HOME}/.ssh" .ssh
	docker-compose build
	docker-compose up

.PHONY: watch
watch:
	nodemon --watch pkg --ext ".go" --exec docker-compose restart stone-api

.PHONY: run
run:
	VERSION=$(VERSION) go run main.go

# commands to test project
.PHONY: test
test:
	go test -count=1 -cover -failfast ./... -coverprofile=coverage.out 

.PHONY: test-cover
test-cover: test
	go tool cover -html=coverage.out

# command to generate mocks
.PHONY: mocks
mocks: 
	rm -rf ./mocks
	mkdir mocks

	# application
	mockgen -source=./app/health/health.go -destination=./mocks/health_app_mock.go -package=mocks -mock_names=App=MockHealthApp

	# stores
	mockgen -source=./store/health/health.go -destination=./mocks/health_mock.go -package=mocks -mock_names=Store=MockHealthStore

# command to generate migration
.PHONY: migration-create
migration-create:
	migrate create -ext sql -dir db/migrations -seq $(NAME)

.PHONY: migration-down
migration-down:
	migrate -source $(MIGRATION_SOURCE) -database $(DATABASE_CONNECT) --verbose down 1
