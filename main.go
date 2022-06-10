package main

import (
	"douyin-file-service/config"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"strconv"
)

func newApp() *iris.Application {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, notFound)
	mvc.Configure(app.Party("/file"), func(app *mvc.Application) {
		app.Handle(new(FileController))
	})
	return app
}

func main() {
	addr := strconv.Itoa(config.AppConfig.GetInt("server.port"))
	app := newApp()
	app.UseGlobal(before)
	app.Run(iris.Addr(":"+addr), iris.WithCharset("UTF-8"), iris.WithoutPathCorrectionRedirection)
}

func before(ctx iris.Context) {
	iris.New().Logger().Info(ctx.Path())
	ctx.Next()
}

func notFound(ctx iris.Context) {
	code := ctx.GetStatusCode()
	msg := "404 Not Found"
	ctx.JSON(iris.Map{
		"Message": msg,
		"Code":    code,
	})
}
