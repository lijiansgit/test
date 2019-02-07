CURRENT_DIR=$(shell pwd)

common:

	go get -v github.com/valyala/fasthttp
	go build -o test

prd: common
	echo "release"

pre: common
	echo "pre"

qa: common
	echo "qa"
	./test
