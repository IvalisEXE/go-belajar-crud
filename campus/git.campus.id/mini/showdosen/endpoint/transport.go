package endpoint

import (
	"context"

	scv "MCCampus/campus/git.campus.id/mini/showdosen/server"

	pb "MCCampus/campus/git.campus.id/mini/showdosen/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcDosenServer struct {
	addDosen              grpctransport.Handler
	readDosenByKdDosen    grpctransport.Handler
	readDosenByKeterangan grpctransport.Handler
	readDosen             grpctransport.Handler

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
		readDosenByKdDosen: grpctransport.NewServer(endpoints.ReadDosenByKdDosenEndpoint,
			decodeReadDosenByKdDosenRequest,
			encodeReadDosenByKdDosenResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadDosenByKdDosen", logger)))...),
		readDosenByKeterangan: grpctransport.NewServer(endpoints.ReadDosenByKeteranganEndpoint,
			decodeReadDosenByKeteranganRequest,
			encodeReadDosenByKeteranganResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadDosenByKeterangan", logger)))...),
		readDosen: grpctransport.NewServer(endpoints.ReadDosenEndpoint,
			decodeReadDosenRequest,
			encodeReadDosenResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadDosen", logger)))...),
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

func decodeReadDosenByKdDosenRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadDosenByKdDosenReq)
	return scv.Dosen{KdDosen: req.KdDosen}, nil
}

func decodeReadDosenByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadDosenByKeteranganReq)
	return scv.Dosen{Keterangan: req.Keterangan}, nil
}

func decodeReadDosenRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func encodeReadDosenByKdDosenResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Dosen)
	return &pb.ReadDosenByKdDosenResp{KdDosen: resp.KdDosen, NamaDosen: resp.NamaDosen,
		Status: resp.Status, CreateBy: resp.CreateBy}, nil
}

func encodeReadDosenByKeteranganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Dosen)
	return &pb.ReadDosenByKeteranganResp{Keterangan: resp.Keterangan, NamaDosen: resp.NamaDosen,
		Status: resp.Status, CreateBy: resp.CreateBy}, nil
}

func encodeReadDosenResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Dosens)

	rsp := &pb.ReadDosenResp{}

	for _, v := range resp {
		itm := &pb.ReadDosenByKdDosenResp{
			KdDosen:   v.KdDosen,
			NamaDosen: v.NamaDosen,
			Status:    v.Status,
			CreateBy:  v.CreateBy,
		}
		rsp.AllDosen = append(rsp.AllDosen, itm)
	}
	return rsp, nil
}

func (s *grpcDosenServer) ReadDosenByKdDosen(ctx oldcontext.Context, kddosen *pb.ReadDosenByKdDosenReq) (*pb.ReadDosenByKdDosenResp, error) {
	_, resp, err := s.readDosenByKdDosen.ServeGRPC(ctx, kddosen)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadDosenByKdDosenResp), nil
}

func (s *grpcDosenServer) ReadDosenByKeterangan(ctx oldcontext.Context, ktrg *pb.ReadDosenByKeteranganReq) (*pb.ReadDosenByKeteranganResp, error) {
	_, resp, err := s.readDosenByKeterangan.ServeGRPC(ctx, ktrg)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadDosenByKeteranganResp), nil
}

func (s *grpcDosenServer) ReadDosen(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadDosenResp, error) {
	_, resp, err := s.readDosen.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadDosenResp), nil
}
