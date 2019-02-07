package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/valyala/fasthttp"
)

func main() {
	log.Println("ENV TEST:", os.Getenv("TEST"))
	log.Println("APP_ENV:", os.Getenv("APP_ENV"))
	log.Println("runtime.NumCPU():", runtime.NumCPU())
	log.Println("PPID:", os.Getppid())
	log.Println("PID:", os.Getpid())

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go receiveSignal(c)

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

// todo
func receiveSignal(c chan os.Signal) {
	for s := range c {
		switch s {
		case syscall.SIGTERM:
			log.Println("test sigterm SIGTERM", s)
			log.Println("sleep 5")
			time.Sleep(5 * 1e9)
			log.Println("exit ok")
			os.Exit(1)
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT:
			log.Println("exit SIGHUP SIGINT SIGQUIT", s)
			os.Exit(2)
		case syscall.SIGUSR1:
			log.Println("usr1", s)
			os.Exit(3)
		case syscall.SIGUSR2:
			log.Println("usr2", s)
			os.Exit(4)
		default:
			log.Println("other", s)
			os.Exit(5)
		}
	}
}
