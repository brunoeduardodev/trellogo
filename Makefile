dev:
	go run ./cmd/rest-api

docker-compose-up:
	docker-compose --env-file .env up -d

docker-compose-down:
	docker-compose down