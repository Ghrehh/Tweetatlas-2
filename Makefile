.PHONY: run test

run:
	@docker-compose build
	@docker-compose run --service-ports --name tweetatlas --rm service

test:
	@docker-compose build test
	@docker-compose run --rm test
	@docker-compose down
