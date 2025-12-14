package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"go-features-showcase/features"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestServer() *Server {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&features.Person{})
	if err != nil {
		panic("failed to migrate test database")
	}
	db.Create(&features.Person{Name: "John", Age: 30})

	return &Server{db: db}
}

func TestGetPeople(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := setupTestServer()
	router := gin.Default()
	router.GET("/people", server.GetPeople)

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/people", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"name":"John"`)
}

func TestCreatePerson(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := setupTestServer()
	router := gin.Default()
	router.POST("/people", server.CreatePerson)

	person := features.Person{Name: "Jane", Age: 25}
	body, _ := json.Marshal(person)
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/people", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"name":"Jane"`)
}

func TestGetPerson(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := setupTestServer()
	router := gin.Default()
	router.GET("/people/:id", server.GetPerson)

	var person features.Person
	server.db.First(&person)

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/people/"+strconv.Itoa(int(person.ID)), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"name":"John"`)
}

func TestUpdatePerson(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := setupTestServer()
	router := gin.Default()
	router.PUT("/people/:id", server.UpdatePerson)

	var person features.Person
	server.db.First(&person)

	updatedPerson := features.Person{Name: "John Doe", Age: 31}
	body, _ := json.Marshal(updatedPerson)
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPut, "/people/"+strconv.Itoa(int(person.ID)), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"name":"John Doe"`)
}

func TestDeletePerson(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := setupTestServer()
	router := gin.Default()
	router.DELETE("/people/:id", server.DeletePerson)

	var person features.Person
	server.db.First(&person)

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodDelete, "/people/"+strconv.Itoa(int(person.ID)), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Person deleted")
}
