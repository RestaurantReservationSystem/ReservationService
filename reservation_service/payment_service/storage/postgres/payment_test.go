package postgres

import (
	"fmt"
	"log"
	"log/slog"
	pb "payment_service/genproto"
	"reflect"
	"testing"
)

//	func TestConnection(t *testing.T) {
//		db, err := Connection()
//		if err != nil {
//			panic(err)
//		}
//
// }
func TestPaymentRepository_CreatePayment(t *testing.T) {
	db, err := Connection()
	if err != nil {
		log.Fatal("error is connection db")
	}
	repo := NewPaymentRepository(db)
	request := pb.CreatePaymentRequest{PaymentStatus: "pending", PaymentMethod: "card", Amount: 4.5, ReservationId: "1234alk;sdl;kf"}
	response, err := repo.CreatePayment(&request)
	if err != nil {
		fmt.Println("+++++++", err)
		//log.Fatalf("error internal server error ->", err.Error())
	}
	expected := &pb.Void{}
	if reflect.DeepEqual(response, expected) {
		slog.StringValue("message is success")
	}
}
