package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	app := fiber.New()

	r, _ := initDB()

	app.Get("/books", r.GetBooks)
	req, _ := http.NewRequest(http.MethodGet, "/books", nil)
	res, err := app.Test(req, -1)

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	t.Log(string(body))

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Nil(t, err)
	assert.NotEmpty(t, string(body))
	fmt.Println(string(body))
}

func TestGetBookByID(t *testing.T) {
	app := fiber.New()
	r, _ := initDB()
	app.Get("/get_books/:id", r.GetBookByID)
	req, _ := http.NewRequest(http.MethodGet, "/get_books/4", nil)
	res, err := app.Test(req, -1)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var book Book
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	err = json.Unmarshal(body, &book)

	fmt.Println(&book)
	assert.Nil(t, err)
	assert.NotNil(t, string(body))

}

func TestCreateBook(t *testing.T) {
	app := fiber.New()
	r, _ := initDB()
	app.Post("/create_books", r.CreateBook)

	//var jsonStr = []byte(`{"author":"thisumnee","title": "Book 1","publisher": "gunasena"}`)

	values := map[string]string{"author": "thisumnee", "title": "Book 1", "publisher": "gunasena"}
	jsonValue, _ := json.Marshal(values)

	req, err := http.NewRequest(http.MethodPost, "/create_books", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)

	var response struct {
		Data    Book   `json:"data"`
		Message string `json:"message"`
	}

	if err := json.Unmarshal([]byte(body), &response); err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	fmt.Println(response.Data)

	assert.Nil(t, err)
	assert.NotNil(t, response.Data)
	assert.Equal(t, "Book 1", response.Data.Title)
}

func TestDeleteBook(t *testing.T) {
	app := fiber.New()
	r, _ := initDB()
	app.Delete("/delete_books/:id", r.DeleteBook)
	req, err := http.NewRequest(http.MethodDelete, "/delete_books/5", nil)
	response, err := app.Test(req, -1)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)

}
