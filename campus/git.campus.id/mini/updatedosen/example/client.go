package main

import (
	"context"
	"time"

	cli "MCCampus/campus/git.campus.id/mini/updatedosen/endpoint"
	svc "MCCampus/campus/git.campus.id/mini/updatedosen/server"
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
	//client.AddDosenService(context.Background(), svc.Dosen{KdDosen: "D010", NamaDosen: "Aripss", CreateBy: "aripssss"})

	//Get Mahasiswa By Nim No
	//cusKdDosen, _ := client.ReadDosenByKdDosenService(context.Background(), "D010")
	//fmt.Println("mahasiswa based on kddosen:", cusKdDosen)

	//List Customer
	//cuss, _ := client.ReadDosenService(context.Background())
	//fmt.Println("all dosens:", cuss)

	//Update Dosen
	client.UpdateDosenService(context.Background(), svc.Dosen{KdDosen: "D010", NamaDosen: "Aripss", Status: 0, UpdateBy: "aripssss", UpdateOn: "Mr.GAGA"})

	/*
		//Get Customer By Email
		cusEmail, _ := client.ReadCustomerByEmailService(context.Background(), "joko@gmail.com")
		fmt.Println("customer based on email:", cusEmail)
	*/
}
