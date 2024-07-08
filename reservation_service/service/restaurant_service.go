package service

import (
	"context"
	pb "reservation_service/genproto"
	"reservation_service/storage/postgres"
)

type RestaurantService struct {
	RestaurantRepo *postgres.RestaurantRepository
	pb.UnimplementedRestaurantServiceServer
}

func NewRestaurantService(repo *postgres.RestaurantRepository) *RestaurantService {
	return &RestaurantService{RestaurantRepo: repo}
}

func (service *RestaurantService) CreateRestaurant(ctx context.Context, in *pb.CreateRestaurantRequest) (*pb.Void, error) {
	return service.RestaurantRepo.CreateRestaurant(in)
}
func (service *RestaurantService) UpdateRestaurant(ctx context.Context, in *pb.UpdateRestaurantRequest) (*pb.Void, error) {
	return service.RestaurantRepo.UpdateRestaurant(in)
}
func (service *RestaurantService) DeleteRestaurant(ctx context.Context, in *pb.IdRequest) (*pb.Void, error) {
	return service.RestaurantRepo.DeleteRestaurant(in)
}
func (service *RestaurantService) GetByIdRestaurant(ctx context.Context, in *pb.IdRequest) (*pb.RestaurantResponse, error) {
	return service.RestaurantRepo.GetByIdRestaurant(in)
}
func (service *RestaurantService) GetAllRestaurants(ctx context.Context, in *pb.GetAllRestaurantRequest) (*pb.RestaurantsResponse, error) {
	return service.RestaurantRepo.GetAllRestaurants(in)
}
