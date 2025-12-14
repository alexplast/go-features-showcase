package features

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math"
)

// Struct to represent a Person
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Method for the Person struct
func (p Person) Greet() {
	greeting := viper.GetString("greeting")
	log.Infof("%s My name is %s and I am %d years old.", greeting, p.Name, p.Age)
}

// Interface for shapes
type Shaper interface {
	Area() float64
}

// Struct for a Circle
type Circle struct {
	Radius float64
}

// Method for Circle to implement the Shaper interface
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Struct for a Rectangle
type Rectangle struct {
	Width, Height float64
}

// Method for Rectangle to implement the Shaper interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func DemonstrateStructsAndInterfaces() {
	log.Info("--- Structs and Methods ---")
	p := Person{Name: "John", Age: 30}
	p.Greet()

	log.Info("\n--- Interfaces ---")
	shapes := []Shaper{
		Circle{Radius: 5},
		Rectangle{Width: 10, Height: 5},
	}
	for _, shape := range shapes {
		log.Infof("Area of shape: %f", shape.Area())
	}
}

func DemonstrateJSON() {
	log.Info("\n--- JSON Marshaling ---")
	p := Person{Name: "John", Age: 30}
	personJSON, _ := json.Marshal(p)
	log.Info("Person as JSON:", string(personJSON))
}

func DemonstrateSlices() {
	log.Info("\n--- Slices ---")
	mySlice := []int{1, 2, 3, 4, 5}
	log.Info("Slice:", mySlice)
	mySlice = append(mySlice, 6)
	log.Info("Appended slice:", mySlice)
}

func DemonstrateMaps() {
	log.Info("\n--- Maps ---")
	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	log.Info("Map:", myMap)
	log.Info("Value of 'one':", myMap["one"])
	delete(myMap, "one")
	log.Info("Map after deleting 'one':", myMap)
}