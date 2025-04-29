APP_NAME=todoApp
CMD_DIR=./cmd
BIN_DIR=./bin

build:
	go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_DIR)/main.go

run: build
	$(BIN_DIR)/$(APP_NAME)

clean:
	rm -rf $(BIN_DIR)

docker-build:
	docker build -t $(APP_NAME)-image .

docker-run:
	docker-compose up

docker-stop:
	docker-compose down
