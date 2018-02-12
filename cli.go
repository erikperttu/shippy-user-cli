package main

import (
	"log"
	"os"

	"context"

	pb "github.com/erikperttu/shippy-user-service/proto/auth"
	"github.com/micro/go-micro"
	microClient "github.com/micro/go-micro/client"
)

func main() {
	srv := micro.NewService(
		micro.Name("shippy.user-cli"),
		micro.Version("latest"),
	)

	srv.Init()

	client := pb.NewAuthClient("shippy.auth", microClient.DefaultClient)

	name := "Test Name"
	email := "test@example.com"
	password := "test123"
	company := "Very large ships inc."

	r, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})

	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	log.Printf("Created: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Failed to list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.Fatalf("Failed to authenticate user: %s error: %v\n", email, err)
	}

	log.Printf("Your access token is: %s \n", authResponse.Token)

	// We good
	os.Exit(0)
}
