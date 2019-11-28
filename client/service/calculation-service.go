package service

import (
	"../../proto"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

// Handle the Add http request
func Add(context *gin.Context) {

	// Connect with gRPC and spin-up a new client service
	client := newCalculationServiceClient()

	// Parse and handle error
	a := getParameterFromURL("a", context)
	b := getParameterFromURL("b", context)

	// Create a new calculation add request
	req := &proto.CalculationRequest{A: a, B: b}

	// Call the add service and send result
	if response, err := client.Add(context, req); err == nil {
		context.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(response.Result),
		})
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}

// Handle the Multiply http request
func Multiply(context *gin.Context) {

	// Connect with gRPC and spin-up a new client service
	client := newCalculationServiceClient()

	// Parse and handle error
	a := getParameterFromURL("a", context)
	b := getParameterFromURL("b", context)

	// Create a new calculation multiply request
	req := &proto.CalculationRequest{A: int64(a), B: int64(b)}

	// Call the add service and send result
	if response, err := client.Multiply(context, req); err == nil {
		context.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(response.Result),
		})
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}

// Private function because it starts with a lower case
func getParameterFromURL(wildcard string, context *gin.Context) int64 {
	parsedWildcard, err := strconv.ParseInt(context.Param(wildcard), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter" + wildcard})
		return -1
	}
	return parsedWildcard
}

func newCalculationServiceClient() proto.CalculationServiceClient {

	// Connect (Dial) to the server with an insecure connection (HTTP)
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	// Create a new client that wants to connect to the proto connection
	return proto.NewCalculationServiceClient(conn)
}
