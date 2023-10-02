package controllers

import (
	"net/http"

	"github.com/Azamjon99/admin-api-gateway/src/application/dtos"
	"github.com/Azamjon99/admin-api-gateway/src/application/services"
	"github.com/gin-gonic/gin"
)

type driverController struct {
	driverApp services.DriverApplicationService
}

func (r *router) initDriverControler() {
	h := &driverController{
		driverApp: r.driverApp,
	}
	r.routes.v1.POST("/drivers", r.middlewares.requireToken, h.handleRegisterDriver)
	// r.routes.v1.DELETE("/drivers/:driver_id", r.middlewares.requireToken, h.handleDeleteDriver)
	r.routes.v1.PUT("/drivers", r.middlewares.requireToken, h.handleUpdateDriverProfile)
	r.routes.v1.GET("/drivers/:driver_id", r.middlewares.requireToken, h.handleGetDriverProfile)
	// r.routes.v1.GET("/drivers", r.middlewares.requireToken, h.handleListDriverProfiles)
}

func (h *driverController) handleRegisterDriver(c *gin.Context) {
	ctx := c.Request.Context()

	var req dtos.RegisterDriverRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.driverApp.RegisterDriver(ctx, &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
// func (h *driverController) handleDeleteDriver(c *gin.Context) {

// 	ctx := c.Request.Context()

// 	driverID := c.Param("driver_id")

// 	response, err := h.driverApp.DeleteDriver(ctx, driverID)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, response)
// }
func (h *driverController) handleUpdateDriverProfile(c *gin.Context) {

	ctx := c.Request.Context()
	var req dtos.UpdateDriverProfileRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.driverApp.UpdateDriverProfile(ctx, &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
func (h *driverController) handleGetDriverProfile(c *gin.Context) {

	ctx := c.Request.Context()

	driverID := c.Param("driver_id")

	response, err := h.driverApp.GetDriverProfile(ctx, driverID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
// func (h *driverController) handleListDriverProfiles(c *gin.Context) {

// 	ctx := c.Request.Context()

// 	response, err := h.driverApp.ListDriverProfiles(ctx)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, response)
// }
