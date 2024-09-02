package usecase

import "rest-api/model"

type TripUsecase struct {
}

func NewTripUsecase() TripUsecase {
	return TripUsecase{}
}

func (tu *TripUsecase) GetTrips() ([]model.Trip, error) {
	return []model.Trip{}, nil
}
