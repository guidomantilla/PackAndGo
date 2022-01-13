package service

import (
	"PackAndGo/src/app/core/exception"
	"PackAndGo/src/app/core/model"
	"PackAndGo/src/app/core/repository"
	"PackAndGo/src/app/core/service/dto"
	"PackAndGo/src/app/misc/transaction"
	"context"
	"database/sql"
	"errors"
)

type TripService interface {
	Create(ctx context.Context, tripDto *dto.TripDTO) *exception.Exception
	FindById(ctx context.Context, id int64) (*dto.TripDTO, *exception.Exception)
	FindAll(ctx context.Context) (*[]dto.TripDTO, *exception.Exception)
}

type DefaultTripService struct {
	transaction.DBTransactionHandler
	tripRepository repository.TripRepository
	cityRepository repository.CityRepository
}

func NewDefaultTripService(dbTransactionHandler transaction.DBTransactionHandler, tripRepository repository.TripRepository, cityRepository repository.CityRepository) *DefaultTripService {
	return &DefaultTripService{
		DBTransactionHandler: dbTransactionHandler,
		tripRepository:       tripRepository,
		cityRepository:       cityRepository,
	}
}

func (service *DefaultTripService) FindById(ctx context.Context, id int64) (*dto.TripDTO, *exception.Exception) {

	var err error
	var tripDto *dto.TripDTO
	err = service.HandleTransaction(func(tx *sql.Tx) error {

		txCtx := context.WithValue(ctx, transaction.DBTransactionContext{}, tx)

		var trip *model.Trip
		if trip, err = service.tripRepository.FindById(txCtx, id); err != nil {
			return err
		}

		var origin *model.City
		if origin, err = service.cityRepository.FindById(txCtx, trip.OriginId); err != nil {
			return err
		}

		var destination *model.City
		if destination, err = service.cityRepository.FindById(txCtx, trip.DestinationId); err != nil {
			return err
		}

		tripDto = &dto.TripDTO{
			Origin:      origin.Name,
			Destination: destination.Name,
			Dates:       trip.Dates,
			Price:       trip.Price,
		}

		return nil
	})

	if err != nil {
		return nil, exception.InternalServerErrorException("error finding the trip", err)
	}

	return tripDto, nil
}

func (service *DefaultTripService) FindAll(ctx context.Context) (*[]dto.TripDTO, *exception.Exception) {

	var err error
	tripsDto := make([]dto.TripDTO, 0)
	err = service.HandleTransaction(func(tx *sql.Tx) error {

		txCtx := context.WithValue(ctx, transaction.DBTransactionContext{}, tx)

		var trips *[]model.Trip
		if trips, err = service.tripRepository.FindAll(txCtx); err != nil {
			return err
		}

		for _, trip := range *trips {

			var origin *model.City
			if origin, err = service.cityRepository.FindById(txCtx, trip.OriginId); err != nil {
				return err
			}

			var destination *model.City
			if destination, err = service.cityRepository.FindById(txCtx, trip.DestinationId); err != nil {
				return err
			}

			tripDto := dto.TripDTO{
				Origin:      origin.Name,
				Destination: destination.Name,
				Dates:       trip.Dates,
				Price:       trip.Price,
			}

			tripsDto = append(tripsDto, tripDto)
		}

		return nil
	})

	if err != nil {
		return nil, exception.InternalServerErrorException("error finding the trips", err)
	}

	return &tripsDto, nil
}

func (service *DefaultTripService) Create(ctx context.Context, tripDto *dto.TripDTO) *exception.Exception {

	var err error
	if err = validateTrip(tripDto); err != nil {
		return exception.BadRequestException("error creating the trip", err)
	}

	err = service.HandleTransaction(func(tx *sql.Tx) error {

		txCtx := context.WithValue(ctx, transaction.DBTransactionContext{}, tx)

		if err = validateCity(service, txCtx, tripDto.OriginId); err != nil {
			return err
		}

		if err = validateCity(service, txCtx, tripDto.DestinationId); err != nil {
			return err
		}

		trip := &model.Trip{
			OriginId:      tripDto.OriginId,
			DestinationId: tripDto.DestinationId,
			Dates:         tripDto.Dates,
			Price:         tripDto.Price,
		}

		if err = service.tripRepository.Create(txCtx, trip); err != nil {
			return err
		}

		tripDto.Id = trip.Id

		return nil
	})
	if err != nil {
		return exception.InternalServerErrorException("error creating the trip", err)
	}

	return nil
}

func validateTrip(tripDto *dto.TripDTO) error {

	if tripDto.Id != 0 {
		return errors.New("trip id must be undefined")
	}

	if tripDto.OriginId <= 0 || tripDto.DestinationId <= 0 {
		return errors.New("origin id and destination id must be greater or equals to one")
	}

	if tripDto.Dates == "" {
		return errors.New("dates must be defined")
	}

	if tripDto.Price <= 0 {
		return errors.New("price must be greater or equals to one")
	}

	return nil
}

func validateCity(service *DefaultTripService, ctx context.Context, cityId int64) error {

	var err error
	if _, err = service.cityRepository.FindById(ctx, cityId); err != nil {
		return err
	}
	return nil
}
