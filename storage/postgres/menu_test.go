package postgres

import (
	"reflect"
	pb "reservation_service/genproto"
	"reservation_service/storage"
	"testing"
)

func ReservationRepo(t *testing.T) *MenuRepository {
	db, err := storage.ConnectionDb()
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
