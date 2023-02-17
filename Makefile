#DSN="host=localhost port=5432 user=postgres password=password dbname=vueapi sslmode=disable timezone=UTC connect_timeout=5"
DSN="host=dpg-cenp3mta4991ihnsqrf0-a.oregon-postgres.render.com port=5432 user=goapp password=rSG4t2vOYkVhvEZsaEGcYzir5TYWj8MQ dbname=vueapi timezone=UTC connect_timeout=10"
BINARY_NAME=vueapi
ENV=development
$PORT=8080

## build: Build binary
build:
	@echo "Building back end..."
	go build -o ${BINARY_NAME} ./cmd/api/
	@echo "Binary built!"

## run: builds and runs the application
run: build
	@echo "Starting back end..."
	@env DSN=${DSN} ENV=${ENV} ./${BINARY_NAME} &
	@echo "Back end started!"

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

## start: an alias to run
start: run

## stop: stops the running application
stop:
	@echo "Stopping back end..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped back end!"

## restart: stops and starts the running application
restart: stop start