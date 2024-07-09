package postgres

import (
	"database/sql"
	pb "reservation_service/genproto"
	storage "reservation_service/help"
	"time"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (repo *OrderRepository) CreateReservation(request *pb.CreateOrderRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("insert into orders(reservation_id,menu_item_id,quantity,created_at) values ($1,$2,$3,$4)", request.ReservationId, request.MenuItemId, request.Quantity, time.Now())
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, err
}
func (repo *OrderRepository) UpdateReservation(request *pb.UpdateOrderRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update orders set reservation_id=$1,menu_item_id=$2,quantity=$3,updated_at=$4 where id=$5 and deleted_at is null", request.ReservationId, request.MenuItemId, request.Quantity, time.Now(), request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, err
}
func (repo *OrderRepository) DeleteOrder(request *pb.IdRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update orders set deleted_at=$1 where id=$2 and deleted_at is null", time.Now(), request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, err
}
func (repo *OrderRepository) GetByIdOrder(request *pb.IdRequest) (*pb.OrderResponse, error) {
	var response pb.OrderResponse
	err := repo.Db.QueryRow("select reservation_id,menu_item_id,quantity  from orders where deleted_at is null and id=$1", request.Id).Scan(&response.ReservationId, &response.MenuItemId, &response.Quantity)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (repo *OrderRepository) GetAllOrder(request *pb.GetAllOrderRequest) (*pb.OrdersResponse, error) {
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
	if len(request.MenuItemId) > 0 {
		params["menu_item_id"] = request.MenuItemId
		filter += " and menu_item_id:=menu_item_id"
	}
	if len(request.Quantity) > 0 {
		params["quantity"] = request.Quantity
		filter += "and quantity:=quantity"
	}

	if request.LimitOffset.Limit > 0 {
		params["limit"] = request.LimitOffset.Limit
		filter += "and limit:=limit"
	}
	if request.LimitOffset.Offset > 0 {
		params["offset"] = request.LimitOffset.Offset
		filter += "and offset:=offset"
	}
	query := "select reservation_id,menu_item_id,quantity from orders where deleted_at is null"

	query = query + filter + limit + offset
	query, arr = storage.ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	var orders []*pb.OrderResponse

	for rows.Next() {
		var order pb.OrderResponse
		err := rows.Scan(&order.ReservationId, &order.MenuItemId, &order.Quantity)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return &pb.OrdersResponse{Orders: orders}, nil
}
