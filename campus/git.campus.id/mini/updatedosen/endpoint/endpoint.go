package endpoint

import (
	"context"

	svc "MCCampus/campus/git.campus.id/mini/updatedosen/server"
	kit "github.com/go-kit/kit/endpoint"
)

type DosenEndpoint struct {
	AddDosenEndpoint           kit.Endpoint
	ReadDosenByKdDosenEndpoint kit.Endpoint
	ReadDosenEndpoint          kit.Endpoint
	UpdateDosenEndpoint        kit.Endpoint
	//ReadCustomerByEmailEndpoint  kit.Endpoint
}

func NewDosenEndpoint(service svc.DosenService) DosenEndpoint {
	addDosenEp := makeAddDosenEndpoint(service)
	readDosenByKdDosenEp := makeReadDosenByKdDosenEndpoint(service)
	readDosenEp := makeReadDosenEndpoint(service)
	updateDosenEp := makeUpdateDosenEndpoint(service)
	//readCustomerByEmailEp := makeReadCustomerByEmailEndpoint(service)
	return DosenEndpoint{AddDosenEndpoint: addDosenEp,
		ReadDosenByKdDosenEndpoint: readDosenByKdDosenEp,
		ReadDosenEndpoint:          readDosenEp,
		UpdateDosenEndpoint:        updateDosenEp,
	}
	//ReadCustomerByEmailEndpoint:  readCustomerByEmailEp,

}

func makeAddDosenEndpoint(service svc.DosenService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Dosen)
		err := service.AddDosenService(ctx, req)
		return nil, err
	}
}

func makeReadDosenByKdDosenEndpoint(service svc.DosenService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Dosen)
		result, err := service.ReadDosenByKdDosenService(ctx, req.KdDosen)
		/*return svc.Customer{CustomerId: result.CustomerId, Name: result.Name,
		CustomerType: result.CustomerType, Mobile: result.Mobile, Email: result.Email,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
		return result, err
	}
}

func makeReadDosenEndpoint(service svc.DosenService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadDosenService(ctx)
		return result, err
	}
}

func makeUpdateDosenEndpoint(service svc.DosenService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Dosen)
		err := service.UpdateDosenService(ctx, req)
		return nil, err
	}
}

/*
func makeReadCustomerByEmailEndpoint(service svc.CustomerService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Customer)
		result, err := service.ReadCustomerByEmailService(ctx, req.Email)
		return result, err
	}
}
*/
