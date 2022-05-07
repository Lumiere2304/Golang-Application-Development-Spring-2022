package main

import (
	"context"
	"flag"
	"fmt"
	"grpcServer/pkg/api"
	"log"

	"google.golang.org/grpc"
)

const requiredArguments int = 2

func main() {
	flag.Parse()
	if flag.NArg() < requiredArguments {
		log.Fatal("You passed not enough arguements")
	}

	inputText := flag.Arg(0)

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	newClient := api.NewAdderClient(conn)

	result, err := newClient.Reverse(context.Background(), &api.Message{Text: inputText})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.GetText())
}