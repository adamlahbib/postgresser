package main

import (
	"context"
	"time"

	pb "github.com/adamlahbib/postgresser/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:5000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// create a new postgresser client
	client := pb.NewPostgresServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create a new Postgres instance
	_, err = client.CreatePostgres(ctx, &pb.CreatePostgresRequest{
		Dbname:     "test",
		Username:   "adm",
		Password:   "adpass",
		Port:       5432,
		Replicas:   1,
		Accessmode: "ReadWriteOnce",
		Capacity:   "1Gi",
	})
	if err != nil {
		panic(err)
	}
}
