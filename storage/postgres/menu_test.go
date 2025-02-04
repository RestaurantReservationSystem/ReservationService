package postgres

import (
	"fmt"
	"reflect"
	pb "reservation_service/genproto"
	"testing"
)

func ReservationRepo(t *testing.T) *MenuRepository {
	db, err := ConnectionDb()
	if err != nil {
		t.Error("ERROR : ", err)
		return nil
	}
	userRepo := NewMenuRepository(db)
	return userRepo
}

func TestCreateMenu(t *testing.T) {
	reservationRepo := ReservationRepo(t)

	request := pb.CreateMenuRequest{
		RestaurantId: "b00dbf91-fecc-4419-9bfe-1dac7c1ed8c6",
		Name:         "Palov",
		Description:  "Mazali va Qazili",
		Price:        100000,
	}

	res, err := reservationRepo.CreateMenu(&request)
	if err != nil {
		t.Error("ERROR : ", err)
		return
	}

	case1 := pb.Void{}

	if !reflect.DeepEqual(res, &case1) {
		t.Error("Result : ", res, "Expected : ", &case1)
		return
	}
}

func TestUpdateMenu(t *testing.T) {
	reservationRepo := ReservationRepo(t)

	request := pb.UpdateMenuRequest{
		Id:           "46e59237-a0e5-49a4-be53-bc6349978f31",
		RestaurantId: "b00dbf91-fecc-4419-9bfe-1dac7c1ed8c6",
		Name:         "Osh",
		Description:  "Mazali va Qazili",
		Price:        120000,
	}

	res, err := reservationRepo.UpdateMenu(&request)
	if err != nil {
		t.Error("ERROR : ", err)
		return
	}

	case1 := pb.Void{}

	if !reflect.DeepEqual(res, &case1) {
		t.Error("Result : ", res, "Expected : ", &case1)
		return
	}
}

func TestDeleteMenu(t *testing.T) {
	reservationRepo := ReservationRepo(t)

	request := pb.IdRequest{Id: "8bcfda5a-df4b-476f-a89f-a0d623350589"}

	res, err := reservationRepo.DeleteMenu(&request)
	if err != nil {
		t.Error("ERROR : ", err)
		return
	}

	case1 := pb.Void{}

	if !reflect.DeepEqual(res, &case1) {
		t.Error("Result : ", res, "Expected : ", &case1)
		return
	}
}

func TestGetByIdMenu(t *testing.T) {
	reservationRepo := ReservationRepo(t)

	request := pb.IdRequest{Id: "46e59237-a0e5-49a4-be53-bc6349978f31"}

	res, err := reservationRepo.GetByIdMenu(&request)
	if err != nil {

		t.Error("ERROR : ", err)
		return
	}

	case1 := pb.MenuResponse{
		RestaurantId: "b00dbf91-fecc-4419-9bfe-1dac7c1ed8c6",
		Name:         "Osh",
		Description:  "Mazali va Qazili",
		Price:        120000,
	}

	if !reflect.DeepEqual(res, &case1) {
		t.Error("Result : ", res, "Expected : ", &case1)
		return
	}
}

func TestGetAllMenu(t *testing.T) {
	reservationRepo := ReservationRepo(t)

	request := pb.GetAllMenuRequest{
		RestaurantId: "b00dbf91-fecc-4419-9bfe-1dac7c1ed8c6",
		Name:         "Osh",
		Description:  "Mazali va Qazili",
		Price:        120000,
	}

	res, err := reservationRepo.GetAllMenu(&request)
	if err != nil {

		t.Error("ERROR : ", err)
		return
	}

	expectedfood := &pb.MenuResponse{
		RestaurantId: "b00dbf91-fecc-4419-9bfe-1dac7c1ed8c6",
		Name:         "Osh",
		Description:  "Mazali va Qazili",
		Price:        120000,
	}

	case1 := pb.MenusResponse{Menus: []*pb.MenuResponse{expectedfood}}

	if !reflect.DeepEqual(res, &case1) {
		fmt.Println("+++++++++++++++", err)
		t.Error("Result : ", res, "Expected : ", &case1)
		return
	}
}
