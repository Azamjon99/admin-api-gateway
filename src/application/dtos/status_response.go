package dtos

type StatusResponse struct {
	Status string `json:"status"`
}

func NewSuccessStatusResponse() *StatusResponse {
	return &StatusResponse{
		Status: "Success",
	}
}
