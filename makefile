APP_BINARY=app.exe

##build: build binary file for service
build:
	@echo "Building front end binary..."
	cd ./cmd/app && go build -o ${APP_BINARY} ./
	@echo "Done!"

##run: build and run service
run: build
	cd ./cmd/app && start /B ${APP_BINARY} &

##stop: stop service
stop:
	@echo "Stopping service..."
	@taskkill /IM "${APP_BINARY}" /F
	@echo "Stopped!"

##test: run tests in service
test:
	go test -v -race -timeout 30s ./...

test_cover:
	go test  -coverprofile=coverage.out -v -race -timeout 30s ./...
	go tool cover -html=coverage.out


## up starts app in docker compose
up:
	@echo Starting Docker images....
	docker-compose up -d 
	@echo Docker images started!

## down: stop docker compose
down:
	@echo Stopping docker compose...
	docker-compose down
	@echo Done!


.DEFAULT_GOAL := up
