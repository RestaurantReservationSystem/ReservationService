package service

import (
	"context"
	pb "reservation_service/genproto"
	"reservation_service/storage/postgres"
)

type MenuService struct {
	MenuRepo *postgres.MenuRepository
	pb.UnimplementedMenuServiceServer
}

func NewMenuService(repo *postgres.MenuRepository) *MenuService {
	return &MenuService{MenuRepo: repo}
}

func (service *MenuService) CreateMenu(ctx context.Context, in *pb.CreateMenuRequest) (*pb.Void, error) {
	return service.MenuRepo.CreateMenu(in)
}
func (service *MenuService) UpdateMenu(ctx context.Context, in *pb.UpdateMenuRequest) (*pb.Void, error) {
	return service.MenuRepo.UpdateMenu(in)
}
func (service *MenuService) DeleteMenu(ctx context.Context, in *pb.IdRequest) (*pb.Void, error) {
	return service.MenuRepo.DeleteMenu(in)
}
func (service *MenuService) GetByIdMenu(ctx context.Context, in *pb.IdRequest) (*pb.MenuResponse, error) {
	return service.MenuRepo.GetByIdMenu(in)
}
func (service *MenuService) GetAllMenu(ctx context.Context, in *pb.GetAllMenuRequest) (*pb.MenusResponse, error) {
	return service.MenuRepo.GetAllMenu(in)
}
