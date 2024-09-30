# Makefile

include .env

postgres_init:
	./scripts/script.sh

api_run:
	@docker build -t $(IMAGE_NAME) .
	@docker run --name ${CONTAINER_NAME} -p 3000:3000 $(IMAGE_NAME)

.PHONY: postgres_init api_run