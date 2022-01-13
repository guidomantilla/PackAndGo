package repository

import (
	"PackAndGo/src/app/core/model"
	"PackAndGo/src/app/misc/mocks"
	"PackAndGo/src/app/misc/transaction"
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestTripRepository(t *testing.T) {

	var err error
	tripRepository := NewDefaultTripRepository()

	t.Run("Test TripRepository FindAll method", func(t *testing.T) {
		dataSource := mocks.NewMockDBDataSource()
		transactionHandler := transaction.NewDefaultDBTransactionHandler(dataSource)

		rows := sqlmock.NewRows([]string{"id", "originId", "destinationId", "dates", "price"}).AddRow(1, 1, 1, "Sat Sun", 24.3).AddRow(2, 2, 2, "Sat Sun", 212.3)
		dataSource.Mock.ExpectBegin()
		dataSource.Mock.ExpectPrepare("select id, originId, destinationId, dates, price from trip")
		dataSource.Mock.ExpectQuery("select id, originId, destinationId, dates, price from trip").WillReturnRows(rows)
		dataSource.Mock.ExpectCommit()

		err = transactionHandler.HandleTransaction(func(tx *sql.Tx) error {
			txCtx := context.WithValue(context.Background(), transaction.DBTransactionContext{}, tx)

			var trips *[]model.Trip
			if trips, err = tripRepository.FindAll(txCtx); err != nil {
				t.Fail()
			}

			assert.Equal(t, 2, len(*trips))
			return nil
		})

		if err != nil {
			t.Fail()
		}
	})

	t.Run("Test TripRepository FindById method", func(t *testing.T) {
		dataSource := mocks.NewMockDBDataSource()
		transactionHandler := transaction.NewDefaultDBTransactionHandler(dataSource)

		tripId := int64(1)
		rows := sqlmock.NewRows([]string{"id", "originId", "destinationId", "dates", "price"}).AddRow(tripId, 1, 1, "Sat Sun", 24.3)
		dataSource.Mock.ExpectBegin()
		dataSource.Mock.ExpectPrepare("select id, originId, destinationId, dates, price from trip where id = ?")
		dataSource.Mock.ExpectQuery("select id, originId, destinationId, dates, price from trip where id = ?").WithArgs(tripId).WillReturnRows(rows)
		dataSource.Mock.ExpectCommit()

		err = transactionHandler.HandleTransaction(func(tx *sql.Tx) error {
			txCtx := context.WithValue(context.Background(), transaction.DBTransactionContext{}, tx)

			var trip *model.Trip
			if trip, err = tripRepository.FindById(txCtx, tripId); err != nil {
				t.Fail()
			}

			assert.Equal(t, tripId, trip.Id)
			return nil
		})

		dataSource.Mock.ExpectCommit()

		if err != nil {
			t.Fail()
		}
	})

	t.Run("Test TripRepository Create method", func(t *testing.T) {
		dataSource := mocks.NewMockDBDataSource()
		transactionHandler := transaction.NewDefaultDBTransactionHandler(dataSource)

		tripId := int64(1)
		result := sqlmock.NewResult(tripId, 1)
		dataSource.Mock.ExpectBegin()
		dataSource.Mock.ExpectPrepare("insert into trip")
		dataSource.Mock.ExpectExec("insert into trip").WithArgs(1, 1, "Sat Sun", 24.3).WillReturnResult(result)
		dataSource.Mock.ExpectCommit()

		err = transactionHandler.HandleTransaction(func(tx *sql.Tx) error {
			txCtx := context.WithValue(context.Background(), transaction.DBTransactionContext{}, tx)

			trip := &model.Trip{OriginId: 1, DestinationId: 1, Dates: "Sat Sun", Price: 24.3}
			if err = tripRepository.Create(txCtx, trip); err != nil {
				t.Fail()
			}

			assert.Equal(t, tripId, trip.Id)
			return nil
		})

		if err != nil {
			t.Fail()
		}
	})
}
