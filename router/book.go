package router

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func initBookRoute(app *iris.Application) {
	booksAPI := app.Party("/books")
	{
		booksAPI.Use(iris.Compression)
		// GET: http://localhost:8080/books
		booksAPI.Get("/", list)
		// POST: http://localhost:8080/books
		booksAPI.Post("/", create)
	}
}

// Book example.
type Book struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desciption"`
	Content string `json:"content"`
}

var Books []Book

func list(ctx iris.Context) {

	ctx.JSON(Books)
	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	// ctx.Negotiate(books)
}

func create(ctx iris.Context) {
	var b Book
	err := ctx.ReadBody(&b)
	// TIP: use ctx.ReadBody(&b) to bind
	// any type of incoming data instead.
	fmt.Println(b)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Book creation failure").DetailErr(err))
		// TIP: use ctx.StopWithError(code, err) when only
		// plain text responses are expected on errors.
		return
	}

	Books = append(Books, b)
	println("Received Book: " + b.Title)

	ctx.StatusCode(iris.StatusCreated)
}
