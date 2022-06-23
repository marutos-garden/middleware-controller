package main

import (
	"context"
	"log"
	pb "github.com/marutos/protobuf-files/device_service"
)

func doGreet(c pb.DeviceServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Device(context.Background(), &pb.DeviceRequest{
		FirstName: "Marwan",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
