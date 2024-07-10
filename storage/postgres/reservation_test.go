package postgres

import (
	"fmt"
	"reflect"
	pb "reservation_service/genproto"
	"reservation_service/storage"
	"testing"
)

func ReservationRep(t *testing.T) *ReservationRepository {
	db, err := storage.ConnectionDb()
	if err != nil {
		t.Error("ERROR : ", err)
		return nil
	}
	userRepo := c(db)
	return userRepo
}

func TestReservationRepository_CreateReservation(t *testing.T) {
	Reservation := ReservationRep(t)

	reser := pb.CreateReservationRequest{
		UserId:          "a12d4f5e-0e9c-42b8-847a-3c54b7f6a4c2",
		RestaurantId:    "30.00",
		ReservationTime: "credit_card",
		Status:          "345rtg",
	}
	response, err := Reservation.CreateReservation(&reser)
	if err != nil {
		fmt.Println("_________", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.ReservationResponse{})

	}
}
func TestReservationRepository_DeleteReservation(t *testing.T) {
	Reservation := ReservationRep(t)
	resDelete := pb.IdRequest{
		Id: "b8811cf2-9352-4dc3-9884-c7dabca7ab8b",
	}
	response, err := Reservation.DeleteReservation(&resDelete)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.ReservationsResponse{})

	}

}
func TestReservationRepository_UpdateReservation(t *testing.T) {
	Reservation := ReservationRep(t)

	reserUpdate := pb.UpdateReservationRequest{
		UserId:          "a12d4f5e-0e9c-42b8-847a-3c54b7f6a4c2",
		RestaurantId:    "30.00",
		ReservationTime: "credit_card",
		Status:          "345rtg",
	}
	response, err := Reservation.UpdateReservation(&reserUpdate)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.ReservationResponse{})

	}
}

func TestReservationRepository_GetReservation(t *testing.T) {
	Reservation := ReservationRep(t)

	getReservation := pb.IdRequest{
		Id: "e7c84f17-6a39-4889-92fe-0a9f3a3c062",
	}

	expected := pb.ReservationResponse{
		UserId:          "a12d4f5e-0e9c-42b8-847a-3c54b7f6a4c2",
		RestaurantId:    "30.00",
		ReservationTime: "credit_card",
		Status:          "345rtg",
	}

	response, err := Reservation.GetByIdReservation(&getReservation)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &expected) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.ReservationResponse{})

	}
}

func TestAllReservationRepository_GetAllReservation(t *testing.T) {
	Reservation := ReservationRep(t)

	getAllReservations := pb.GetAllReservationRequest{
		UserId:          "a12d4f5e-0e9c-42b8-847a-3c54b7f6a4c2",
		RestaurantId:    "30.00",
		ReservationTime: "credit_card",
		Status:          "345rtg",
	}

	expected := pb.ReservationResponse{}

	response, err := Reservation.GetAllReservation(&getAllReservations)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &expected) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.ReservationResponse{})

	}
}
