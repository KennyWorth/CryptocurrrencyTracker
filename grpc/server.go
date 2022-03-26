package ctrackergrpc

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"ctracker/pb"
	"ctracker/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = "50051"
)

func NewGrpcServer() {
	grpcServer := grpc.NewServer()
	service := service.NewCryptoGrpcService()
	pb.RegisterGetAllCoinsServer(grpcServer, service)
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatalf("failed to start tcp listener on %v , %v", port, err)
	}
	time := time.Now()
	fmt.Println("starting gRPC server")
	fmt.Printf("Starting. Time: %v", time)
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Println("failed to serve gRPC server")
	}
	fmt.Println("shutdown completed")
}

func UpdateFromApi(command string) (pb.CoinResponse, error) {
	url := fmt.Sprintf("api.coingecko.com/api/v3/%v", command)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return pb.CoinResponse{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		return pb.CoinResponse{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Print(body)
	return pb.CoinResponse{}, err
}
