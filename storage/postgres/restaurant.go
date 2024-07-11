package postgres

import (
	"database/sql"
	"fmt"
	pb "reservation_service/genproto"
	storage "reservation_service/help"
	"time"
)

type RestaurantRepository struct {
	Db *sql.DB
}

func NewRestaurantRepo(db *sql.DB) *RestaurantRepository {
	return &RestaurantRepository{Db: db}
}

func (repo *RestaurantRepository) CreateRestaurant(request *pb.CreateRestaurantRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("insert into restaurants(name,address,phone_number,description)values ($1,$2,$3,$4)", request.Name, request.Address, request.PhoneNumber, request.Description)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (repo *RestaurantRepository) UpdateRestaurant(request *pb.UpdateRestaurantRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update restaurants set name=$1,address=$2,phone_number=$3,description=$4 ,updated_at = $5  where id=$6 and deleted_at is null", request.Name, request.Address, request.PhoneNumber, request.Description, time.Now(), request.Id)
	if err != nil {
		fmt.Println("+++++++++++++++")
		return nil, err
	}
	return &pb.Void{}, nil
}

func (repo *RestaurantRepository) DeleteRestaurant(request *pb.IdRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update  restaurants set  deleted_at=current_timestamp where  deleted_at is null and id=$1", request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (repo *RestaurantRepository) GetByIdRestaurant(request *pb.IdRequest) (*pb.RestaurantResponse, error) {
	var restaurant pb.RestaurantResponse
	err := repo.Db.QueryRow("select name,address,phone_number,description  from restaurants where id=$1 and deleted_at is null", request.Id).Scan(&restaurant.Name, &restaurant.Address, &restaurant.PhoneNumber, &restaurant.Description)
	if err != nil {
		return nil, err
	}
	return &restaurant, err

}
func (repo *RestaurantRepository) GetAllRestaurants(request *pb.GetAllRestaurantRequest) (*pb.RestaurantsResponse, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	filter := ""

	// Constructing the WHERE clause based on request parameters
	if len(request.Address) > 0 {
		params["address"] = request.Address
		filter += " AND address = :address"
	}

	// Handling LIMIT and OFFSET for pagination
	if request.LimitOffset.Limit > 0 {
		params["limit"] = request.LimitOffset.Limit
		limit = " LIMIT :limit"
	}

	if request.LimitOffset.Offset > 0 {
		params["offset"] = request.LimitOffset.Offset
		offset = " OFFSET :offset"
	}

	// Construct the SQL query
	query := "SELECT name, address, phone_number, description FROM restaurants WHERE deleted_at IS NULL"
	query += filter + limit + offset

	query, arr = storage.ReplaceQueryParams(query, params)

	// Execute the query
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over rows and populate restaurants slice
	var restaurants []*pb.RestaurantResponse
	for rows.Next() {
		var restaurant pb.RestaurantResponse
		err := rows.Scan(&restaurant.Name, &restaurant.Address, &restaurant.PhoneNumber, &restaurant.Description)
		if err != nil {
			return nil, err
		}
		restaurants = append(restaurants, &restaurant)
	}

	return &pb.RestaurantsResponse{
		Restaurants: restaurants,
	}, nil
}
