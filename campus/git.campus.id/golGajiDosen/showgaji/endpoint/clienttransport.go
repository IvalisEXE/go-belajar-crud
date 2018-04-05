package endpoint

import (
	"context"
	"time"

	svc "MCCampus/campus/git.campus.id/golGajiDosen/showgaji/server"

	pb "MCCampus/campus/git.campus.id/golGajiDosen/showgaji/grpc"

	util "MCCampus/campus/git.campus.id/golGajiDosen/util/grpc"
	disc "MCCampus/campus/git.campus.id/golGajiDosen/util/microservice"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcName = "grpc.GajiDosenService"
)

func NewGRPCGajiDosenClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.GajiDosenService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addGajiDosenEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddGajiDosenEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addGajiDosenEp = retry
	}

	var readGajiDosenByKdGGDosenEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadGajiDosenByKdGGDosenEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readGajiDosenByKdGGDosenEp = retry
	}

	var readGajiDosenEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadGajiDosenEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readGajiDosenEp = retry
	}
	/*
		var updateCustomerEp endpoint.Endpoint
		{
			factory := util.EndpointFactory(makeClientUpdateCustomer, creds, timeout, tracer, logger)
			endpointer := sd.NewEndpointer(instancer, factory, logger)
			balancer := lb.NewRoundRobin(endpointer)
			retry := lb.Retry(retryMax, retryTimeout, balancer)
			updateCustomerEp = retry
		}

		var readCustomerByEmailEp endpoint.Endpoint
		{
			factory := util.EndpointFactory(makeClientReadCustomerByEmail, creds, timeout, tracer, logger)
			endpointer := sd.NewEndpointer(instancer, factory, logger)
			balancer := lb.NewRoundRobin(endpointer)
			retry := lb.Retry(retryMax, retryTimeout, balancer)
			readCustomerByEmailEp = retry
		}
	*/
	return GajiDosenEndpoint{AddGajiDosenEndpoint: addGajiDosenEp, ReadGajiDosenByKdGGDosenEndpoint: readGajiDosenByKdGGDosenEp,
		ReadGajiDosenEndpoint: readGajiDosenEp}, nil
}

func encodeAddGajiDosenRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.GajiDosen)
	return &pb.AddGajiDosenReq{
		KdGGDosen:  req.KdGGDosen,
		JumlahGaji: req.JumlahGaji,
		Keterangan: req.Keterangan,
		Status:     req.Status,
		CreateBy:   req.CreateBy,
		CreateOn:   req.CreateOn,
		UpdateBy:   req.UpdateBy,
		UpdateOn:   req.UpdateOn,
	}, nil
}

func encodeReadGajiDosenByKdGGDosenRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.GajiDosen)
	return &pb.ReadGajiDosenByKdGGDosenReq{KdGGDosen: req.KdGGDosen}, nil
}

func encodeReadGajiDosenRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

/*

func encodeUpdateCustomerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Customer)
	return &pb.UpdateCustomerReq{
		CustomerId:    req.CustomerId,
		Name:          req.Name,
		CustomerType:  req.CustomerType,
		Mobile:        req.Mobile,
		Email:         req.Email,
		Gender:        req.Gender,
		CallbackPhone: req.CallbackPhone,
		Status:        req.Status,
	}, nil
}

func encodeReadCustomerByEmailRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Customer)
	return &pb.ReadCustomerByEmailReq{Email: req.Email}, nil
}
*/

func decodeDosenResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadGajiDosenByKdGGDosenRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadGajiDosenByKdGGDosenResp)
	return svc.GajiDosen{
		KdGGDosen:  resp.KdGGDosen,
		JumlahGaji: resp.JumlahGaji,
		Keterangan: resp.Keterangan,
		Status:     resp.Status,
		CreateBy:   resp.CreateBy,
		CreateOn:   resp.CreateOn,
		UpdateBy:   resp.UpdateBy,
		UpdateOn:   resp.UpdateOn,
	}, nil
}

func decodeReadGajiDosenResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadGajiDosenResp)
	var rsp svc.GajiDosens

	for _, v := range resp.AllGajiDosen {
		itm := svc.GajiDosen{
			KdGGDosen:  v.KdGGDosen,
			JumlahGaji: v.JumlahGaji,
			Keterangan: v.Keterangan,
			Status:     v.Status,
			CreateBy:   v.CreateBy,
			CreateOn:   v.CreateOn,
			UpdateBy:   v.UpdateBy,
			UpdateOn:   v.UpdateOn,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

/*
func decodeReadCustomerByEmailRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadCustomerByEmailResp)
	return svc.Customer{
		CustomerId:    resp.CustomerId,
		Name:          resp.Name,
		CustomerType:  resp.CustomerType,
		Mobile:        resp.Mobile,
		Email:         resp.Email,
		Gender:        resp.Gender,
		CallbackPhone: resp.CallbackPhone,
		Status:        resp.Status,
	}, nil
}
*/
func makeClientAddGajiDosenEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddGajiDosen",
		encodeAddGajiDosenRequest,
		decodeReadGajiDosenResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddGajiDosen")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddGajiDosen",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadGajiDosenByKdGGDosenEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadGajiDosenByKdGGDosen",
		encodeReadGajiDosenByKdGGDosenRequest,
		decodeReadGajiDosenByKdGGDosenRespones,
		pb.ReadGajiDosenByKdGGDosenResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadGajiDosenByKdGGDosen")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadGajiDosenByKdGGDosen",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadGajiDosenEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadGajiDosen",
		encodeReadGajiDosenRequest,
		decodeReadGajiDosenResponse,
		pb.ReadGajiDosenResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadGajiDosen")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadGajiDosen",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

/*
func makeClientUpdateCustomer(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateCustomer",
		encodeUpdateCustomerRequest,
		decodeCustomerResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateCustomer")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateCustomer",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadCustomerByEmail(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadCustomerByEmail",
		encodeReadCustomerByEmailRequest,
		decodeReadCustomerByEmailRespones,
		pb.ReadCustomerByEmailResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadCustomerByEmail")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadCustomerByEmail",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
*/
