package main

import (
	"context"
	"fmt"
	"time"

	cli "MCCampus/campus/git.campus.id/mini/showdosen/endpoint"
	//svc "git.bluebird.id/bluebird/util/server"
	opt "MCCampus/campus/git.campus.id/mini/util/grpc"
	util "MCCampus/campus/git.campus.id/mini/util/microservice"
	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCDosenClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Mahasiswa
	//client.AddDosenService(context.Background(), svc.Dosen{KdDosen: "D010", NamaDosen: "Aripss", CreateBy: "lolo"})

	//Get Mahasiswa By Nim No
	//cusKdDosen, _ := client.ReadDosenByKdDosenService(context.Background(), "D010")
	//fmt.Println("dosen based on kddosen:", cusKdDosen)

	//List Customer
	//
	cuss, _ := client.ReadDosenService(context.Background())
	fmt.Println("all dosens:", cuss)
	/*
		//Update Customer
		client.UpdateCustomerService(context.Background(), svc.Customer{Name: "Joko", CustomerType: 1, Mobile: "0876", Email: "joko@gmail.com", Gender: "M", CallbackPhone: "0876", Status: 1, CustomerId: 2})


		//Get Customer By Email
		cusEmail, _ := client.ReadCustomerByEmailService(context.Background(), "joko@gmail.com")
		fmt.Println("customer based on email:", cusEmail)
	*/
}
