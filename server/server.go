package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-features-showcase/features"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

type Server struct {
	db *gorm.DB
}

func NewServer() *Server {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&features.Person{})

	return &Server{db: db}
}

func (s *Server) Run() {
	r := gin.Default()

	r.GET("/people", s.GetPeople)
	r.POST("/people", s.CreatePerson)
	r.GET("/people/:id", s.GetPerson)
	r.PUT("/people/:id", s.UpdatePerson)
	r.DELETE("/people/:id", s.DeletePerson)

	r.Run()
}

func (s *Server) GetPeople(c *gin.Context) {
	var people []features.Person
	s.db.Find(&people)
	c.JSON(http.StatusOK, people)
}

func (s *Server) CreatePerson(c *gin.Context) {
	var person features.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.db.Create(&person)
	c.JSON(http.StatusOK, person)
}

func (s *Server) GetPerson(c *gin.Context) {
	var person features.Person
	if err := s.db.First(&person, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, person)
}

func (s *Server) UpdatePerson(c *gin.Context) {
	var person features.Person
	if err := s.db.First(&person, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.db.Save(&person)
	c.JSON(http.StatusOK, person)
}

func (s *Server) DeletePerson(c *gin.Context) {
	var person features.Person
	if err := s.db.First(&person, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	s.db.Delete(&person)
	c.JSON(http.StatusOK, gin.H{"message": "Person deleted"})
}
