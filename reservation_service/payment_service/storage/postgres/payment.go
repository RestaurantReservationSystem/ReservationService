package postgres

import (
	"database/sql"
	pb "payment_service/genproto"
	help "payment_service/help"
	"time"
)

type PaymentRepository struct {
	Db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{Db: db}
}

func (repo *PaymentRepository) CreatePayment(request *pb.CreatePaymentRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("insert into  payments(reservation_id,amount,payment_method,payment_status,created_at) values ($1,$2,$3,$4,$5)", request.ReservationId, request.Amount, request.PaymentMethod, request.PaymentStatus, time.Now())
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, err
}
func (repo *PaymentRepository) UpdatePayment(request *pb.UpdatePaymentRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update  payments set reservation_id=$1,amount=$2,payment_method=$3,payment_status=$4,updated_at=$5 where id=$6 and deleted_at is null", request.ReservationId, request.Amount, request.PaymentMethod, request.PaymentStatus, time.Now(), request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, err
}

func (repo *PaymentRepository) DeletedPayment(request *pb.IdRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update  payments set deleted_at=$1 where id=$2 and deleted_at is null ", time.Now(), request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, err
}
func (repo *PaymentRepository) GetByIdPayment(request *pb.IdRequest) (*pb.PaymentResponse, error) {
	var response pb.PaymentResponse
	err := repo.Db.QueryRow("select reservation_id, amount,payment_method,payment_status from payments where  id=$1 and delted_at is null ", request.Id).Scan(&response.ReservationId, &response.Amount, &response.PaymentMethod, &response.PaymentStatus)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
func (repo *PaymentRepository) GetAllPayment(request *pb.GetAllPaymentRequest) (*pb.PaymentsResponse, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	filter := ""
	if len(request.ReservationId) > 0 {
		params["reservation_id"] = request.ReservationId
		filter += " and reservation_id = :reservation_id "
	}
	if request.Amount > 0 {
		params["amount"] = request.Amount
		filter += " and amount:=amount"
	}
	if len(request.PaymentMethod) > 0 {
		params["payment_method"] = request.PaymentStatus
		filter += "and payment_method:=payment_method"
	}
	if len(request.PaymentStatus) > 0 {
		params["payment_status"] = request.PaymentStatus
		filter += "and payment_status:=payment_status"
	}
	if request.LimitOffset.Limit > 0 {
		params["limit"] = request.LimitOffset.Limit
		filter += "and limit:=limit"
	}
	if request.LimitOffset.Offset > 0 {
		params["offset"] = request.LimitOffset.Offset
		filter += "and offset:=offset"
	}
	query := "select user_id,restaurant_id,reservation_time,status from reservations where deleted_at is null"
	query = query + filter + limit + offset
	query, arr = help.ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	var payments []*pb.PaymentResponse

	for rows.Next() {
		var payment pb.PaymentResponse
		err := rows.Scan(&payment.ReservationId, &payment.Amount, &payment.PaymentMethod, &payment.PaymentStatus)
		if err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}
	return &pb.PaymentsResponse{Payments: payments}, nil
}
