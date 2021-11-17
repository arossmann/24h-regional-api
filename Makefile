DOCKER_IMAGE_VERSION=0.1
DOCKER_IMAGE_NAME=24h-regional-api
DOCKER_IMAGE_TAGNAME=$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION)
PORT_IN=8888
PORT_OUT=8888

default: build ## default = build

build: ## build the image
	docker build -t $(DOCKER_IMAGE_TAGNAME) .
	docker tag $(DOCKER_IMAGE_TAGNAME) $(DOCKER_IMAGE_NAME):latest

test: ## test the image
	docker run --rm $(DOCKER_IMAGE_TAGNAME) /bin/echo "Success."

rmi: ## remove the image
	docker rmi -f $(DOCKER_IMAGE_TAGNAME)

rebuild: rmi build ## rebuild it

run: ## run the image
	docker run --env-file=.env -d -p  $(PORT_OUT):$(PORT_IN) $(DOCKER_IMAGE_NAME):latest

help: ## This help dialog
	@IFS=$$'\n' ; \
		help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//'`); \
		for help_line in $${help_lines[@]}; do \
			IFS=$$'#' ; \
			help_split=($$help_line) ; \
			help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
			help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
			printf "%-10s %s\n" $$help_command $$help_info ; \
		done

.PHONY: help run list build test rmi rebuild