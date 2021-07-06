package main

import (
	"context"
	"fmt"
	"go-gRpc/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}
func (*server) Greet(ctx context.Context, res *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("%v \n", res)
	firstName := res.GetGreeting().GetFirstName()
	lastName := res.GetGreeting().GetLastName()
	result := "Hello " + firstName + " " + lastName
	resp := &greetpb.GreetResponse{
		Result: result,
	}
	return resp, nil
}


func main()  {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis) ; err != nil {
		log.Fatalf("Error: %v", err)
	}

}
