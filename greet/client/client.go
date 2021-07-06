package main

import (
	"context"
	"fmt"
	"go-gRpc/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cc,err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)

}
func doUnary(c greetpb.GreetServiceClient){
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Abdulmalik",
			LastName: "Giwa :) ",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil{
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("%v", res.Result)
}
