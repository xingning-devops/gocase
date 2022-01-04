export: name=httpserver

build:
	mkdir bin/
	GOOS=linux GOARCH=amd64 go build -o bin/$(name) -v


clean:
	rm -rf bin/