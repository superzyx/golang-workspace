export tag=v2.0

main:
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/main

release:
	docker build -t superzyx12138/product-server:${tag} .
	docker push superzyx12138/product-server:${tag}
