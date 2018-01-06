.PHONY: run test

run:
	@docker-compose build
	@docker-compose run --service-ports --name tweetatlas --rm service

test:
	@docker-compose up --exit-code-from test --force-recreate --build test
	@docker-compose rm -f test
