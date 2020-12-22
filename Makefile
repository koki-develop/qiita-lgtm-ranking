.PHONY: test build clean deploy

init:
	docker-compose build
	docker-compose run --rm app yarn install --check-files

test:
	docker-compose run --rm app ./bin/test

build: test
	docker-compose run --rm app ./bin/build

clean:
	rm -rf ./build ./vendor

deploy: clean build
	docker-compose run --rm app yarn run sls deploy --verbose
