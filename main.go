package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/valyala/fasthttp"
)

func main() {
	log.Println("ENV TEST:", os.Getenv("TEST"))
	log.Println("runtime.NumCPU():", runtime.NumCPU())

	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			rootFunc(ctx)
		case "/ping":
			pingFunc(ctx)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}

	log.Println("fasthttp Listen: 8080")
	err := fasthttp.ListenAndServe(":8080", m)
	if err != nil {
		panic(err)
	}
}

func rootFunc(ctx *fasthttp.RequestCtx) {
	log.Println(string(ctx.RequestURI()))
	fmt.Fprintf(ctx, "hello world!")
}

func pingFunc(ctx *fasthttp.RequestCtx) {
	log.Println(string(ctx.RequestURI()))
	fmt.Fprintf(ctx, "ok")
}
