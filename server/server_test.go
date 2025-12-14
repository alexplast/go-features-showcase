package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go-features-showcase/features"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestServer initializes a server with an in-memory SQLite database
// and seeds it with data for testing.
func setupTestServer() *Server {
	// Use a new in-memory SQLite database for each test setup.
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Drop table to ensure a clean state
	_ = db.Migrator().DropTable(&features.Person{})

	err = db.AutoMigrate(&features.Person{})
	if err != nil {
		panic("failed to migrate test database")
	}

	// Seed database with test data
	people := []features.Person{
		{Name: "John", Age: 30},
		{Name: "Jane", Age: 25},
		{Name: "Johnathan", Age: 35},
		{Name: "Zoe", Age: 28},
	}
	db.Create(&people)

	return &Server{db: db}
}

// TestGetPeople tests the default retrieval of people with pagination.
func TestGetPeople(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := setupTestServer()
	router := gin.Default()
	router.GET("/people", server.GetPeople)

	req, _ := http.NewRequestWithContext(t.Context(), http.MethodGet, "/people", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data  []features.Person `json:"data"`
		Total int               `json:"total"`
		Page  int               `json:"page"`
		Limit int               `json:"limit"`
	}

	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, 4, response.Total)
	assert.Equal(t, 1, response.Page)
	assert.Equal(t, 10, response.Limit)
	assert.Len(t, response.Data, 4)
	assert.Equal(t, "John", response.Data[0].Name)
}

// TestGetPeoplePagination tests the pagination feature.
func TestGetPeoplePagination(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := setupTestServer()
	router := gin.Default()
	router.GET("/people", server.GetPeople)

	req, _ := http.NewRequestWithContext(t.Context(), http.MethodGet, "/people?page=2&limit=2", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data  []features.Person `json:"data"`
		Total int64             `json:"total"`
		Page  int               `json:"page"`
	}

	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, int64(4), response.Total)
	assert.Equal(t, 2, response.Page)
	assert.Len(t, response.Data, 2)
	assert.Equal(t, "Johnathan", response.Data[0].Name) // Assuming default order is by ID
}

// TestGetPeopleFilter tests the name filtering feature.
func TestGetPeopleFilter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := setupTestServer()
	router := gin.Default()
	router.GET("/people", server.GetPeople)

	req, _ := http.NewRequestWithContext(t.Context(), http.MethodGet, "/people?name=John", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data  []features.Person `json:"data"`
		Total int64             `json:"total"`
	}

	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, int64(2), response.Total)
	assert.Len(t, response.Data, 2)
	assert.Equal(t, "John", response.Data[0].Name)
	assert.Equal(t, "Johnathan", response.Data[1].Name)
}

// TestGetPeopleFilterAndPagination tests combining filtering and pagination.
func TestGetPeopleFilterAndPagination(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := setupTestServer()
	router := gin.Default()
	router.GET("/people", server.GetPeople)

	req, _ := http.NewRequestWithContext(t.Context(), http.MethodGet, "/people?name=John&page=2&limit=1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data  []features.Person `json:"data"`
		Total int64             `json:"total"`
		Page  int               `json:"page"`
	}

	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, int64(2), response.Total)
	assert.Equal(t, 2, response.Page)
	assert.Len(t, response.Data, 1)
	assert.Equal(t, "Johnathan", response.Data[0].Name)
}

func TestCreatePerson(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := setupTestServer()
	router := gin.Default()
	router.POST("/people", server.CreatePerson)

	person := features.Person{Name: "NewGuy", Age: 25}
	body, _ := json.Marshal(person)
	req, _ := http.NewRequestWithContext(t.Context(), http.MethodPost, "/people", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var createdPerson features.Person
	err := json.Unmarshal(w.Body.Bytes(), &createdPerson)
	require.NoError(t, err)
	assert.Equal(t, "NewGuy", createdPerson.Name)
}

func TestGetPerson(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := setupTestServer()
	router := gin.Default()
	router.GET("/people/:id", server.GetPerson)

	var person features.Person
	server.db.First(&person, "name = ?", "John")

	req, _ := http.NewRequestWithContext(t.Context(), http.MethodGet, "/people/"+strconv.Itoa(int(person.ID)), nil)
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
	server.db.First(&person, "name = ?", "John")

	updatedPerson := features.Person{Name: "John Doe", Age: 31}
	body, _ := json.Marshal(updatedPerson)
	req, _ := http.NewRequestWithContext(t.Context(), http.MethodPut, "/people/"+strconv.Itoa(int(person.ID)), bytes.NewBuffer(body))
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
	server.db.First(&person, "name = ?", "John")

	req, _ := http.NewRequestWithContext(t.Context(), http.MethodDelete, "/people/"+strconv.Itoa(int(person.ID)), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Person deleted")

	// Verify the person was actually deleted
	var deletedPerson features.Person
	err := server.db.First(&deletedPerson, person.ID).Error
	assert.Error(t, err) // Should be an error (record not found)
}
