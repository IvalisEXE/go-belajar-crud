package endpoint

import (
	"context"

	sv "MCCampus/campus/git.campus.id/mini/adddosen/server"
)

func (me DosenEndpoint) AddDosenService(ctx context.Context, dosen sv.Dosen) error {
	_, err := me.AddDosenEndpoint(ctx, dosen)
	return err
}

/*
func (ce CustomerEndpoint) ReadCustomerByMobileService(ctx context.Context, mobile string) (sv.Customer, error) {
	req := sv.Customer{Mobile: mobile}
	fmt.Println(req)
	resp, err := ce.ReadCustomerByMobileEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Customer)
	return cus, err
}

func (ce CustomerEndpoint) ReadCustomerService(ctx context.Context) (sv.Customers, error) {
	resp, err := ce.ReadCustomerEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Customers), err
}

func (ce CustomerEndpoint) UpdateCustomerService(ctx context.Context, cus sv.Customer) error {
	_, err := ce.UpdateCustomerEndpoint(ctx, cus)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

func (ce CustomerEndpoint) ReadCustomerByEmailService(ctx context.Context, email string) (sv.Customer, error) {
	req := sv.Customer{Email: email}
	resp, err := ce.ReadCustomerByEmailEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Customer)
	return cus, err
}
*/
