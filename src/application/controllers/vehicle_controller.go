package controllers

import (
	"net/http"

	"github.com/Azamjon99/admin-api-gateway/src/application/dtos"
	pb "github.com/Azamjon99/admin-api-gateway/src/application/protos/logistics_vehicle"
	"github.com/gin-gonic/gin"
)

type vehicleController struct {
	vehicleClient pb.VehicleServiceClient
}

func (r *router) initVehicleControler() {
	h := &vehicleController{
		vehicleClient: r.logisticsVehicleClient,
	}
	r.routes.v1.POST("/drivers/:driver_id/vehicles", r.middlewares.requireToken, h.handleAddVehicle)
	r.routes.v1.PUT("/drivers/vehicles", r.middlewares.requireToken, h.handleUpdateVehicle)
	r.routes.v1.DELETE("/drivers/vehicles/:vehicle_id", r.middlewares.requireToken, h.handleDeleteVehicle)
	r.routes.v1.GET("/drivers/vehicles/:vehicle_id", r.middlewares.requireToken, h.handleGetVehicle)
	r.routes.v1.GET("/drivers/:driver_id/vehicles", r.middlewares.requireToken, h.handleListVehiclesByDriver)
}

func (h *vehicleController) handleAddVehicle(c *gin.Context) {

	ctx := c.Request.Context()

	var req dtos.AddVehicleRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	req.DriverId = c.Param("driver_id")
	response, err := h.vehicleClient.CreateVehicle(ctx, &pb.CreateVehicleRequest{
		DriverId: req.DriverId,
		Model:    req.Model,
		Make:     req.Make,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
func (h *vehicleController) handleUpdateVehicle(c *gin.Context) {
	ctx := c.Request.Context()

	var req pb.Vehicle
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	response, err := h.vehicleClient.UpdateVehicle(ctx, &pb.UpdateVehicleRequest{
		Vehicle: &req,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
func (h *vehicleController) handleDeleteVehicle(c *gin.Context) {
	ctx := c.Request.Context()

	vehicleID := c.Param("vehicle_id")

	_, err := h.vehicleClient.DeleteVehicle(ctx, &pb.DeleteVehicleRequest{
		VehicleId: vehicleID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dtos.NewSuccessStatusResponse())
}
func (h *vehicleController) handleGetVehicle(c *gin.Context) {
	ctx := c.Request.Context()

	vehicleID := c.Param("vehicle_id")

	response, err := h.vehicleClient.GetVehicle(ctx, &pb.GetVehicleRequest{
		VehicleId: vehicleID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
func (h *vehicleController) handleListVehiclesByDriver(c *gin.Context) {
	ctx := c.Request.Context()

	driverID := c.Param("driver_id")

	response, err := h.vehicleClient.ListVehicle(ctx, &pb.ListVehicleRequest{
		DriverId: driverID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
