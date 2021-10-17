export tag=v1.0

main:
	mkdir -p bin/client/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/client/amd64

release:
    docker build -t hub.docker.com/httpclient:${tag}
	docker push hub.docker.com/httpclient:${tag}