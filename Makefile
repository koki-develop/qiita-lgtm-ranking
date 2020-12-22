.PHONY: build clean deploy

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o build/update src/handlers/update/main.go

clean:
	rm -rf ./build ./vendor go.sum

deploy: clean build
	sls deploy --verbose
