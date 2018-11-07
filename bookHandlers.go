package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func addBook(context echo.Context) error {
	client := getRedis()

	author := context.FormValue("Author")
	publishDate, pubDateErr := time.Parse("2006-01-02", context.FormValue("PublishDate"))
	publisher := context.FormValue("Publisher")
	rating, ratingErr := strconv.Atoi(context.FormValue("Rating"))
	status := context.FormValue("Status")
	title := context.FormValue("Title")

	if ratingErr != nil {
		rating = 2
	}
	if pubDateErr != nil {
		publishDate = time.Now()
	}

	id := uuid.New().String()

	book := Book{
		Author:      author,
		PublishDate: publishDate,
		Publisher:   publisher,
		Rating:      rating,
		Status:      status,
		Title:       title,
	}

	bookJSON, _ := json.Marshal(book)

	err := client.Set(id, bookJSON, 0).Err()
	if err != nil {
		fmt.Println(err)
		return context.String(http.StatusBadRequest, "Could not add "+title)
	}

	returnString := title + " by " + author + " has been added with id " + id
	return context.String(http.StatusOK, returnString)
}

func emptyBookList(context echo.Context) error {
	getRedis().FlushDB()

	return context.String(http.StatusOK, "Book List Empty")
}

func getBook(context echo.Context) error {
	client := getRedis()

	id := context.Param("id")

	bookJSON, err := client.Get(id).Result()
	if err != nil {
		return context.String(http.StatusBadRequest, "No book with id "+id)
	}

	// Want to parse the json better that comes back from redis
	return context.JSON(http.StatusOK, bookJSON)
}

func getBookList(context echo.Context) error {
	client := getRedis()
	// Not sure why this doesn't work yet
	keys := client.Scan(0, "*", 100)

	return context.String(http.StatusOK, keys.String())
}

func getRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func removeBook(context echo.Context) error {
	client := getRedis()

	id := context.Param("id")

	err := client.Del(id).Err()
	if err != nil {
		return context.String(http.StatusOK, "Could not remove book with id "+id)
	}

	return context.String(http.StatusOK, "Successfully removed book")
}
