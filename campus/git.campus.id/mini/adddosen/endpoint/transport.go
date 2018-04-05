package endpoint

import (
	"context"

	pb "MCCampus/campus/git.campus.id/mini/adddosen/grpc"
	scv "MCCampus/campus/git.campus.id/mini/adddosen/server"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcDosenServer struct {
	addDosen grpctransport.Handler
	//readCustomerByMobile grpctransport.Handler
	//readCustomer         grpctransport.Handler
	//updateCustomer       grpctransport.Handler
	//readCustomerByEmail  grpctransport.Handler
}

func NewGRPCDosenServer(endpoints DosenEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.DosenServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcDosenServer{
		addDosen: grpctransport.NewServer(endpoints.AddDosenEndpoint,
			decodeAddDosenRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddDosen", logger)))...),
	}
}

func decodeAddDosenRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddDosenReq)
	return scv.Dosen{KdDosen: req.GetKdDosen(), NamaDosen: req.GetNamaDosen(),
		Status: req.GetStatus(), CreateBy: req.GetCreateBy()}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcDosenServer) AddDosen(ctx oldcontext.Context, dosen *pb.AddDosenReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addDosen.ServeGRPC(ctx, dosen)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}
