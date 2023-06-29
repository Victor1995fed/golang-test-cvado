include .env
export

up:
	docker-compose up -d --build
	sleep 10
	docker-compose exec -T db mysql -u ${DB_USERNAME} -p${MYSQL_ROOT_PASSWORD} -D ${MYSQL_DATABASE} < ./data.sql
	$(MAKE) up_server

up_server:
	@echo "Starting build application..."
	go build  -o test_server cmd/server/main.go
	./test_server
	@echo "Application started!"

up_client:
	@echo "Starting build application..."
	@echo "Build client..."
	go build  -o test_client cmd/client/main.go
	./test_client
	@echo "Done!"
stop:
	docker-compose stop

tests:
	go test -v ./...
gen-protoc:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  proto/server.proto