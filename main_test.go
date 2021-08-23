package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/toDo-list/internal/handlers"
	"log"
	"net/http"
	"net/http/httptest"
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

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
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

	if resp.StatusCode != 201 {
		t.Fatalf("Expected status code 201, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}
