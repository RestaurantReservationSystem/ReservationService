package postgres

import (
	"database/sql"
	pb "reservation_service/genproto"
	"reservation_service/help"
	"time"
)

type MenuRepository struct {
	Db *sql.DB
}

func NewMenuRepository(db *sql.DB) *MenuRepository {
	return &MenuRepository{Db: db}
}

func (repo *MenuRepository) CreateMenu(request *pb.CreateMenuRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("insert into menus(restaurant_id,name,description,price,created_at) values ($1,$2,$3,$4,$5)", request.RestaurantId, request.Name, request.Description, request.Price, time.Now())
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, err
}
func (repo *MenuRepository) UpdateMenu(request *pb.UpdateMenuRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update menus set restaurant_id=$1,name=$2,description=$3,price=$4,updated_at=$5 where id=$6 and deleted_at is null", request.RestaurantId, request.Name, request.Description, request.Price, time.Now(), request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, err
}
func (repo *MenuRepository) DeleteMenu(request *pb.IdRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update menus set deleted_at=$1 where id=$2 and deleted_at is null", time.Now(), request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, err
}
func (repo *MenuRepository) GetByIdMenu(request *pb.IdRequest) (*pb.MenuResponse, error) {
	var response pb.MenuResponse
	err := repo.Db.QueryRow("select restaurant_id,name,description,price from menus where deleted_at is null and id=$1", request.Id).Scan(&response.RestaurantId, &response.Name, &response.Description, &response.Price)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (repo *MenuRepository) GetAllMenu(request *pb.GetAllMenuRequest) (*pb.MenusResponse, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	filter := ""
	if len(request.RestaurantId) > 0 {
		params["restaurant_id"] = request.RestaurantId
		filter += " and restaurant_id = :restaurant_id "
	}
	if len(request.Name) > 0 {
		params["name"] = request.Name
		filter += " and name:=name"
	}
	if len(request.Description) > 0 {
		params["description"] = request.Description
		filter += "and description:=description"
	}
	if request.Price > 0 {
		params["price"] = request.Price
		filter += "and price:=price"
	}
	if request.LimitOffset.Limit > 0 {
		params["limit"] = request.LimitOffset.Limit
		filter += "and limit:=limit"
	}
	if request.LimitOffset.Offset > 0 {
		params["offset"] = request.LimitOffset.Offset
		filter += "and offset:=offset"
	}
	query := "select restaurant_id,name,description,price from menus where deleted_at is null"

	query = query + filter + limit + offset
	query, arr = help.ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	var menus []*pb.MenuResponse

	for rows.Next() {
		var menu pb.MenuResponse
		err := rows.Scan(&menu.RestaurantId, &menu.Name, &menu.Description, &menu.Price)
		if err != nil {
			return nil, err
		}
		menus = append(menus, &menu)
	}
	return &pb.MenusResponse{Menus: menus}, nil
}
