package main

import (
	"context"
	"log"
	"net"

	pb "github.com/kodinggo/grpc-server-gb2/pb/grpc_server"
	"google.golang.org/grpc"
)

func main() {
	// Run grpc server
	grpcServer := grpc.NewServer()

	storygRPCHandler := NewStorygRPCHandler()

	pb.RegisterStoryServiceServer(grpcServer, storygRPCHandler)

	httpListener, err := net.Listen("tcp", ":5555")
	if err != nil {
		log.Panicf("create http listener: %v", err)
	}

	log.Println("grpc server is running")

	grpcServer.Serve(httpListener)
}

type StorygRPCHandler struct {
	pb.UnimplementedStoryServiceServer
}

func NewStorygRPCHandler() pb.StoryServiceServer {
	return &StorygRPCHandler{}
}

func (h *StorygRPCHandler) FindAll(context.Context, *pb.StoryFindAllRequest) (*pb.Stories, error) {
	stories := &pb.Stories{
		Stories: []*pb.Story{
			{
				Id:      1,
				Title:   "Example",
				Content: "Example",
			},
		},
	}

	return stories, nil
}
