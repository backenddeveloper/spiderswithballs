default:

clean:
	find . -name *.wasm -delete

compile:
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o html/spiderswithballs.wasm

build: compile
	docker build -t backenddeveloper/spiderswithballs:$$(git rev-parse --verify HEAD) .
	docker tag backenddeveloper/spiderswithballs:$$(git rev-parse --verify HEAD) backenddeveloper/spiderswithballs:latest

localhost:
	docker run -v `pwd`/html:/usr/share/nginx/html -p 8080:80 backenddeveloper/spiderswithballs:latest
