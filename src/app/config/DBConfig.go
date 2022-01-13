package config

import (
	"PackAndGo/src/app/misc/datasource"
	"PackAndGo/src/app/misc/environment"

	"go.uber.org/zap"
)

var singletonDataSource datasource.DBDataSource

func StopDB() {

	if err := singletonDataSource.GetDatabase().Close(); err != nil {
		zap.L().Fatal("Error closing DB")
		return
	}
}

func InitDB(environment environment.Environment) datasource.DBDataSource {
	username := environment.GetValue(DATASOURCE_USERNAME).AsString()
	password := environment.GetValue(DATASOURCE_PASSWORD).AsString()
	url := environment.GetValue(DATASOURCE_URL).AsString()
	singletonDataSource = datasource.NewMysqlDataSource(username, password, url)
	return singletonDataSource
}
