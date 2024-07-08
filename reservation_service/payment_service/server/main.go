package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"payment_service/config"
	pb "payment_service/genproto"
	"payment_service/service"
	"payment_service/storage/postgres"
)

func main() {
	db, err := postgres.Connection()
	if err != nil {
		log.Fatalf("error-> database connection error->%s", err.Error())
	}
	cnf := config.Load()
	listen, err := net.Listen("tcp", cnf.HTTPPort)
	if err != nil {
		log.Fatalf("error-> tcp connection error->%s", err.Error())
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, service.NewPaymentService(postgres.NewPaymentRepository(db)))
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("error-> api_get_way connection error->%s", err.Error())
	}

}
