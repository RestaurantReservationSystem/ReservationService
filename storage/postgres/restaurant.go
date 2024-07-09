package postgres

import (
	"database/sql"
	pb "reservation_service/genproto"
	storage "reservation_service/help"
	"time"
)

type RestaurantRepo struct {
	Db *sql.DB
}

func NewRestaurantRepo(db *sql.DB) *RestaurantRepo {
	return &RestaurantRepo{Db: db}
}

func (repo *RestaurantRepo) CreateRestaurant(request *pb.CreateRestaurantRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("insert into restaurants(name,address,phone_number,description)values ($1,$2,$3,$4)", request.Name, request.Address, request.PhoneNumber, request.Description)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (repo *RestaurantRepo) UpdateRestaurant(request *pb.UpdateRestaurantRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update restaurants set name=$1,address=$2,phone_number=$3,description=$4 ,updated_at = $5  where id=$6 and deleted_at is null", request.Name, request.Address, request.PhoneNumber, request.Description, time.Now(), request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (repo *RestaurantRepo) DeleteRestaurant(request *pb.IdRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update  restaurants set  deleted_at=current_timestamp where  deleted_at is null and id=$1", request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (repo *RestaurantRepo) GetByIdRestaurant(request *pb.IdRequest) (*pb.RestaurantResponse, error) {
	var restaurant pb.RestaurantResponse
	err := repo.Db.QueryRow("select name,address,phone_number,description  from restaurants where id=$1 and deleted_at is null", request.Id).Scan(&restaurant.Name, &restaurant.Address, &restaurant.PhoneNumber, &restaurant.Description)
	if err != nil {
		return nil, err
	}
	return &restaurant, err

}
func (repo *RestaurantRepo) GetAllRestaurants(request *pb.GetAllRestaurantRequest) (*pb.RestaurantsResponse, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	filter := ""

	if len(request.Address) > 0 {
		params["address"] = request.Address
		arr = append(arr, " ADDRESS :address ")
	}

	if request.LimitOffset.Limit > 0 {
		params["limit"] = request.LimitOffset.Limit
		limit += "LIMIT :limit "
	}

	if request.LimitOffset.Offset > 0 {
		params["offset"] = request.LimitOffset.Offset
		offset += "OFFSET :offset "
	}

	query := "select name,address ,phone_number,description from restaurants  where  deleted_at is null "

	query = query + filter + limit + offset
	query, arr = storage.ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
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
