package router

import "github.com/kataras/iris/v12"

func InitRoute(app *iris.Application) {
	initBookRoute(app)
	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("helldd")
	})
}
