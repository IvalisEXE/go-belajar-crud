package endpoint

import (
	"context"

	svc "MCCampus/campus/git.campus.id/mini/adddosen/server"
	kit "github.com/go-kit/kit/endpoint"
)

type DosenEndpoint struct {
	AddDosenEndpoint kit.Endpoint
	//ReadCustomerByMobileEndpoint kit.Endpoint
	//ReadCustomerEndpoint         kit.Endpoint
	//UpdateCustomerEndpoint       kit.Endpoint
	//ReadCustomerByEmailEndpoint  kit.Endpoint
}

func NewDosenEndpoint(service svc.DosenService) DosenEndpoint {
	addDosenEp := makeAddDosenEndpoint(service)
	//readCustomerByMobileEp := makeReadCustomerByMobileEndpoint(service)
	//readCustomerEp := makeReadCustomerEndpoint(service)
	//updateCustomerEp := makeUpdateCustomerEndpoint(service)
	//readCustomerByEmailEp := makeReadCustomerByEmailEndpoint(service)
	return DosenEndpoint{AddDosenEndpoint: addDosenEp} //ReadCustomerByMobileEndpoint: readCustomerByMobileEp,
	//ReadCustomerEndpoint:         readCustomerEp,
	//UpdateCustomerEndpoint:       updateCustomerEp,
	//ReadCustomerByEmailEndpoint:  readCustomerByEmailEp,

}

func makeAddDosenEndpoint(service svc.DosenService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Dosen)
		err := service.AddDosenService(ctx, req)
		return nil, err
	}
}

/*
func makeReadCustomerByMobileEndpoint(service svc.CustomerService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Customer)
		result, err := service.ReadCustomerByMobileService(ctx, req.Mobile)
		/*return svc.Customer{CustomerId: result.CustomerId, Name: result.Name,
		CustomerType: result.CustomerType, Mobile: result.Mobile, Email: result.Email,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
/*return result, err
	}
}*/
/*
func makeReadCustomerEndpoint(service svc.CustomerService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadCustomerService(ctx)
		return result, err
	}
}

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
