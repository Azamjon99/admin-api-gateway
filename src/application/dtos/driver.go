package dtos

type RegisterDriverRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type UpdateDriverProfileRequest struct {
	DriverId string `json:"driver_id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}
