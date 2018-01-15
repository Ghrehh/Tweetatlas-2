.PHONY: run test

run:
	@docker-compose build service
	@docker-compose up service

test:
	@docker-compose build test
	@docker-compose run --rm test
