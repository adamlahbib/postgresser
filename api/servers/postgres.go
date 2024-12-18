package servers

import (
	"context"

	pb "github.com/adamlahbib/postgresser/api/proto"
	"github.com/adamlahbib/postgresser/models"
	"github.com/adamlahbib/postgresser/services"
)

type PostgresServer struct {
	postgresService services.Service
	pb.UnimplementedPostgresServiceServer
}

func NewPostgres(postgresService services.Service) pb.PostgresServiceServer {
	return &PostgresServer{
		postgresService: postgresService,
	}
}

func (p *PostgresServer) CreatePostgres(ctx context.Context, req *pb.CreatePostgresRequest) (*pb.CreatePostgresResponse, error) {
	resp, err := p.postgresService.Create(ctx, models.CreateRequest{
		DBName:     req.GetDbname(),
		Username:   req.GetUsername(),
		Password:   req.GetPassword(),
		Port:       req.GetPort(),
		Replicas:   req.GetReplicas(),
		Capacity:   req.GetCapacity(),
		AccessMode: req.GetAccessmode(),
	})
	return &pb.CreatePostgresResponse{
		Id: resp.Id,
	}, err
}

func (p *PostgresServer) UpdatePostgres(ctx context.Context, req *pb.UpdatePostgresRequest) (*pb.UpdatePostgresResponse, error) {
	err := p.postgresService.Update(ctx, models.UpdateRequest{
		Id:       req.GetId(),
		Replicas: req.GetReplicas(),
	})
	return &pb.UpdatePostgresResponse{}, err
}

func (p *PostgresServer) DeletePostgres(ctx context.Context, req *pb.DeletePostgresRequest) (*pb.DeletePostgresResponse, error) {
	err := p.postgresService.Delete(ctx, models.DeleteRequest{
		Id: req.GetId(),
	})
	return &pb.DeletePostgresResponse{}, err
}
