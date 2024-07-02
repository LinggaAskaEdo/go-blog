SQL_PATH = './etc/sql/'

.PHONY: build
build: 
	@go mod tidy && \
		go generate ./src/cmd && \
		go build -o ./build/app ./src/cmd

.PHONY: live-run
live-run: build
	@air --build.cmd "go build -o ./build/app ./src/cmd" --build.bin "./build/app"

.PHONY: native-run
native-run: build
	@./build/app

.PHONY: db-start
db-start: 
	@docker start mysql-docker postgres-docker redis-docker

.PHONY: db-stop
db-stop: 
	@docker stop mysql-docker postgres-docker redis-docker

.PHONY: rabbit-start
rabbit-start: 
	@docker start rabbitmq-docker

.PHONY: rabbit-stop
rabbit-stop: 
	@docker stop rabbitmq-docker

# .PHONY: generate-migrate-file
# generate-migrate-file: 
# 	@if [ -z "$(name)" ]; then \
# 		echo "Param name is missing !!!\nex: make generate-migrate-file name=create_user_table"; \
# 	else \
# 		migrate create -ext sql -dir ${SQL_PATH} -seq $(name); \
# 	fi

PHONY: create-migration-file
create-migration-file: 
	@if [ -z "$(name)" ]; then \
		echo "Param name is missing !!!\nex: make generate-migrate-file name=create_user_table"; \
	else \
		goose -dir ${SQL_PATH} create $(name) sql; \
	fi	