package usecase

import (
	"rest-api/model"
	"rest-api/repository"
)

type TripUsecase struct {
	repository repository.TripRepository
}

func NewTripUsecase(repo repository.TripRepository) TripUsecase {
	return TripUsecase{
		repository: repo,
	}
}

func (tu *TripUsecase) GetTrips() ([]model.Trip, error) {
	return tu.repository.GetTrips()
}

func (tu *TripUsecase) CreateTrip(trip model.Trip) (model.Trip, error) {
	tripId, err := tu.repository.CreateTrip(trip)
	if err != nil {
		return model.Trip{}, err
	}

	trip.TripId = tripId

	return trip, nil
}

func (tu *TripUsecase) GetById(id int) (*model.Trip, error) {
	trip, err := tu.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return trip, nil
}
