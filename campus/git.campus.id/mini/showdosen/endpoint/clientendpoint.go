package endpoint

import (
	"context"
	"fmt"

	sv "MCCampus/campus/git.campus.id/mini/showdosen/server"
)

func (me DosenEndpoint) AddDosenService(ctx context.Context, dosen sv.Dosen) error {
	_, err := me.AddDosenEndpoint(ctx, dosen)
	return err
}

func (me DosenEndpoint) ReadDosenByKdDosenService(ctx context.Context, kddosen string) (sv.Dosen, error) {
	req := sv.Dosen{KdDosen: kddosen}
	fmt.Println(req)
	resp, err := me.ReadDosenByKdDosenEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Dosen)
	return cus, err
}

func (me DosenEndpoint) ReadDosenByKeteranganService(ctx context.Context, ktrg string) (sv.Dosen, error) {
	req := sv.Dosen{Keterangan: ktrg}
	fmt.Println(req)
	resp, err := me.ReadDosenByKeteranganEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Dosen)
	return cus, err
}

func (me DosenEndpoint) ReadDosenService(ctx context.Context) (sv.Dosens, error) {
	resp, err := me.ReadDosenEndpoint(ctx, nil)
	fmt.Println("me resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Dosens), err
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
