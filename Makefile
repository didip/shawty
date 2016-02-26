BASEDIR  = $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

NAME = go-shorten/webapp

build:
	docker build -t registry.dev.databricks.com/$(NAME) $(BASEDIR)

push:
	docker push registry.dev.databricks.com/$(NAME)

restart_pods:
	kubectl rolling-update go-shorten --image registry.dev.databricks.com/$(DOCKER_NAME):latest --update-period=5s

.PHONY: build push restart_pods
