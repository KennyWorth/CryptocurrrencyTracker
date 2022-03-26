package service

import (
	"context"
	"ctracker/apicall"
	"ctracker/cache"
	"ctracker/pb"
	"fmt"
)

type CryptoService struct {
	// holds the service and cache
	cache cache.CryptoCache
}

func NewService() CryptoService {
	fmt.Println("Hello, World!")
	c := cache.NewCryptoCache()
	// build cluster if not present
	c.NewCluster()
	return CryptoService{
		cache: c,
	}
}

type CryptoGrpcService struct {
	// holds the service and cache
	pb.UnimplementedGetAllCoinsServer
	// cache cache.CryptoCache
}

func NewCryptoGrpcService() *CryptoGrpcService {
	// c := cache.NewCryptoCache()
	// build cluster if not present
	// c.NewCluster()
	return &CryptoGrpcService{
		// cache: c,
	}
}
func (s *CryptoGrpcService) Coins(ctx context.Context, empty *pb.Empty) (*pb.CoinResponse, error) {
	fmt.Print("Coins here")
	response, err := apicall.CoinListRoutine("coins/list")
	if err != nil {
		return nil, err
	}
	fmt.Println("no errors before return response")
	return response, nil
}
