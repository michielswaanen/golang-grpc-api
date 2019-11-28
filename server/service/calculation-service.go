package service

import (
	"../../proto"
	"context"
)

type calculationServer struct{}

// Add 2 integers together and return result
func (s *calculationServer) Add(ctx context.Context, request *proto.CalculationRequest) (*proto.CalculationResponse, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.CalculationResponse{Result: result}, nil
}

// Multiply 2 integers together and return result
func (s *calculationServer) Multiply(ctx context.Context, request *proto.CalculationRequest) (*proto.CalculationResponse, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto.CalculationResponse{Result: result}, nil
}

// Create a new instance of the calculation server
func InitializeCalculationServer() *calculationServer {
	return new(calculationServer)
}
