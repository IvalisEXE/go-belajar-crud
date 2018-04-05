package endpoint

import (
	"context"
	"fmt"

	sv "MCCampus/campus/git.campus.id/golGajiDosen/showgaji/server"
)

func (me GajiDosenEndpoint) AddGajiDosenService(ctx context.Context, gajidosen sv.GajiDosen) error {
	_, err := me.AddGajiDosenEndpoint(ctx, gajidosen)
	return err
}

func (me GajiDosenEndpoint) ReadGajiDosenByKdGGDosenService(ctx context.Context, kdgajidosen string) (sv.GajiDosen, error) {
	req := sv.GajiDosen{KdGGDosen: kdgajidosen}
	fmt.Println(req)
	resp, err := me.ReadGajiDosenByKdGGDosenEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.GajiDosen)
	return cus, err
}

func (me GajiDosenEndpoint) ReadGajiDosenService(ctx context.Context) (sv.GajiDosens, error) {
	resp, err := me.ReadGajiDosenEndpoint(ctx, nil)
	fmt.Println("me resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.GajiDosens), err
}

/*
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
