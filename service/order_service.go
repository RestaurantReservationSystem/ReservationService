package service

import (
	"context"
	pb "reservation_service/genproto"
	"reservation_service/storage/postgres"
)

type OrderService struct {
	OrderRepo *postgres.OrderRepository
	pb.UnimplementedOrderServiceServer
}

func NewOrderService(repo *postgres.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: repo}
}

func (service *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.Void, error) {
	return service.OrderRepo.CreateOrder(in)
}
func (service *OrderService) UpdateOrder(ctx context.Context, in *pb.UpdateOrderRequest) (*pb.Void, error) {
	return service.OrderRepo.UpdateOrder(in)
}
func (service *OrderService) DeleteOrder(ctx context.Context, in *pb.IdRequest) (*pb.Void, error) {
	return service.OrderRepo.DeleteOrder(in)
}
func (service *OrderService) GetByIdOrder(ctx context.Context, in *pb.IdRequest) (*pb.OrderResponse, error) {
	return service.OrderRepo.GetByIdOrder(in)
}
func (service *OrderService) GetAllOrder(ctx context.Context, in *pb.GetAllOrderRequest) (*pb.OrdersResponse, error) {
	return service.OrderRepo.GetAllOrder(in)
}
