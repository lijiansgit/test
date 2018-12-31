package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/valyala/fasthttp"
)

func main() {
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

	log.Println("fasthttp Listen: 80")
	err := fasthttp.ListenAndServe(":80", m)
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
