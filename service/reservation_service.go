package service

import (
	"context"
	pb "reservation_service/genproto"
	"reservation_service/storage/postgres"
)

type ReservationService struct {
	ReservationRepo *postgres.ReservationRepository
	MenuRepo        *postgres.MenuRepository
	RestaurantRepo  *postgres.RestaurantRepository
	OrderRepo       *postgres.OrderRepository
	pb.UnimplementedReservationServiceServer
}

func NewReservationService(reservation *postgres.ReservationRepository, menu *postgres.MenuRepository, restaurant *postgres.RestaurantRepository, order *postgres.OrderRepository) *ReservationService {
	return &ReservationService{ReservationRepo: reservation, MenuRepo: menu, RestaurantRepo: restaurant, OrderRepo: order}
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

func (service *ReservationService) CreateMenu(ctx context.Context, in *pb.CreateMenuRequest) (*pb.Void, error) {
	return service.MenuRepo.CreateMenu(in)
}
func (service *ReservationService) UpdateMenu(ctx context.Context, in *pb.UpdateMenuRequest) (*pb.Void, error) {
	return service.MenuRepo.UpdateMenu(in)
}
func (service *ReservationService) DeleteMenu(ctx context.Context, in *pb.IdRequest) (*pb.Void, error) {
	return service.MenuRepo.DeleteMenu(in)
}
func (service *ReservationService) GetByIdMenu(ctx context.Context, in *pb.IdRequest) (*pb.MenuResponse, error) {
	return service.MenuRepo.GetByIdMenu(in)
}
func (service *ReservationService) GetAllMenu(ctx context.Context, in *pb.GetAllMenuRequest) (*pb.MenusResponse, error) {
	return service.MenuRepo.GetAllMenu(in)
}

func (service *ReservationService) CreateRestaurant(ctx context.Context, in *pb.CreateRestaurantRequest) (*pb.Void, error) {
	return service.RestaurantRepo.CreateRestaurant(in)
}
func (service *ReservationService) UpdateRestaurant(ctx context.Context, in *pb.UpdateRestaurantRequest) (*pb.Void, error) {
	return service.RestaurantRepo.UpdateRestaurant(in)
}
func (service *ReservationService) DeleteRestaurant(ctx context.Context, in *pb.IdRequest) (*pb.Void, error) {
	return service.RestaurantRepo.DeleteRestaurant(in)
}
func (service *ReservationService) GetByIdRestaurant(ctx context.Context, in *pb.IdRequest) (*pb.RestaurantResponse, error) {
	return service.RestaurantRepo.GetByIdRestaurant(in)
}
func (service *ReservationService) GetAllRestaurants(ctx context.Context, in *pb.GetAllRestaurantRequest) (*pb.RestaurantsResponse, error) {
	return service.RestaurantRepo.GetAllRestaurants(in)
}

func (service *ReservationService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.Void, error) {
	return service.OrderRepo.CreateOrder(in)
}
func (service *ReservationService) UpdateOrder(ctx context.Context, in *pb.UpdateOrderRequest) (*pb.Void, error) {
	return service.OrderRepo.UpdateOrder(in)
}
func (service *ReservationService) DeleteOrder(ctx context.Context, in *pb.IdRequest) (*pb.Void, error) {
	return service.OrderRepo.DeleteOrder(in)
}
func (service *ReservationService) GetByIdOrder(ctx context.Context, in *pb.IdRequest) (*pb.OrderResponse, error) {
	return service.OrderRepo.GetByIdOrder(in)
}
func (service *ReservationService) GetAllOrder(ctx context.Context, in *pb.GetAllOrderRequest) (*pb.OrdersResponse, error) {
	return service.OrderRepo.GetAllOrder(in)
}
