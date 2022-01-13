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

func TestCityRepository(t *testing.T) {

	var err error
	cityRepository := NewDefaultCityRepository()

	t.Run("Test CityRepository FindAll method", func(t *testing.T) {
		dataSource := mocks.NewMockDBDataSource()
		transactionHandler := transaction.NewDefaultDBTransactionHandler(dataSource)

		rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "MockCity01").AddRow(2, "MockCity02")
		dataSource.Mock.ExpectBegin()
		dataSource.Mock.ExpectPrepare("select id, name from city")
		dataSource.Mock.ExpectQuery("select id, name from city").WillReturnRows(rows)
		dataSource.Mock.ExpectCommit()

		err = transactionHandler.HandleTransaction(func(tx *sql.Tx) error {
			txCtx := context.WithValue(context.Background(), transaction.DBTransactionContext{}, tx)

			var cities *[]model.City
			if cities, err = cityRepository.FindAll(txCtx); err != nil {
				t.Fail()
			}

			assert.Equal(t, 2, len(*cities))
			return nil
		})

		dataSource.Mock.ExpectCommit()

		if err != nil {
			t.Fail()
		}
	})

	t.Run("Test CityRepository FindByName method", func(t *testing.T) {
		dataSource := mocks.NewMockDBDataSource()
		transactionHandler := transaction.NewDefaultDBTransactionHandler(dataSource)

		cityId := int64(1)
		cityName := "MockCity01"
		rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(cityId, cityName)
		dataSource.Mock.ExpectBegin()
		dataSource.Mock.ExpectPrepare("select id, name from city where name = ?")
		dataSource.Mock.ExpectQuery("select id, name from city where name = ?").WithArgs(cityName).WillReturnRows(rows)
		dataSource.Mock.ExpectCommit()

		err = transactionHandler.HandleTransaction(func(tx *sql.Tx) error {
			txCtx := context.WithValue(context.Background(), transaction.DBTransactionContext{}, tx)

			var city *model.City
			if city, err = cityRepository.FindByName(txCtx, cityName); err != nil {
				t.Fail()
			}

			assert.Equal(t, cityId, city.Id)
			assert.Equal(t, cityName, city.Name)
			return nil
		})

		dataSource.Mock.ExpectCommit()

		if err != nil {
			t.Fail()
		}
	})

	t.Run("Test CityRepository FindById method", func(t *testing.T) {
		dataSource := mocks.NewMockDBDataSource()
		transactionHandler := transaction.NewDefaultDBTransactionHandler(dataSource)

		cityId := int64(1)
		cityName := "MockCity01"
		rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(cityId, cityName)
		dataSource.Mock.ExpectBegin()
		dataSource.Mock.ExpectPrepare("select id, name from city where id = ?")
		dataSource.Mock.ExpectQuery("select id, name from city where id = ?").WithArgs(cityId).WillReturnRows(rows)
		dataSource.Mock.ExpectCommit()

		err = transactionHandler.HandleTransaction(func(tx *sql.Tx) error {
			txCtx := context.WithValue(context.Background(), transaction.DBTransactionContext{}, tx)

			var city *model.City
			if city, err = cityRepository.FindById(txCtx, cityId); err != nil {
				t.Fail()
			}

			assert.Equal(t, cityId, city.Id)
			assert.Equal(t, cityName, city.Name)
			return nil
		})

		dataSource.Mock.ExpectCommit()

		if err != nil {
			t.Fail()
		}
	})

	t.Run("Test CityRepository Create method", func(t *testing.T) {
		dataSource := mocks.NewMockDBDataSource()
		transactionHandler := transaction.NewDefaultDBTransactionHandler(dataSource)

		cityId := int64(1)
		cityName := "MockCity01"
		result := sqlmock.NewResult(cityId, 1)
		dataSource.Mock.ExpectBegin()
		dataSource.Mock.ExpectPrepare("insert into city")
		dataSource.Mock.ExpectExec("insert into city").WithArgs(cityName).WillReturnResult(result)
		dataSource.Mock.ExpectCommit()

		err = transactionHandler.HandleTransaction(func(tx *sql.Tx) error {
			txCtx := context.WithValue(context.Background(), transaction.DBTransactionContext{}, tx)

			city := &model.City{Name: cityName}
			if err = cityRepository.Create(txCtx, city); err != nil {
				t.Fail()
			}

			assert.Equal(t, cityId, city.Id)
			assert.Equal(t, cityName, city.Name)
			return nil
		})

		dataSource.Mock.ExpectCommit()

		if err != nil {
			t.Fail()
		}
	})
}
