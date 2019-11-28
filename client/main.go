package main

import (
	"./service"
	"github.com/gin-gonic/gin"
	"log"
)

// Start the client
func main() {

	// Initialize Routing library
	router := gin.Default()

	// Define all the routes that start with localhost:8080/calculation
	calculation := router.Group("/calculation")
	{
		calculation.GET("/add/:a/:b", service.Add)
		calculation.GET("/multiply/:a/:b", service.Multiply)
	}

	// Define all the routes that start with localhost:8080/hashing
	hashing := router.Group("/hashing")
	{
		hashing.GET("/encode/:plainText", service.Encode)
		hashing.GET("/decode/:hashedText", service.Decode)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
