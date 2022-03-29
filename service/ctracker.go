package service

import (
	"context"
	"ctracker/apicall"
	cache "ctracker/cache"
	"ctracker/pb"
	"fmt"
	"strconv"
)

type CryptoService struct {
	// holds the service and cache
	cache cache.CryptoCache
}

func NewService() CryptoService {
	c := cache.NewCryptoCache()
	// build cluster if not present
	// c.NewCluster()
	return CryptoService{
		cache: c,
	}
}

type CryptoGrpcService struct {
	// holds the service and cache
	pb.UnimplementedGetCoinServer
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
func (s *CryptoGrpcService) Coins(ctx context.Context, empty *pb.Empty) (*pb.CoinListResponse, error) {
	response, err := apicall.CoinListRoutine("coins/list")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (v *CryptoGrpcService) CoinPrice(ctx context.Context, marketPrice *pb.MarketPriceRequest) (*pb.MarketPriceResponse, error) {
	id := marketPrice.Id
	vsCurrency := marketPrice.VsCurrency
	daysRequest := strconv.FormatInt(marketPrice.Days, 10)
	apiCommandFormat := ("coins/" + id + "/market_chart?vs_currency=" + vsCurrency + "&days=" + daysRequest)
	fmt.Println("\nRequest: " + apiCommandFormat)
	response, err := apicall.MarketPriceRoutine(apiCommandFormat, id)
	if err != nil {
		return nil, err
	}
	fmt.Println("no errors before return response")
	// fmt.Println(response)
	return response, nil
}
