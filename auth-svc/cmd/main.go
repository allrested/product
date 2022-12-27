package main

import (
	"fmt"
	"github.com/allrested/product/auth-svc/pkg/config"
	"github.com/allrested/product/auth-svc/pkg/db"
	pb "github.com/allrested/product/auth-svc/pkg/pb"
	services "github.com/allrested/product/auth-svc/pkg/services"
	"github.com/allrested/product/auth-svc/pkg/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          c.Issuer,
		ExpirationHours: c.ExpHours,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
