.PHONY: build
build: 
	@go mod tidy && \
		go generate ./src/cmd && \
		go build -o ./build/app ./src/cmd

.PHONY: run
run: 
	@air --build.cmd "go build -o ./build/app ./src/cmd" --build.bin "./build/app"
	
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