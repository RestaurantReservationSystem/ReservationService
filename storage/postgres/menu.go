package postgres

import (
	"database/sql"
	"fmt"
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
	_, err := repo.Db.Exec("insert into menu(restaurant_id,name,description,price,created_at) values ($1,$2,$3,$4,$5)", request.RestaurantId, request.Name, request.Description, request.Price, time.Now())
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, err
}

func (repo *MenuRepository) UpdateMenu(request *pb.UpdateMenuRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update menu set restaurant_id=$1,name=$2,description=$3,price=$4,updated_at=$5  where id=$6 and deleted_at is null", request.RestaurantId, request.Name, request.Description, request.Price, time.Now(), request.Id)
	if err != nil {
		fmt.Println("+++++++++++++++", err)
		return nil, err
	}
	return &pb.Void{}, err
}

func (repo *MenuRepository) DeleteMenu(request *pb.IdRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update menu set deleted_at=$1 where id=$2 and deleted_at is null", time.Now(), request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, err
}

func (repo *MenuRepository) GetByIdMenu(request *pb.IdRequest) (*pb.MenuResponse, error) {
	var response pb.MenuResponse
	err := repo.Db.QueryRow("select restaurant_id,name,description,price from menu where deleted_at is null and id=$1", request.Id).Scan(&response.RestaurantId, &response.Name, &response.Description, &response.Price)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (repo *MenuRepository) GetAllMenu(request *pb.GetAllMenuRequest) (*pb.MenusResponse, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		filter string
		limit  string
		offset string
	)
	
	if len(request.RestaurantId) > 0 {
		params["restaurant_id"] = request.RestaurantId
		filter += " AND restaurant_id = :restaurant_id"
	}

	if len(request.Name) > 0 {
		params["name"] = request.Name
		filter += " AND name = :name"
	}
	if len(request.Description) > 0 {
		params["description"] = request.Description
		filter += " AND description = :description"
	}
	if request.Price > 0 {
		params["price"] = request.Price
		filter += " AND price = :price"
	}

	if request.LimitOffset.Limit > 0 {
		params["limit"] = request.LimitOffset.Limit
		limit = " LIMIT :limit"
	}
	if request.LimitOffset.Offset > 0 {
		params["offset"] = request.LimitOffset.Offset
		offset = " OFFSET :offset"
	}
	query := "SELECT restaurant_id, name, description, price FROM menu WHERE deleted_at IS NULL"
	if filter != "" {
		query += " AND " + filter[4:] // Remove leading " AND"
	}
	query += limit + offset

	query, arr = help.ReplaceQueryParams(query, params)
	fmt.Println("Query:", query)

	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menus []*pb.MenuResponse
	for rows.Next() {
		var menu pb.MenuResponse
		if err := rows.Scan(&menu.RestaurantId, &menu.Name, &menu.Description, &menu.Price); err != nil {
			return nil, err
		}
		menus = append(menus, &menu)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println("++++++++", menus)

	return &pb.MenusResponse{Menus: menus}, nil
}

