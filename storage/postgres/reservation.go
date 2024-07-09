package postgres

import (
	"database/sql"
	pb "reservation_service/genproto"
	storage "reservation_service/help"
	"time"
)

type ReservationRepository struct {
	Db *sql.DB
}

func NewReservationRepository(db *sql.DB) *ReservationRepository {
	return &ReservationRepository{Db: db}
}

func (repo *ReservationRepository) CreateReservation(request *pb.CreateReservationRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("insert into reservations(user_id,restaurant_id,reservation_time,status,created_at) values ($1,$2,$3,$4,$5)", request.UserId, request.RestaurantId, request.ReservationTime, request.Status, time.Now())
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, err
}
func (repo *ReservationRepository) UpdateReservation(request *pb.UpdateReservationRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update reservations set user_id=$1,restaurant_id=$2,reservation_time=$3,status=$4,updated_at=$5 where id=$6 and deleted_at is null", request.UserId, request.RestaurantId, request.ReservationTime, request.Status, time.Now(), request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, err
}
func (repo *ReservationRepository) DeleteReservation(request *pb.IdRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update reservations set deleted_at=$1 where id=$2 and deleted_at is null", time.Now(), request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, err
}
func (repo *ReservationRepository) GetByIdReservation(request *pb.IdRequest) (*pb.ReservationResponse, error) {
	var response pb.ReservationResponse
	err := repo.Db.QueryRow("select user_id,restaurant_id,reservation_time,status from reservations where deleted_at is null and id=$1", request.Id).Scan(&response.UserId, &response.RestaurantId, &response.ReservationTime, &response.Status)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (repo *ReservationRepository) GetAllReservation(request *pb.GetAllReservationRequest) (*pb.ReservationsResponse, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	filter := ""
	if len(request.UserId) > 0 {
		params["user_id"] = request.UserId
		filter += " and user_id = :user_id "
	}
	if len(request.RestaurantId) > 0 {
		params["restaurant_id"] = request.RestaurantId
		filter += " and restaurant_id:=restaurant_id"
	}
	if len(request.ReservationTime) > 0 {
		params["reservation_time"] = request.ReservationTime
		filter += "and reservation_time:=reservation_time"
	}
	if len(request.Status) > 0 {
		params["status"] = request.Status
		filter += "and status:=status"
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
	query, arr = storage.ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	var reservations []*pb.ReservationResponse

	for rows.Next() {
		var response pb.ReservationResponse
		err := rows.Scan(&response.UserId, &response.RestaurantId, &response.ReservationTime, &response.Status)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, &response)
	}
	return &pb.ReservationsResponse{Reservations: reservations}, nil
}
