package dtos

type AddVehicleRequest struct {
	DriverId string `json:"driver_id"`
	Model    string `json:"model"`
	Make     string `json:"make"`
}
