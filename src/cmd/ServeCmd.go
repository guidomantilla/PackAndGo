package cmd

import (
	"PackAndGo/src/app/config"
	"PackAndGo/src/app/core/repository"
	"PackAndGo/src/app/core/service"
	"PackAndGo/src/app/core/ws"
	"PackAndGo/src/app/mgmt"
	"PackAndGo/src/app/misc/transaction"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func CreateServeCmd() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {

			environment := config.InitParams(&args)
			defer config.StopParams()

			config.InitMonitoring(environment)
			defer config.StopMonitoring()

			dataSource := config.InitDB(environment)
			defer config.StopDB()

			infoWs := mgmt.NewDefaultInfoWs()
			envWs := mgmt.NewDefaultEnvWs(environment)
			metricsWs := mgmt.NewDefaultMetricsWs()
			healthWs := mgmt.NewDefaultHealthWs()

			transactionHandler := transaction.NewDefaultDBTransactionHandler(dataSource)

			tripRepository := repository.NewDefaultTripRepository()
			cityRepository := repository.NewDefaultCityRepository()

			tripService := service.NewDefaultTripService(transactionHandler, tripRepository, cityRepository)

			tripWs := ws.NewDefaultTripWs(tripService)

			if err := config.InitWebServer(environment, tripWs, infoWs, envWs, metricsWs, healthWs); err != nil {
				zap.L().Fatal("error starting the server.")
			}
			defer config.StopWebServer()
		},
	}
}
