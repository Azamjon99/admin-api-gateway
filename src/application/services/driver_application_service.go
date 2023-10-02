package services

import (
	"context"

	"github.com/Azamjon99/admin-api-gateway/src/application/dtos"
	pb "github.com/Azamjon99/admin-api-gateway/src/application/protos/logistics_staff"
)

type DriverApplicationService interface {
	RegisterDriver(ctx context.Context, req *dtos.RegisterDriverRequest) (*pb.RegisterDriverResponse, error)
	GetDriverProfile(ctx context.Context, driverID string) (*pb.GetDriverProfileResponse, error)
	UpdateDriverProfile(ctx context.Context, req *dtos.UpdateDriverProfileRequest) (*pb.UpdateDriverProfileResponse, error)
	// DeleteDriver(ctx context.Context, driverID string) (*dtos.StatusResponse, error)
	// ListDriverProfiles(ctx context.Context) (*pb.ListDriverProfilesResponse, error)
}

type driverAppSvcImpl struct {
	driverClient pb.StaffServiceClient
}

func NewDriverApplicationService(driverClient pb.StaffServiceClient) DriverApplicationService {
	return &driverAppSvcImpl{
		driverClient: driverClient,
	}
}
func (s *driverAppSvcImpl) RegisterDriver(ctx context.Context, req *dtos.RegisterDriverRequest) (*pb.RegisterDriverResponse, error) {

	result, err := s.driverClient.RegisterDriver(ctx, &pb.RegisterDriverRequest{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *driverAppSvcImpl) GetDriverProfile(ctx context.Context, driverID string) (*pb.GetDriverProfileResponse, error) {

	result, err := s.driverClient.GetDriverProfile(ctx, &pb.GetDriverProfileRequest{
		DriverId: driverID,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *driverAppSvcImpl) UpdateDriverProfile(ctx context.Context, req *dtos.UpdateDriverProfileRequest) (*pb.UpdateDriverProfileResponse, error) {

	result, err := s.driverClient.UpdateDriverProfile(ctx, &pb.UpdateDriverProfileRequest{
		DriverId: req.DriverId,
		Name:     req.Name,
		ImageUrl: req.ImageUrl,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
// func (s *driverAppSvcImpl) DeleteDriver(ctx context.Context, driverID string) (*dtos.StatusResponse, error) {

// 	_, err := s.driverClient.DeleteDriver(ctx, &pb.DeleteDriverRequest{
// 		DriverId: driverID,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return dtos.NewSuccessStatusResponse(), nil
// }
// func (s *driverAppSvcImpl) ListDriverProfiles(ctx context.Context) (*pb.ListDriverProfilesResponse, error) {

// 	result, err := s.driverClient.ListDriverProfiles(ctx, &pb.ListDriverProfilesRequest{})

// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }
