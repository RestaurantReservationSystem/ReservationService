package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"reservation_service/config"
	pb "reservation_service/genproto"
	"reservation_service/service"
	"reservation_service/storage/postgres"
)

func main() {
	cnf := config.Load()

	db, err := postgres.ConnectionDb()
	if err != nil {
		log.Fatalf("error-> connection db", err.Error())
	}

	listen, err := net.Listen("tcp", cnf.HTTPPort)
	if err != nil {
		log.Fatalf("error is listen up tcp connection in port->", err.Error())
	}
	grpcServer := grpc.NewServer()
	pb.RegisterReservationServiceServer(grpcServer, service.NewReservationService(postgres.NewReservationRepository(db), postgres.NewMenuRepository(db), postgres.NewRestaurantRepo(db), postgres.NewOrderRepository(db)))
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("error-> connection ->%s", err.Error())
	}
}
