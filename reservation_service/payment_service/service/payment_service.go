package service

import (
	"context"
	pb "payment_service/genproto"
	"payment_service/storage/postgres"
)

type PaymentService struct {
	PaymentRepo *postgres.PaymentRepository
	pb.UnimplementedPaymentServiceServer
}

func NewPaymentService(repo *postgres.PaymentRepository) *PaymentService {
	return &PaymentService{PaymentRepo: repo}
}
func (service *PaymentService) CreatePayment(ctx context.Context, in *pb.CreatePaymentRequest) (*pb.Void, error) {
	return service.PaymentRepo.CreatePayment(in)
}
func (service *PaymentService) UpdatePayment(ctx context.Context, in *pb.UpdatePaymentRequest) (*pb.Void, error) {
	return service.PaymentRepo.UpdatePayment(in)
}
func (service *PaymentService) DeletePayment(ctx context.Context, in *pb.IdRequest) (*pb.Void, error) {
	return service.PaymentRepo.DeletedPayment(in)
}
func (service *PaymentService) GetByIdPayment(ctx context.Context, in *pb.IdRequest) (*pb.PaymentResponse, error) {
	return service.PaymentRepo.GetByIdPayment(in)
}
func (service *PaymentService) GetAllPayment(ctx context.Context, in *pb.GetAllPaymentRequest) (*pb.PaymentsResponse, error) {
	return service.PaymentRepo.GetAllPayment(in)
}
