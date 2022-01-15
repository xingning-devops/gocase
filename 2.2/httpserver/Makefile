export tag=v1.0

build:
	echo "build binary"
	mkdir bin/
	GOOS=linux GOARCH=amd64  CGO_ENABLED=0 go build -o bin/$(name) -v

release: build
	echo "build images"
	docker build -t xingning03/httpserver:${tag} .
	

push: release
	echo "push"
	docker push xingning03/httpserver:${tag}

clean:
	echo "rm build"
	rm -rf bin/
	docker rmi xingning03/httpserver:${tag}
