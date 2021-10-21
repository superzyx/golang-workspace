export tag=v1.0

main:
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/main

release:
	docker build -t superzyx12138/httpserver:${tag} .
	docker push superzyx12138/httpserver:${tag}
