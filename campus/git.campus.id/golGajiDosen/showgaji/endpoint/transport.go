package endpoint

import (
	"context"

	scv "MCCampus/campus/git.campus.id/golGajiDosen/showgaji/server"

	pb "MCCampus/campus/git.campus.id/golGajiDosen/showgaji/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcGajiDosenServer struct {
	addGajiDosen             grpctransport.Handler
	readGajiDosenByKdGGDosen grpctransport.Handler
	readGajiDosen            grpctransport.Handler
	//updateCustomer       grpctransport.Handler
	//readCustomerByEmail  grpctransport.
}

func NewGRPCGajiDosenServer(endpoints GajiDosenEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.GajiDosenServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcGajiDosenServer{
		addGajiDosen: grpctransport.NewServer(endpoints.AddGajiDosenEndpoint,
			decodeAddGajiDosenRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddGajiDosen", logger)))...),
		readGajiDosenByKdGGDosen: grpctransport.NewServer(endpoints.ReadGajiDosenByKdGGDosenEndpoint,
			decodeReadGajiDosenByKdGGDosenRequest,
			encodeReadGajiDosenByKdGGDosenResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadGajiDosenByKdGGDosen", logger)))...),
		readGajiDosen: grpctransport.NewServer(endpoints.ReadGajiDosenEndpoint,
			decodeReadGajiDosenRequest,
			encodeReadGajiDosenResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadGajiDosen", logger)))...),
	}
}

func decodeAddGajiDosenRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddGajiDosenReq)
	return scv.GajiDosen{KdGGDosen: req.GetKdGGDosen(), JumlahGaji: req.GetJumlahGaji(), Status: req.GetStatus(), CreateBy: req.GetCreateBy(),
		CreateOn: req.GetCreateOn(), UpdateBy: req.GetUpdateBy(), UpdateOn: req.GetUpdateOn()}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcGajiDosenServer) AddGajiDosen(ctx oldcontext.Context, dosen *pb.AddGajiDosenReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addGajiDosen.ServeGRPC(ctx, dosen)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func decodeReadGajiDosenByKdGGDosenRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadGajiDosenByKdGGDosenReq)
	return scv.GajiDosen{KdGGDosen: req.KdGGDosen}, nil
}

func decodeReadGajiDosenRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func encodeReadGajiDosenByKdGGDosenResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.GajiDosen)
	return &pb.ReadGajiDosenByKdGGDosenResp{KdGGDosen: resp.KdGGDosen, JumlahGaji: resp.JumlahGaji, Keterangan: resp.Keterangan,
		Status: resp.Status, CreateBy: resp.CreateBy, CreateOn: resp.CreateOn, UpdateBy: resp.UpdateBy, UpdateOn: resp.UpdateOn}, nil
}

func encodeReadGajiDosenResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.GajiDosens)

	rsp := &pb.ReadGajiDosenResp{}

	for _, v := range resp {
		itm := &pb.ReadGajiDosenByKdGGDosenResp{
			KdGGDosen:  v.KdGGDosen,
			JumlahGaji: v.JumlahGaji,
			Keterangan: v.Keterangan,
			Status:     v.Status,
			CreateBy:   v.CreateBy,
			CreateOn:   v.CreateOn,
			UpdateBy:   v.UpdateBy,
			UpdateOn:   v.UpdateOn,
		}
		rsp.AllGajiDosen = append(rsp.AllGajiDosen, itm)
	}
	return rsp, nil
}

func (s *grpcGajiDosenServer) ReadGajiDosenByKdGGDosen(ctx oldcontext.Context, kdgajidosen *pb.ReadGajiDosenByKdGGDosenReq) (*pb.ReadGajiDosenByKdGGDosenResp, error) {
	_, resp, err := s.readGajiDosenByKdGGDosen.ServeGRPC(ctx, kdgajidosen)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadGajiDosenByKdGGDosenResp), nil
}

func (s *grpcGajiDosenServer) ReadGajiDosen(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadGajiDosenResp, error) {
	_, resp, err := s.readGajiDosen.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadGajiDosenResp), nil
}
