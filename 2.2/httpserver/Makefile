export tag=v1.0

release:
	echo "build images"
	docker build -t xingning03/httpserver:${tag} .
	

push: release
	echo "push"
	docker push xingning03/httpserver:${tag}
