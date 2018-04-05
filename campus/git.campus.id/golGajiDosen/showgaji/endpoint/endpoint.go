package endpoint

import (
	"context"

	svc "MCCampus/campus/git.campus.id/golGajiDosen/showgaji/server"
	kit "github.com/go-kit/kit/endpoint"
)

type GajiDosenEndpoint struct {
	AddGajiDosenEndpoint             kit.Endpoint
	ReadGajiDosenByKdGGDosenEndpoint kit.Endpoint
	ReadGajiDosenEndpoint            kit.Endpoint
	//UpdateCustomerEndpoint       kit.Endpoint
	//ReadCustomerByEmailEndpoint  kit.Endpoint
}

func NewGajiDosenEndpoint(service svc.GajiDosenService) GajiDosenEndpoint {
	addGajiDosenEp := makeAddGajiDosenEndpoint(service)
	readGajiDosenByKdGGDosenEp := makeReadGajiDosenByKdGGDosenEndpoint(service)
	readGajiDosenEp := makeReadGajiDosenEndpoint(service)
	//updateCustomerEp := makeUpdateCustomerEndpoint(service)
	//readCustomerByEmailEp := makeReadCustomerByEmailEndpoint(service)
	return GajiDosenEndpoint{AddGajiDosenEndpoint: addGajiDosenEp,
		ReadGajiDosenByKdGGDosenEndpoint: readGajiDosenByKdGGDosenEp,
		ReadGajiDosenEndpoint:            readGajiDosenEp,
	}
	//UpdateCustomerEndpoint:       updateCustomerEp,
	//ReadCustomerByEmailEndpoint:  readCustomerByEmailEp,

}

func makeAddGajiDosenEndpoint(service svc.GajiDosenService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.GajiDosen)
		err := service.AddGajiDosenService(ctx, req)
		return nil, err
	}
}

func makeReadGajiDosenByKdGGDosenEndpoint(service svc.GajiDosenService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.GajiDosen)
		result, err := service.ReadGajiDosenByKdGGDosenService(ctx, req.KdGGDosen)
		/*return svc.Customer{CustomerId: result.CustomerId, Name: result.Name,
		CustomerType: result.CustomerType, Mobile: result.Mobile, Email: result.Email,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
		return result, err
	}
}

func makeReadGajiDosenEndpoint(service svc.GajiDosenService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadGajiDosenService(ctx)
		return result, err
	}
}

/*
func makeUpdateCustomerEndpoint(service svc.CustomerService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Customer)
		err := service.UpdateCustomerService(ctx, req)
		return nil, err
	}
}

func makeReadCustomerByEmailEndpoint(service svc.CustomerService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Customer)
		result, err := service.ReadCustomerByEmailService(ctx, req.Email)
		return result, err
	}
}
*/
