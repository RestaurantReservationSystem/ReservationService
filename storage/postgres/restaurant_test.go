package postgres

import (
	"reflect"
	pb "reservation_service/genproto"
	"testing"
)

func RestaurantRepo(t *testing.T) *RestaurantRepository {
	db, err := ConnectionDb()
	if err != nil {
		t.Error("ERROR : ", err)
		return nil
	}
	userRepo := NewRestaurantRepo(db)
	return userRepo
}

func TestCreateRestaurant(t *testing.T) {
	RestaurantRepo := RestaurantRepo(t)

	request := pb.CreateRestaurantRequest{Name: "Rayhon", Address: "Chilonzor", PhoneNumber: "+998954674546", Description: "Oilaviy resatoran"}

	res, err := RestaurantRepo.CreateRestaurant(&request)
	if err != nil {
		t.Error("Error : ", err)
		return
	}

	case1 := pb.Void{}

	if !reflect.DeepEqual(res, &case1) {
		t.Error("Result : ", res, "Expected : ", &case1)
		return
	}
}

func TestUpdateRestaurant(t *testing.T) {
	RestaurantRepo := RestaurantRepo(t)

	updateReq := pb.UpdateRestaurantRequest{Id: "92d120c7-56c7-4c51-a1e1-f46d006308eb", Name: "Nordon", Address: "Shuhrat", PhoneNumber: "+998991234567", Description: "Hamma uchun"}

	res, err := RestaurantRepo.UpdateRestaurant(&updateReq)
	if err != nil {
		t.Error("Error : ", err)
		return
	}

	case1 := pb.Void{}

	if !reflect.DeepEqual(res, &case1) {
		t.Error("Result : ", res, "Expected : ", &case1)
		return
	}
}

func TestDeleteRestaurant(t *testing.T) {
	restaurantRepo := RestaurantRepo(t)

	id := pb.IdRequest{Id: "6ba3b94d-dbaf-4242-8f2b-faa223813029"}

	res, err := restaurantRepo.DeleteRestaurant(&id)
	if err != nil {
		t.Error("Error : ", err)
		return
	}

	case1 := pb.Void{}

	if !reflect.DeepEqual(res, &case1) {
		t.Error("Result : ", res, "Expected : ", &case1)
		return
	}
}

func TestGetByIdRestaurant(t *testing.T) {
	restaurantRepo := RestaurantRepo(t)

	id := pb.IdRequest{Id: "b00dbf91-fecc-4419-9bfe-1dac7c1ed8c6"}

	res, err := restaurantRepo.GetByIdRestaurant(&id)
	if err != nil {
		t.Error("Error : ", err)
		return
	}

	case1 := pb.RestaurantResponse{Name: "Rayhon", Address: "Chilonzor", PhoneNumber: "+998954674546", Description: "Oilaviy resatoran"}

	if !reflect.DeepEqual(res, &case1) {
		t.Error("Result : ", res, "Expected : ", &case1)
		return
	}
}

func TestGetAllRestaurants(t *testing.T) {
	restaurantRepo := RestaurantRepo(t)

	req := pb.GetAllRestaurantRequest{Address: "Chilonzor", LimitOffset: &pb.Filter{Limit: 1}}

	res, err := restaurantRepo.GetAllRestaurants(&req)
	if err != nil {
		t.Error("Error : ", err)
		return
	}

	expectedRestaurant := &pb.RestaurantResponse{
		Name:        "Rayhon",
		Address:     "Chilonzor",
		PhoneNumber: "+998954674546",
	}

	case1 := pb.RestaurantsResponse{Restaurants: []*pb.RestaurantResponse{expectedRestaurant}}

	if reflect.DeepEqual(res, &case1) {
		t.Error("Result : ", res, "Expected : ", &case1)
		return
	}
}
