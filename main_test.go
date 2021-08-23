package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/toDo-list/internal/handlers"
	"github.com/toDo-list/internal/models"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func setupGetServer() *gin.Engine {
	r := gin.Default()

	r.GET("/todos", handlers.GetTodoList)

	return r
}

func setupPostServer() *gin.Engine {
	r := gin.Default()

	r.POST("/todo", handlers.CreateTodo)

	return r
}

// CDC Tests
func TestGetTodoListRoute(t *testing.T) {
	ts := httptest.NewServer(setupGetServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/todos", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		var todoModelArray []models.Todo
		err = json.Unmarshal(bodyBytes, &todoModelArray)
		if err != nil {
			t.Fatal(err)
		}

		if reflect.TypeOf(todoModelArray) != reflect.TypeOf([]models.Todo{}) {
			t.Fatalf("Expected %s, got %s", reflect.TypeOf([]models.Todo{}), reflect.TypeOf(todoModelArray))
		}
	}
}

func TestCreateTodoRoute(t *testing.T) {
	ts := httptest.NewServer(setupPostServer())
	defer ts.Close()

	data := map[string]string{"description": "test"}
	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(fmt.Sprintf("%s/todo", ts.URL), "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code 201, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}

	if resp.StatusCode == http.StatusCreated {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		var todoModel models.Todo
		err = json.Unmarshal(bodyBytes, &todoModel)
		if err != nil {
			t.Fatal(err)
		}

		if reflect.TypeOf(todoModel) != reflect.TypeOf(models.Todo{}) {
			t.Fatalf("Expected %s, got %s", reflect.TypeOf([]models.Todo{}), reflect.TypeOf(todoModel))
		}
	}
}
