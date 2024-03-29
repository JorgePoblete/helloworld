package main

import (
	"log"
	"math"

	"github.com/kataras/iris/v12"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.Application().Logger().Info(ctx.Request().URL)
	})

	app.Get("/healthcheck", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
	})

	app.Get("/", func(ctx iris.Context) {
		ctx.Application().Logger().Info(ctx.Request().URL)
		ctx.StatusCode(iris.StatusOK)
	})

	app.Get("/metrics", func(ctx iris.Context) {
		promhttp.Handler().ServeHTTP(ctx.ResponseWriter(), ctx.Request())
	})

	app.Get("/hello", func(ctx iris.Context) {
		log.Printf("%s", ctx.Request().URL)
		count := 1.0
		for i := 0; i < 100000; i++ {
			count += math.Sqrt(count)
		}
		ctx.Text("Hello, world %.10f times\n", count)
		ctx.StatusCode(iris.StatusOK)
	})

	app.Listen("0.0.0.0:8080")
}
