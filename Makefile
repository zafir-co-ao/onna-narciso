APP_NAME = narciso
GO = go
GO_TEST = $(GO) test -cover
TEST_DIR = ./...

.PHONY: tests

tests:
	$(GO_TEST) $(TEST_DIR)

build:
	rm -fr bin
	mkdir bin
	cp -r web/assets bin/assets
	cp -r web/static bin/static
	templ generate
	bunx tailwindcss -i web/assets/css/main.css -o web/assets/css/styles.css
	go build -o bin/$(APP_NAME) cmd/web/main.go

build-dev:
	rm -fr bin
	mkdir bin

run-dev: build-dev
	air
