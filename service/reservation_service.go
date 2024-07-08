package service

import (
	"context"
	pb "reservation_service/genproto"
	"reservation_service/storage/postgres"
)

type ReservationService struct {
	ReservationRepo *postgres.ReservationRepository
	pb.UnimplementedReservationServiceServer
}

func NewReservationService(repo *postgres.ReservationRepository) *ReservationService {
	return &ReservationService{ReservationRepo: repo}
}
func (service *ReservationService) CreateReservation(ctx context.Context, in *pb.CreateReservationRequest) (*pb.Void, error) {
	return service.ReservationRepo.CreateReservation(in)
}
func (service *ReservationService) UpdateReservation(ctx context.Context, in *pb.UpdateReservationRequest) (*pb.Void, error) {
	return service.ReservationRepo.UpdateReservation(in)
}
func (service *ReservationService) DeleteReservation(ctx context.Context, in *pb.IdRequest) (*pb.Void, error) {
	return service.ReservationRepo.DeleteReservation(in)
}
func (service *ReservationService) GetByIdReservation(ctx context.Context, in *pb.IdRequest) (*pb.ReservationResponse, error) {
	return service.ReservationRepo.GetByIdReservation(in)
}
func (service *ReservationService) GetAllReservation(ctx context.Context, in *pb.GetAllReservationRequest) (*pb.ReservationsResponse, error) {
	return service.ReservationRepo.GetAllReservation(in)
}
