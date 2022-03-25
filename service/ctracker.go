package service

import (
	"ctracker/cache"
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
