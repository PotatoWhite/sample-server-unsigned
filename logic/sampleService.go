package logic

import (
	"context"
	"log"
	pb "potato/demo/sample-server/grpc/go/proto"
)

type Service struct {
}

func (a *Service) GetAllSamples(ctx context.Context, request *pb.GetAllSamplesRequest) (*pb.GetAllSamplesResponse, error) {
	log.Println(request)

	samples := make([]*pb.Sample, 10)

	var index uint32
	for ; index < 10; index++ {
		samples[index] = &pb.Sample{
			Data: request.GetData(),
			Id:   request.Max + index,
		}

		log.Println(request.Max + index)
	}

	return &pb.GetAllSamplesResponse{
		Samples: samples,
	}, nil
}
