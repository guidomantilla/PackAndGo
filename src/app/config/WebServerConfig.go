package config

import (
	"PackAndGo/src/app/core/ws"
	"PackAndGo/src/app/mgmt"
	"PackAndGo/src/app/misc/environment"

	"github.com/gin-gonic/gin"
)

var singletonEngine *gin.Engine

func StopWebServer() {
	//Nothing to do here yet
}

func InitWebServer(environment environment.Environment, tripWs ws.TripWs,
	infoWs mgmt.InfoWs, envWs mgmt.EnvWs, metricsWs mgmt.MetricsWs, healthWs mgmt.HealthWs) error {

	singletonEngine = gin.Default()

	loadApiRoutes(tripWs)
	loadMgmtRoutes(infoWs, envWs, metricsWs, healthWs)

	hostAddress := environment.GetValueOrDefault(HOST_POST, HOST_POST_DEFAULT_VALUE).AsString()
	return singletonEngine.Run(hostAddress)
}

func loadApiRoutes(tripWs ws.TripWs) {

	group := singletonEngine.Group("/api")
	group.GET("/trip", tripWs.FindAll)
	group.GET("/trip/:id", tripWs.FindById)
	group.POST("/trip", tripWs.Create)
}

func loadMgmtRoutes(infoWs mgmt.InfoWs, envWs mgmt.EnvWs, metricsWs mgmt.MetricsWs, healthWs mgmt.HealthWs) {

	group := singletonEngine.Group("/mgmt")
	group.GET("/info", infoWs.Get)
	group.GET("/env", envWs.Get)
	group.GET("/metrics", metricsWs.Get)
	group.GET("/health", healthWs.Get)
}
