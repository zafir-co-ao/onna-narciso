APP_NAME = narciso
GO = go
GO_TEST = $(GO) test -cover
TEST_DIR = ./...

.PHONY: tests

tests:
	$(GO_TEST) $(TEST_DIR)


clean:
	rm -fr bin
	find . -name '*_templ.go' -type f -exec rm {} \;

build: clean
	mkdir bin
	cp -r web/static bin/static
	templ generate
	bunx tailwindcss -i web/assets/css/main.css -o web/static/css/styles.css
	go build -o bin/$(APP_NAME) cmd/web/main.go

build-dev: clean
	mkdir -p bin/static

run-dev: build-dev
	air
