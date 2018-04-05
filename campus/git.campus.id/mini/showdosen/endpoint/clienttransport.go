package endpoint

import (
	"context"
	"time"

	svc "MCCampus/campus/git.campus.id/mini/showdosen/server"

	pb "MCCampus/campus/git.campus.id/mini/showdosen/grpc"

	util "MCCampus/campus/git.campus.id/mini/util/grpc"
	disc "MCCampus/campus/git.campus.id/mini/util/microservice"

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
	grpcName = "grpc.DosenService"
)

func NewGRPCDosenClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.DosenService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addDosenEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddDosenEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addDosenEp = retry
	}

	var readDosenByKdDosenEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadDosenByKdDosenEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readDosenByKdDosenEp = retry
	}

	var readDosenByKeteranganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadDosenByKeteranganEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readDosenByKeteranganEp = retry
	}

	var readDosenEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadDosenEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readDosenEp = retry
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
	return DosenEndpoint{AddDosenEndpoint: addDosenEp, ReadDosenByKdDosenEndpoint: readDosenByKdDosenEp, ReadDosenByKeteranganEndpoint: readDosenByKeteranganEp,
		ReadDosenEndpoint: readDosenEp}, nil
}

func encodeAddDosenRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Dosen)
	return &pb.AddDosenReq{
		KdDosen:   req.KdDosen,
		NamaDosen: req.NamaDosen,
		Status:    req.Status,
		CreateBy:  req.CreateBy,
	}, nil
}

func encodeReadDosenByKdDosenRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Dosen)
	return &pb.ReadDosenByKdDosenReq{KdDosen: req.KdDosen}, nil
}

func encodeReadDosenByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Dosen)
	return &pb.ReadDosenByKeteranganReq{Keterangan: req.Keterangan}, nil
}

func encodeReadDosenRequest(_ context.Context, request interface{}) (interface{}, error) {
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

func decodeReadDosenByKdDosenRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadDosenByKdDosenResp)
	return svc.Dosen{
		KdDosen:   resp.KdDosen,
		NamaDosen: resp.NamaDosen,
		Status:    resp.Status,
		CreateBy:  resp.CreateBy,
	}, nil
}

func decodeReadDosenByKeteranganRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadDosenByKeteranganResp)
	return svc.Dosen{
		Keterangan: resp.Keterangan,
		NamaDosen:  resp.NamaDosen,
		Status:     resp.Status,
		CreateBy:   resp.CreateBy,
	}, nil
}

func decodeReadDosenResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadDosenResp)
	var rsp svc.Dosens

	for _, v := range resp.AllDosen {
		itm := svc.Dosen{
			KdDosen:   v.KdDosen,
			NamaDosen: v.NamaDosen,
			Status:    v.Status,
			CreateBy:  v.CreateBy,
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
func makeClientAddDosenEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddDosen",
		encodeAddDosenRequest,
		decodeDosenResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddDosen")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddDosen",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadDosenByKdDosenEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadDosenByKdDosen",
		encodeReadDosenByKdDosenRequest,
		decodeReadDosenByKdDosenRespones,
		pb.ReadDosenByKdDosenResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadDosenByKdDosen")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadDosenByKdDosen",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadDosenByKeteranganEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadDosenByKeterangan",
		encodeReadDosenByKeteranganRequest,
		decodeReadDosenByKeteranganRespones,
		pb.ReadDosenByKeteranganResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadDosenByKeterangan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadDosenByKeterangan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadDosenEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadDosen",
		encodeReadDosenRequest,
		decodeReadDosenResponse,
		pb.ReadDosenResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadDosen")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadDosen",
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
