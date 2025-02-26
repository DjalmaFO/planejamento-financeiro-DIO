build-image:
	docker build -t djalmafo/finance .

run-app:
	docker-compose -f ./app.yml up -d

test-unit:
	go test ./...