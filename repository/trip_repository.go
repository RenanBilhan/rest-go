package repository

import (
	"database/sql"
	"fmt"
	"rest-api/model"
)

type TripRepository struct {
	connection *sql.DB
}

func NewTripRepository(connection *sql.DB) TripRepository {
	return TripRepository{connection: connection}
}

func (pr *TripRepository) GetTrips() ([]model.Trip, error) {
	query := "SELECT trip_id, trip_name, description FROM trip.trip"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Trip{}, err
	}
	var tripList []model.Trip
	var tripObj model.Trip

	for rows.Next() {
		err = rows.Scan(
			&tripObj.TripId,
			&tripObj.TripName,
			&tripObj.Description)

		if err != nil {
			fmt.Println(err)
			return []model.Trip{}, err
		}

		tripList = append(tripList, tripObj)
	}
	rows.Close()
	return tripList, nil
}

func (tr *TripRepository) CreateTrip(trip model.Trip) (int, error) {
	var id int
	query, err := tr.connection.Prepare("INSERT INTO trip.trip" +
		"(trip_name, description)" +
		"VALUES($1, $2) RETURNING trip_id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	err = query.QueryRow(trip.TripName, trip.Description).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil
}
func (tr *TripRepository) GetById(id int) (*model.Trip, error) {

	query, err := tr.connection.Prepare("SELECT trip_id, trip_name, description FROM trip.trip WHERE trip_id = $1 ")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer query.Close()

	var trip model.Trip

	err = query.QueryRow(id).Scan(
		&trip.TripId,
		&trip.TripName,
		&trip.Description)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	query.Close()
	return &trip, nil

}
