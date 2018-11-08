package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestAddBook(t *testing.T) {
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
		assert.Equal(t, "Area-X by Jeff VanderMeer has been added", addBookRec.Body.String())
	}
}

func TestEmptyList(t *testing.T) {
	e := echo.New()

	clearDBReq := httptest.NewRequest(http.MethodDelete, "/emptylist", nil)
	clearDBRec := httptest.NewRecorder()
	clearDBContext := e.NewContext(clearDBReq, clearDBRec)

	if assert.NoError(t, emptyBookList(clearDBContext)) {
		assert.Equal(t, "Book List Empty", clearDBRec.Body.String())
	}
}

func TestGetBook(t *testing.T) {
	e := echo.New()

	getBookReq := httptest.NewRequest(http.MethodGet, "/getbook/:id", nil)
	getBookRec := httptest.NewRecorder()
	getBookContext := e.NewContext(getBookReq, getBookRec)
	getBookContext.SetParamNames("id")
	getBookContext.SetParamValues("JeffVanderMeer-Area-X")

	if assert.NoError(t, getBook(getBookContext)) {
		var book string
		json.Unmarshal([]byte(getBookRec.Body.String()), &book)
		expected := "{\"Author\":\"Jeff VanderMeer\",\"PublishDate\":\"2014-11-18T00:00:00Z\",\"Publisher\":\"FSG Originals\",\"Rating\":3,\"Status\":\"Checked-In\",\"Title\":\"Area-X\"}"
		assert.Equal(t, expected, book)
	}
}

func TestRemoveBook(t *testing.T) {
	e := echo.New()

	removeBookReq := httptest.NewRequest(http.MethodDelete, "/removebook/:id", nil)
	removeBookRec := httptest.NewRecorder()
	removeBookContext := e.NewContext(removeBookReq, removeBookRec)
	removeBookContext.SetParamNames("id")
	removeBookContext.SetParamValues("JeffVanderMeer-Area-X")

	if assert.NoError(t, removeBook(removeBookContext)) {
		assert.Equal(t, "Successfully removed book", removeBookRec.Body.String())
	}
}
