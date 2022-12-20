package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/expbuild/expbuild/proto/gen/remote_execution"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCapabilitiesClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := pb.GetCapabilitiesRequest{}

	r, err := client.GetCapabilities(ctx, &req)
	fmt.Println(r.String())
}
