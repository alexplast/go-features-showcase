package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-features-showcase/features"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	err = db.AutoMigrate(&features.Person{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return &Server{db: db}
}

func (s *Server) Run() {
	r := gin.Default()

	r.GET("/people", s.GetPeople)
	r.POST("/people", s.CreatePerson)
	r.GET("/people/:id", s.GetPerson)
	r.PUT("/people/:id", s.UpdatePerson)
	r.DELETE("/people/:id", s.DeletePerson)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
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
