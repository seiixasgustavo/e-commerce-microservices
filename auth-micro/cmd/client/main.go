package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	v1 "github.com/seiixasgustavo/e-commerce-microservices/auth-micro/pkg/api/v1"
	"google.golang.org/grpc"
)

func main() {

	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Problem opening .env file: %v", err)
	}

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Problems dialing to grpc server: %v", err)
	}

	client := v1.NewAuthClient(conn)
	clientUser := v1.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, rErr := client.SignUp(ctx, &v1.UserAuthRequest{
		Username: "gustavo",
		Email:    "gustavo@gustavo.com",
		Password: "gustavo",
	})

	if rErr != nil {
		log.Fatalf("Server problem: %v", rErr)
	}
	fmt.Printf("Server responded: %v\n", res.GetStatus())

	resL, lErr := client.Login(ctx, &v1.LoginRequest{
		Username: "gustavo",
		Password: "gustavo",
	})

	if lErr != nil {
		log.Fatalf("Server problem: %v", lErr)
	}
	fmt.Printf("Server responded: %v\n", resL.GetStatus())

	res1, err1 := clientUser.Create(ctx, &v1.UserRequest{
		User: &v1.UserStruct{
			Username: "gustavo1",
			Email:    "gustavo1@gustavo.com",
			Password: "gustavo",
		},
	})

	if err1 != nil {
		log.Fatalf("Server problem: %v", err1)
	}
	fmt.Printf("Server responded: %v\n", res1.GetStatus())

	res2, err2 := clientUser.Delete(ctx, &v1.IdRequest{
		ID: uint64(1),
	})

	if err2 != nil {
		log.Fatalf("Server problem: %v", err2)
	}
	fmt.Printf("Server responded: %v\n", res2.GetStatus())

	res3, err3 := clientUser.Update(ctx, &v1.UserIdRequest{
		User: &v1.UserStruct{
			Username: "gustavo2",
			Email:    "gustavo2@gustavo.com",
			Password: "gustavo",
		},
		ID: uint64(2),
	})

	if err3 != nil {
		log.Fatalf("Server problem: %v", err3)
	}
	fmt.Printf("Server responded: %v\n", res3.GetStatus())

	res4, err4 := clientUser.ChangePassword(ctx, &v1.PasswordRequest{
		ID:       uint64(2),
		Password: "gustavo1",
	})

	if err4 != nil {
		log.Fatalf("Server problem: %v", err4)
	}
	fmt.Printf("Server responded: %v\n", res4.GetStatus())

	res5, err5 := clientUser.FindByPk(ctx, &v1.IdRequest{
		ID: uint64(2),
	})

	if err5 != nil {
		log.Fatalf("Server problem: %v", err5)
	}
	fmt.Printf("Server responded: %v - %v\n", res5.GetStatus(), res5.GetUser())

	res6, err6 := clientUser.FindByUsername(ctx, &v1.UsernameRequest{
		Username: "gustavo1",
	})

	if err6 != nil {
		log.Fatalf("Server problem: %v", err6)
	}
	fmt.Printf("Server responded: %v - %v\n", res6.GetStatus(), res6.GetUser())

}
