.PHONY: env
dev:
	ssh -i jcb-dev.pem centos@35.77.156.137

.PHONY: env
uat:
	ssh -i jcb-uat.pem centos@52.193.130.242

.PHONY: env
pro:
	ssh -i jcb-prod.pem centos@54.238.37.123

#secret:
.PHONY: sql
sql: ## mysql : root : PXDN93VRKUm8TeE7
	scp -i jcb-dev.pem  sql/iss.sql  centos@ec2-3-115-4-116.ap-northeast-1.compute.amazonaws.com:/home/centos/sql


#secret:
.PHONY: iss
iss: ##
	scp -i jcb-dev.pem ../app  centos@ec2-3-115-4-116.ap-northeast-1.compute.amazonaws.com:/home/centos/www



#secret:
.PHONY: build
build: ##
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

