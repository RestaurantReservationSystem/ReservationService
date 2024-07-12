package postgres

import (
	"fmt"
	"reflect"
	pb "reservation_service/genproto"
	"testing"
)

func TestOrderRepository_CreateOrder(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	fmt.Println("+++++++++++++")

	ord := NewOrderRepository(db)
	fmt.Println(ord)
	ords := pb.CreateOrderRequest{
		ReservationId: "a12d4f5e-0e9c-42b8-847a-3c54b7f6a4c2",
		MenuItemId:    "30.00",
		Quantity:      89,
	}
	response, err := ord.CreateOrder(&ords)
	if err != nil {
		fmt.Println("_________", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.OrderResponse{})

	}
}
func TestOrderRepository_DeleteOrder(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {

		panic(err)
	}

	pay := NewOrderRepository(db)
	fmt.Println(pay)
	ordDelete := pb.IdRequest{
		Id: "b8811cf2-9352-4dc3-9884-c7dabca7ab8b",
	}
	response, err := pay.DeleteOrder(&ordDelete)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.OrderResponse{})

	}

}
func TestOrderRepository_UpdateOrder(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {

		panic(err)
	}

	ord := NewOrderRepository(db)
	fmt.Println(ord)
	orderUpdate := pb.UpdateOrderRequest{
		Id:            "b271e2a3-f90d-4b51-bfd5-808bd65f1756",
		ReservationId: "d3dcbdff-de1c-452d-94da-2bb783f1016a",
		MenuItemId:    "753.5",
		Quantity:      12,
	}
	response, err := ord.UpdateOrder(&orderUpdate)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.OrderResponse{})

	}
}

func TestOrderRepository_GetOrder(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {

		panic(err)
	}

	pay := NewOrderRepository(db)

	getOrder := pb.IdRequest{
		Id: "e7c84f17-6a39-4889-92fe-0a9f3a3c062",
	}

	fmt.Println(pay)
	expected := pb.OrderResponse{
		ReservationId: "a12d4f5e-0e9c-42b8-847a-3c54b7f6a4c2",
		MenuItemId:    "credit_card",
		Quantity:      89,
	}

	response, err := pay.GetByIdOrder(&getOrder)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &expected) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.OrderResponse{})

	}
}

func TestAllOrderRepository_GetAllOrder(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {

		panic(err)
	}

	payment := NewOrderRepository(db)

	getAllOrders := pb.GetAllOrderRequest{
		ReservationId: "d3dcbdff-de1c-452d-94da-2bb783f1016a",
		MenuItemId:    "salom",
		Quantity:      89,
	}

	fmt.Println(payment)
	expected := pb.PaymentsResponse{}

	response, err := payment.GetAllOrder(&getAllOrders)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &expected) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.OrderResponse{})

	}
}
