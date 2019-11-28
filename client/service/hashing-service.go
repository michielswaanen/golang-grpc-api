package service

import (
	"../../proto"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
)

// Handle the Encode http request
func Encode(context *gin.Context) {

	// Connect with gRPC and spin-up a new client service
	client := newHashingServiceClient()

	// Get parameter from url
	plainText := context.Param("plainText")

	// Create a new hashing encode request
	req := &proto.HashingEncodeRequest{PlainText: plainText}

	// Encode and send response
	if response, err := client.Encode(context, req); err == nil {
		context.JSON(http.StatusOK, gin.H{
			"encoded": fmt.Sprint(response.HashedText),
		})
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
}

// Handle the Decode http request
func Decode(context *gin.Context) {

	// Connect with gRPC and spin-up a new client service
	client := newHashingServiceClient()

	// Get parameter from url
	hash := context.Param("hashedText")

	// Create a new hashing decode request
	req := &proto.HashingDecodeRequest{HashedText: hash}

	// Decode and send response
	if response, err := client.Decode(context, req); err == nil {
		context.JSON(http.StatusOK, gin.H{
			"decoded": fmt.Sprint(response.PlainText),
		})
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}

// Private function that creates a new hashing service when the client connects
func newHashingServiceClient() proto.HashingServiceClient {

	// Connect (Dial) to the server with an insecure connection (HTTP)
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	return proto.NewHashingServiceClient(conn)
}
