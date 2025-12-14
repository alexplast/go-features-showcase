package features

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"strings"
)

// Function that demonstrates error handling
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func DemonstrateErrorHandling() {
	log.Info("\n--- Error Handling ---")
	divResult, err := Divide(10, 2)
	if err != nil {
		log.Error("Error:", err)
	} else {
		log.Info("10 / 2 =", divResult)
	}

	divResult, err = Divide(10, 0)
	if err != nil {
		log.Error("Error:", err)
	} else {
		log.Info("10 / 0 =", divResult)
	}
}

// Function to demonstrate the defer statement
func DemonstrateDefer() {
	defer log.Info("\nThis will be printed last.")
	log.Info("\nThis will be printed first.")
}

func DemonstratePointers() {
	log.Info("\n--- Pointers ---")
	x := 10
	y := &x
	log.Infof("x = %d, y = %p", x, y)
	*y = 20
	log.Infof("x = %d, y = %p", x, y)
}

func DemonstrateStringManipulation() {
	log.Info("\n--- String Manipulation ---")
	myString := "Hello, Go!"
	log.Info("Original string:", myString)
	log.Info("Uppercase:", strings.ToUpper(myString))
	log.Info("Contains 'Go':", strings.Contains(myString, "Go"))
}