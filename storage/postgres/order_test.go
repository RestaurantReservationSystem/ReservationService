package postgres

import (
	"fmt"
	"reflect"
	pb "reservation_service/genproto"
	"reservation_service/storage"
	"testing"
)

func OrderRepo(t *testing.T) *OrderRepository {
	db, err := storage.ConnectionDb()
	if err != nil {
		t.Error("ERROR : ", err)
		return nil
	}
	userRepo := NewOrderRepository(db)
	return userRepo
}

func TestOrderRepository_CreateOrder(t *testing.T) {
	RestaurantRepo := OrderRepo(t)

	ords := pb.CreateOrderRequest{
		ReservationId: "a12d4f5e-0e9c-42b8-847a-3c54b7f6a4c2",
		MenuItemId:    "30.00",
		Quantity:      "credit_card",
	}
	response, err := RestaurantRepo.CreateOrder(&ords)
	if err != nil {
		fmt.Println("_________", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.OrderResponse{})

	}
}
func TestOrderRepository_DeleteOrder(t *testing.T) {
	RestaurantRepo := OrderRepo(t)

	ordDelete := pb.IdRequest{
		Id: "b8811cf2-9352-4dc3-9884-c7dabca7ab8b",
	}
	response, err := RestaurantRepo.DeleteOrder(&ordDelete)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.OrderResponse{})

	}

}
func TestOrderRepository_UpdateOrder(t *testing.T) {
	RestaurantRepo := OrderRepo(t)

	orderUpdate := pb.UpdateOrderRequest{
		Id:            "b271e2a3-f90d-4b51-bfd5-808bd65f1756",
		ReservationId: "d3dcbdff-de1c-452d-94da-2bb783f1016a",
		MenuItemId:    "753.5",
		Quantity:      "salom",
	}
	response, err := RestaurantRepo.UpdateOrder(&orderUpdate)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.OrderResponse{})

	}
}

func TestOrderRepository_GetOrder(t *testing.T) {
	RestaurantRepo := OrderRepo(t)

	getOrder := pb.IdRequest{
		Id: "e7c84f17-6a39-4889-92fe-0a9f3a3c062",
	}

	expected := pb.OrderResponse{
		ReservationId: "a12d4f5e-0e9c-42b8-847a-3c54b7f6a4c2",
		MenuItemId:    "credit_card",
		Quantity:      "pending",
	}

	response, err := RestaurantRepo.GetByIdOrder(&getOrder)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &expected) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.OrderResponse{})

	}
}

func TestAllOrderRepository_GetAllOrder(t *testing.T) {
	RestaurantRepo := OrderRepo(t)

	getAllOrders := pb.GetAllOrderRequest{
		ReservationId: "d3dcbdff-de1c-452d-94da-2bb783f1016a",
		MenuItemId:    "salom",
		Quantity:      "+++++++++++++++++++++++",
	}

	expected := pb.PaymentsResponse{}

	response, err := RestaurantRepo.GetAllOrder(&getAllOrders)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &expected) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.OrderResponse{})

	}
}
