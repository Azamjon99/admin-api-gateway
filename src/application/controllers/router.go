package controllers

import (
	logisticspb "github.com/Azamjon99/admin-api-gateway/src/application/protos/logistics_vehicle"
	"github.com/Azamjon99/admin-api-gateway/src/application/services"
	"github.com/Azamjon99/admin-api-gateway/src/infrastructure/jwt"

	"github.com/gin-gonic/gin"
)

type routes struct {
	root *gin.Engine
	v1   *gin.RouterGroup
}
type middlewares struct {
	requireToken gin.HandlerFunc
}
type router struct {
	routes                 *routes
	middlewares            *middlewares
	tokenSvc               jwt.Service
	driverApp              services.DriverApplicationService
	logisticsVehicleClient logisticspb.VehicleServiceClient
}

func Init(
	root *gin.Engine,
	tokenSvc jwt.Service,
	driverApp services.DriverApplicationService,
	logisticsVehicleClient logisticspb.VehicleServiceClient,
) {
	v1 := root.Group("/v1")
	routes := routes{
		root: root,
		v1:   v1,
	}
	middlewares := middlewares{
		requireToken: tokenSvc.Middleware().RequireToken(),
	}

	router := router{
		routes:        &routes,
		middlewares:   &middlewares,
		tokenSvc:      tokenSvc,
		driverApp:     driverApp,
		logisticsVehicleClient: logisticsVehicleClient,
	}
	router.initDriverControler()
	router.initVehicleControler()
}
