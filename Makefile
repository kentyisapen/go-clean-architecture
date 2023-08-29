build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/api/main.go

run: 
	docker-compose up --build server

db.init:
	docker compose down && \
	rm -rf storage/* && \
	rm -rf .data/db && \
	docker compose up -d
