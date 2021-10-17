export tag=v1.0

main:
	mkdir -p ../bin/server/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/server/amd64/main

release:
    docker build -t hub.docker.com/httpserver:${tag}
    docker push hub.docker.com/httpserver:${tag}