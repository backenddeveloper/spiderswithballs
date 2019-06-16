default:

build:
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o spiderswithballs.wasm main.go
	docker build -t backenddeveloper/spiderswithballs:$$(git rev-parse --verify HEAD)
	docker tag backenddeveloper/spiderswithballs:$$(git rev-parse --verify HEAD) backenddeveloper/spiderswithballs:latest

localhost:
	docker run -v `pwd`:/usr/share/nginx/html -p 8080:80 nginx:1.17.0
