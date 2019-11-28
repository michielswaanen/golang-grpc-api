package service

import (
	"../../proto"
	"context"
	"encoding/base64"
)

type hashingServer struct{}

// Encode a string into a hashed piece of text
func (s *hashingServer) Encode(ctx context.Context, request *proto.HashingEncodeRequest) (*proto.HashingEncodeResponse, error) {
	plainText := request.GetPlainText()

	hashedText := base64.StdEncoding.EncodeToString([]byte(plainText))

	return &proto.HashingEncodeResponse{HashedText: hashedText}, nil
}

// Decode a string into a plain piece of text
func (s *hashingServer) Decode(ctx context.Context, request *proto.HashingDecodeRequest) (*proto.HashingDecodeResponse, error) {
	hashedText := request.GetHashedText()

	bytePlainText, err := base64.StdEncoding.DecodeString(hashedText)

	if err == nil {
		return &proto.HashingDecodeResponse{PlainText: string(bytePlainText)}, nil
	} else {
		return &proto.HashingDecodeResponse{PlainText: err.Error()}, nil
	}
}

// Create a new instance of the hashing server
func InitializeHashingServer() *hashingServer {
	return new(hashingServer)
}
