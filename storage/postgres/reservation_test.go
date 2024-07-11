package postgres

import (
	"fmt"
	"reflect"
	pb "reservation_service/genproto"
	"testing"
)

func TestReservationRepository_CreateReservation(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	fmt.Println("+++++++++++++")

	res := NewReservationRepository(db)
	fmt.Println(res)
	reser := pb.CreateReservationRequest{
		UserId:          "a12d4f5e-0e9c-42b8-847a-3c54b7f6a4c2",
		RestaurantId:    "30.00",
		ReservationTime: "credit_card",
		Status:          "345rtg",
	}
	response, err := res.CreateReservation(&reser)
	if err != nil {
		fmt.Println("_________", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.ReservationResponse{})

	}
}
func TestReservationRepository_DeleteReservation(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {

		panic(err)
	}

	res := NewReservationRepository(db)
	fmt.Println(res)
	resDelete := pb.IdRequest{
		Id: "b8811cf2-9352-4dc3-9884-c7dabca7ab8b",
	}
	response, err := res.DeleteReservation(&resDelete)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.ReservationsResponse{})

	}

}
func TestReservationRepository_UpdateReservation(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {

		panic(err)
	}

	res := NewReservationRepository(db)
	fmt.Println(res)
	reserUpdate := pb.UpdateReservationRequest{
		UserId:          "a12d4f5e-0e9c-42b8-847a-3c54b7f6a4c2",
		RestaurantId:    "30.00",
		ReservationTime: "credit_card",
		Status:          "345rtg",
	}
	response, err := res.UpdateReservation(&reserUpdate)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.ReservationResponse{})

	}
}

func TestReservationRepository_GetReservation(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {

		panic(err)
	}

	res := NewReservationRepository(db)

	getReservation := pb.IdRequest{
		Id: "e7c84f17-6a39-4889-92fe-0a9f3a3c062",
	}

	fmt.Println(res)
	expected := pb.ReservationResponse{
		UserId:          "a12d4f5e-0e9c-42b8-847a-3c54b7f6a4c2",
		RestaurantId:    "30.00",
		ReservationTime: "credit_card",
		Status:          "345rtg",
	}

	response, err := res.GetByIdReservation(&getReservation)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &expected) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.ReservationResponse{})

	}
}

func TestAllReservationRepository_GetAllReservation(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {

		panic(err)
	}

	reservation := NewReservationRepository(db)

	getAllReservations := pb.GetAllReservationRequest{
		UserId:          "a12d4f5e-0e9c-42b8-847a-3c54b7f6a4c2",
		RestaurantId:    "30.00",
		ReservationTime: "credit_card",
		Status:          "345rtg",
	}

	fmt.Println(reservation)
	expected := pb.ReservationResponse{}

	response, err := reservation.GetAllReservation(&getAllReservations)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &expected) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.ReservationResponse{})

	}
}
