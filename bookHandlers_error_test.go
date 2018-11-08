package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestAddBookError(t *testing.T) {
	e := echo.New()

	formValues := make(url.Values)
	formValues.Set("Author", "Jeff VanderMeer")
	formValues.Set("PublishDate", "2014-11-18")
	formValues.Set("Publisher", "FSG Originals")
	formValues.Set("Rating", "3")
	formValues.Set("Status", "Checked-In")
	formValues.Set("Title", "Area-X")

	addBookReq := httptest.NewRequest(
		http.MethodPost,
		"/addbook",
		strings.NewReader(formValues.Encode()),
	)
	addBookReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	addBookRec := httptest.NewRecorder()
	addBookContext := e.NewContext(addBookReq, addBookRec)

	if assert.NoError(t, addBook(addBookContext)) {
		assert.Equal(t, "Could not add Area-X", addBookRec.Body.String())
	}
}

func TestGetBookError(t *testing.T) {
	e := echo.New()

	getBookReq := httptest.NewRequest(http.MethodGet, "/getbook/:id", nil)
	getBookRec := httptest.NewRecorder()
	getBookContext := e.NewContext(getBookReq, getBookRec)
	getBookContext.SetParamNames("id")
	getBookContext.SetParamValues("JeffVanderMeer-Area-X")

	if assert.NoError(t, getBook(getBookContext)) {
		assert.Equal(t, "No book with id JeffVanderMeer-Area-X", getBookRec.Body.String())
	}
}

func TestRemoveBookError(t *testing.T) {
	e := echo.New()

	removeBookReq := httptest.NewRequest(http.MethodDelete, "/removebook/:id", nil)
	removeBookRec := httptest.NewRecorder()
	removeBookContext := e.NewContext(removeBookReq, removeBookRec)
	removeBookContext.SetParamNames("id")
	removeBookContext.SetParamValues("JeffVanderMeer-Area-X")

	if assert.NoError(t, removeBook(removeBookContext)) {
		assert.Equal(t, "Could not remove book with id JeffVanderMeer-Area-X", removeBookRec.Body.String())
	}
}

func TestSetBookRatingError(t *testing.T) {
	e := echo.New()

	formValues := make(url.Values)
	formValues.Set("Rating", "2")
	setRatingReq := httptest.NewRequest(
		http.MethodPost,
		"/setrating/:id",
		strings.NewReader(formValues.Encode()),
	)
	setRatingReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	setRatingRec := httptest.NewRecorder()
	setRatingContext := e.NewContext(setRatingReq, setRatingRec)
	setRatingContext.SetParamNames("id")
	setRatingContext.SetParamValues("JeffVanderMeer-Area-X")

	if assert.NoError(t, setBookRating(setRatingContext)) {
		assert.Equal(t, "No book to update with id JeffVanderMeer-Area-X", setRatingRec.Body.String())
	}
}

func TestSetBookStatusError(t *testing.T) {
	e := echo.New()

	formValues := make(url.Values)
	formValues.Set("Status", "Checked-Out")
	setRatingReq := httptest.NewRequest(
		http.MethodPost,
		"/setrating/:id",
		strings.NewReader(formValues.Encode()),
	)
	setRatingReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	setRatingRec := httptest.NewRecorder()
	setRatingContext := e.NewContext(setRatingReq, setRatingRec)
	setRatingContext.SetParamNames("id")
	setRatingContext.SetParamValues("JeffVanderMeer-Area-X")

	if assert.NoError(t, setBookStatus(setRatingContext)) {
		assert.Equal(t, "No book to update with id JeffVanderMeer-Area-X", setRatingRec.Body.String())
	}
}
