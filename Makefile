
export COMPOSE_PROJECT_NAME=urgence-tree
export UID=$(shell id -u)
export NAME_UID=$(shell id -u -n)
export GUID=$(shell id -g)
export NAME_GUID=$(shell id -g -n)

.PHONY: init
init:
	git config --global diff.submodule log
	git submodule init
	git submodule update

.PHONY: pull-submodule
pull-submodule:
	@cd Common; git pull; cd ..

.PHONY: build
build:
	docker-compose -f docker-compose.dev.yaml build --force-rm --parallel

############################################ DEV ############################################

############## SERV ##############

.PHONY: dev-db
dev-db:
	docker-compose -f docker-compose.dev.yaml up -d --remove-orphans --build dev-db

.PHONY: dev-decision-tree
dev-decision-tree:
	docker-compose -f docker-compose.dev.yaml up --remove-orphans --build decision-tree

.PHONY: test-image
test-image: dev-db
	docker-compose -f docker-compose.dev.yaml up -d --remove-orphans --build image-decision-tree

############################################ CLEAR ############################################

.PHONY: clear
clear:
	docker-compose -f docker-compose.dev.yaml down --remove-orphans


.PHONY: clear-volumes
clear-volumes:
	docker-compose -f docker-compose.dev.yaml down --remove-orphans --volumes

############################################ DOWN ############################################

.PHONY: stop-db
stop-db:
	docker-compose -f docker-compose.dev.yaml stop
