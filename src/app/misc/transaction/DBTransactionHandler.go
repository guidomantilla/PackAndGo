package transaction

import (
	"PackAndGo/src/app/misc/datasource"
	"database/sql"
	"log"
)

type DBTransactionContext struct{}

type DBTransactionHandlerFunction func(tx *sql.Tx) error

type DBTransactionHandler interface {
	HandleTransaction(fn DBTransactionHandlerFunction) error
}

type DefaultDBTransactionHandler struct {
	datasource.DBDataSource
}

func NewDefaultDBTransactionHandler(dbDatasource datasource.DBDataSource) *DefaultDBTransactionHandler {
	return &DefaultDBTransactionHandler{
		DBDataSource: dbDatasource,
	}
}

func (handler *DefaultDBTransactionHandler) HandleTransaction(fn DBTransactionHandlerFunction) error {
	tx, err := handler.GetDatabase().Begin()
	if err != nil {
		handleError(err)
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			handleError(tx.Rollback())
			log.Println(err)
		} else if err != nil {
			// something went wrong, rollback
			handleError(tx.Rollback())
		} else {
			// all good, commit
			err = tx.Commit()
			handleError(err)
		}
	}()

	err = fn(tx)
	return err
}
func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
