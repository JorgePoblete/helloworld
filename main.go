package main

import (
	"log"

	"github.com/kataras/iris/v12"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	app := iris.New()

	app.Get("/healthcheck", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
	})

	app.Get("/metrics", func(ctx iris.Context) {
		promhttp.Handler().ServeHTTP(ctx.ResponseWriter(), ctx.Request())
	})

	app.Get("/hello", func(ctx iris.Context) {
		log.Printf("%s", ctx.Request().URL)
		count := 1
		for i := 0; i <= 1000000; i++ {
			count = i
		}
		ctx.Text("Hello, world %d times\n", count)
		ctx.StatusCode(iris.StatusOK)
	})

	app.Listen(":8080")
}
