package repository

import (
	"PackAndGo/src/app/core/model"
	"PackAndGo/src/app/misc/transaction"
	"context"
	"database/sql"
	"fmt"
	"log"
)

type CityRepository interface {
	Create(ctx context.Context, trip *model.City) error
	FindById(ctx context.Context, id int64) (*model.City, error)
	FindByName(ctx context.Context, name string) (*model.City, error)
	FindAll(ctx context.Context) (*[]model.City, error)
}

type DefaultCityRepository struct {
	statementCreate     string
	statementFindById   string
	statementFindByName string
	statementFind       string
}

func NewDefaultCityRepository() *DefaultCityRepository {
	return &DefaultCityRepository{
		statementCreate:     "insert into city (name) values (?)",
		statementFindById:   "select id, name from city where id = ?",
		statementFindByName: "select id, name from city where name = ?",
		statementFind:       "select id, name from city",
	}
}

func (repository *DefaultCityRepository) Create(ctx context.Context, city *model.City) error {

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
	if result, err = statement.Exec(city.Name); err != nil {
		return err
	}

	if city.Id, err = result.LastInsertId(); err != nil {
		return err
	}

	return nil
}

func (repository *DefaultCityRepository) FindById(ctx context.Context, id int64) (*model.City, error) {

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

	var city model.City
	if err = row.Scan(&city.Id, &city.Name); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("city with id %d not found", id)
		}
		return nil, err
	}

	return &city, nil
}

func (repository *DefaultCityRepository) FindByName(ctx context.Context, name string) (*model.City, error) {

	var tx = ctx.Value(transaction.DBTransactionContext{}).(*sql.Tx)

	var err error
	var statement *sql.Stmt

	if statement, err = tx.Prepare(repository.statementFindByName); err != nil {
		return nil, err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Println("Error closing the statement")
		}
	}(statement)

	row := statement.QueryRow(name)

	var city model.City
	if err = row.Scan(&city.Id, &city.Name); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("city with name %s not found", name)
		}
		return nil, err
	}

	return &city, nil
}

func (repository *DefaultCityRepository) FindAll(ctx context.Context) (*[]model.City, error) {

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

	cities := make([]model.City, 0)
	for rows.Next() {

		var city model.City
		if err = rows.Scan(&city.Id, &city.Name); err != nil {
			return nil, err
		}

		cities = append(cities, city)
	}

	return &cities, nil
}
