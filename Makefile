######################################  org.inspiry.iss ###################################
######################################  ---------    	###################################
######################################  ISS快捷入口  	    ###################################
######################################  启动项目注意事项	###################################
######### 注意事项：
# 启动项目:
# 					s1. 先启动: make docker-env-start
#					s2. 后启动: make docker-start
# 停止项目:
# 					s1. 先停止: make docker-stop
#					s2. 后停止: make docker-env-stop
# 重启项目:
# 					重启项目: make docker-restart
#					重启环境: make docker-env-restart

RELEASE_VERSION  = v1.1.1.0

.PHONY: all
all: help


.PHONY: go
tidy: ## Updates the go modules
	go mod tidy

.PHONY: go
cloc: ## Show project type
	go install github.com/hhatto/gocloc/cmd/gocloc@latest && gocloc --include-lang=go  --match-d= ./apps/ .


.PHONY: go
cover:  ## Displays test coverage in the iss  packages
	go test -coverprofile=cover-iss.out  && go tool cover -html=cover-iss.out

.PHONY: go
test: ## Tests the entire project
	go test -count=1 \
			-race \
			-coverprofile=coverage.txt \
			-covermode=atomic \
			./...


.PHONY: go
lint: ## Lints the entire project
	@golangci-lint run --timeout=3m

.PHONY: go
build: ## Compilation main.go to iss file
	#@go build -o app  main.go
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app1

.PHONY: go
run:build ## Compilation run iss
	@./app

.PHONY: go
back-run: ## Compilation back run iss
	sudo nohup ./app >> logs/app.log &




.PHONY: git-tag
git-tag: ## Creates release tag
	git tag $(RELEASE_VERSION)
	git push origin $(RELEASE_VERSION)


.PHONY: env
env-start: ## Start Iss Env file {docker-compose-env.yml}
	printf 	 `make docker-env-start`
	docker-compose -p iss_env -f docker-compose-env.yml up -d

.PHONY: env
env-stop: ## Stop Iss Env file {docker-compose-env.yml}
				 ## make docker-env-stop
	docker-compose -p iss_env -f docker-compose-env.yml down

.PHONY: env
env-restart:  ## Restart Iss Env file {docker-compose-env.yml}
			   ## make docker-env-restart
	docker-compose  -p iss_env restart

.PHONY: docker
docker-start:  ## Start Iss file {docker-compose.yml}
			   ## make docker-start
	docker-compose -p iss_app  up -d

.PHONY: docker

docker-stop:  ## Stop Iss file {docker-compose.yml}
		## make docker-stop
	docker-compose -p iss_app down

	#docker-compose docker-compose.yml down

.PHONY: docker
docker-restart:  ## Restart Iss file {docker-compose.yml}
	docker-compose  -p iss_app restart

.PHONY: docker
docker-exec-mysql:  ## Exec mysql, {use iss;} export iss.sql{source /etc/sql/iss.sql}
	docker exec -it mysql bash

.PHONY: help
help: ## Display available commands

	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk \
		'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'



.PHONY: atlas
atlas: #:
	atlas schema inspect -u "mysql://root:root@localhost:3306/unity" > schema.hcl



###
# - crawLab - craw
# - zinc	- nosql
# - gotify  - websocket
###
