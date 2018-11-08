package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
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

	id := strings.Replace(author, " ", "", -1) + "-" + title

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
		return context.String(http.StatusBadRequest, "Could not add "+title)
	}

	returnString := title + " by " + author + " has been added"
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

func setBookRating(context echo.Context) error {
	client := getRedis()
	id := context.Param("id")
	rating, ratingErr := strconv.Atoi(context.FormValue("Rating"))

	if ratingErr != nil {
		return context.String(http.StatusBadRequest, "No rating provided")
	}

	bookString, err := client.Get(id).Result()
	if err != nil {
		return context.String(http.StatusBadRequest, "No book to update with id "+id)
	}

	var bookJSON Book
	json.Unmarshal([]byte(bookString), &bookJSON)

	book := Book{
		Author:      bookJSON.Author,
		PublishDate: bookJSON.PublishDate,
		Publisher:   bookJSON.Publisher,
		Rating:      rating,
		Status:      bookJSON.Status,
		Title:       bookJSON.Title,
	}

	newBookJSON, _ := json.Marshal(book)

	setErr := client.Set(id, newBookJSON, 0).Err()
	if setErr != nil {
		fmt.Println(err)
		return context.String(http.StatusBadRequest, "Could not update "+bookJSON.Title)
	}

	returnString := "Successfully set rating for " + bookJSON.Title + " to " + strconv.Itoa(rating)
	return context.String(http.StatusOK, returnString)
}

func setBookStatus(context echo.Context) error {
	client := getRedis()
	id := context.Param("id")
	status := context.FormValue("Status")

	if status == "" {
		return context.String(http.StatusBadRequest, "No status provided")
	}

	bookString, err := client.Get(id).Result()
	if err != nil {
		return context.String(http.StatusBadRequest, "No book to update with id "+id)
	}

	var bookJSON Book
	json.Unmarshal([]byte(bookString), &bookJSON)

	book := Book{
		Author:      bookJSON.Author,
		PublishDate: bookJSON.PublishDate,
		Publisher:   bookJSON.Publisher,
		Rating:      bookJSON.Rating,
		Status:      status,
		Title:       bookJSON.Title,
	}

	newBookJSON, _ := json.Marshal(book)

	setErr := client.Set(id, newBookJSON, 0).Err()
	if setErr != nil {
		fmt.Println(err)
		return context.String(http.StatusBadRequest, "Could not update "+bookJSON.Title)
	}

	returnString := "Successfully set status for " + bookJSON.Title + " to " + status
	return context.String(http.StatusOK, returnString)
}
