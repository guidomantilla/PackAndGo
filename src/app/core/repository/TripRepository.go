package repository

import (
	"PackAndGo/src/app/core/model"
	"PackAndGo/src/app/misc/transaction"
	"context"
	"database/sql"
	"fmt"
	"log"
)

type TripRepository interface {
	Create(ctx context.Context, trip *model.Trip) error
	FindById(ctx context.Context, id int64) (*model.Trip, error)
	FindAll(ctx context.Context) (*[]model.Trip, error)
}

type DefaultTripRepository struct {
	statementCreate   string
	statementFindById string
	statementFind     string
}

func NewDefaultTripRepository() *DefaultTripRepository {
	return &DefaultTripRepository{
		statementCreate:   "insert into trip (originId, destinationId, dates, price) values (?, ?, ?, ?)",
		statementFindById: "select id, originId, destinationId, dates, price from trip where id = ?",
		statementFind:     "select id, originId, destinationId, dates, price from trip",
	}
}

func (repository *DefaultTripRepository) Create(ctx context.Context, trip *model.Trip) error {

	var tx = ctx.Value(transaction.DBTransactionContext{}).(*sql.Tx)

	var err error
	var statement *sql.Stmt

	if statement, err = tx.Prepare(repository.statementCreate); err != nil {
		return err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Println("Error closing the statement")
		}
	}(statement)

	var result sql.Result
	if result, err = statement.Exec(trip.OriginId, trip.DestinationId, trip.Dates, trip.Price); err != nil {
		return err
	}

	if trip.Id, err = result.LastInsertId(); err != nil {
		return err
	}

	return nil
}

func (repository *DefaultTripRepository) FindById(ctx context.Context, id int64) (*model.Trip, error) {

	var tx = ctx.Value(transaction.DBTransactionContext{}).(*sql.Tx)

	var err error
	var statement *sql.Stmt

	if statement, err = tx.Prepare(repository.statementFindById); err != nil {
		return nil, err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Println("Error closing the statement")
		}
	}(statement)

	row := statement.QueryRow(id)

	var trip model.Trip
	if err = row.Scan(&trip.Id, &trip.OriginId, &trip.DestinationId, &trip.Dates, &trip.Price); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("trip with id %d not found", id)
		}
		return nil, err
	}

	return &trip, nil
}

func (repository *DefaultTripRepository) FindAll(ctx context.Context) (*[]model.Trip, error) {

	var tx = ctx.Value(transaction.DBTransactionContext{}).(*sql.Tx)

	var err error
	var statement *sql.Stmt

	if statement, err = tx.Prepare(repository.statementFind); err != nil {
		return nil, err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Println("Error closing the statement")
		}
	}(statement)

	var rows *sql.Rows
	if rows, err = statement.Query(); err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error closing the result set")
		}
	}(rows)

	trips := make([]model.Trip, 0)
	for rows.Next() {

		var trip model.Trip
		if err = rows.Scan(&trip.Id, &trip.OriginId, &trip.DestinationId, &trip.Dates, &trip.Price); err != nil {
			return nil, err
		}

		trips = append(trips, trip)
	}

	return &trips, nil
}
