export tag=v1.0-metrics

build:
	echo "build images"
	docker build -t xingning03/httpserver:${tag} .

push: build
	echo "push"
	docker push xingning03/httpserver:${tag}
