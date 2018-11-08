package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.DELETE("/emptylist", emptyBookList)
	e.DELETE("/removebook/:id", removeBook)

	e.GET("/booklist", getBookList)
	e.GET("/book/:id", getBook)

	e.POST("/addbook", addBook)
	e.POST("/setrating/:id", setBookRating)
	e.POST("/setstatus/:id", setBookStatus)

	e.Logger.Fatal(e.Start(":1323"))
}
