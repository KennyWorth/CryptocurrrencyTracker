package cache

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"
)

type CryptoCache struct {
	cache *elasticache.ElastiCache
}

func NewCache() *elasticache.ElastiCache {
	mySession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.UsEast2RegionID),
	}))

	// Create an ElastiCache client from just a session.
	svc := elasticache.New(mySession)

	return svc
}

func NewCryptoCache() CryptoCache {
	c := NewCache()
	return CryptoCache{
		cache: c,
	}
}

func (s *CryptoCache) NewCluster() {

	input := &elasticache.CreateCacheClusterInput{
		CacheClusterId:       aws.String("my-redis-cluster"),
		CacheNodeType:        aws.String("cache.t2.micro"),
		CacheSubnetGroupName: aws.String("default"),
		Engine:               aws.String("redis"),
		EngineVersion:        aws.String("6.2"),
		NumCacheNodes:        aws.Int64(1),
		Port:                 aws.Int64(11211),
	}

	result, err := s.cache.CreateCacheCluster(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticache.ErrCodeReplicationGroupNotFoundFault:
				fmt.Println(elasticache.ErrCodeReplicationGroupNotFoundFault, aerr.Error())
			case elasticache.ErrCodeInvalidReplicationGroupStateFault:
				fmt.Println(elasticache.ErrCodeInvalidReplicationGroupStateFault, aerr.Error())
			case elasticache.ErrCodeCacheClusterAlreadyExistsFault:
				fmt.Println(elasticache.ErrCodeCacheClusterAlreadyExistsFault, aerr.Error())
				return
			case elasticache.ErrCodeInsufficientCacheClusterCapacityFault:
				fmt.Println(elasticache.ErrCodeInsufficientCacheClusterCapacityFault, aerr.Error())
			case elasticache.ErrCodeCacheSecurityGroupNotFoundFault:
				fmt.Println(elasticache.ErrCodeCacheSecurityGroupNotFoundFault, aerr.Error())
			case elasticache.ErrCodeCacheSubnetGroupNotFoundFault:
				fmt.Println(elasticache.ErrCodeCacheSubnetGroupNotFoundFault, aerr.Error())
			case elasticache.ErrCodeClusterQuotaForCustomerExceededFault:
				fmt.Println(elasticache.ErrCodeClusterQuotaForCustomerExceededFault, aerr.Error())
			case elasticache.ErrCodeNodeQuotaForClusterExceededFault:
				fmt.Println(elasticache.ErrCodeNodeQuotaForClusterExceededFault, aerr.Error())
			case elasticache.ErrCodeNodeQuotaForCustomerExceededFault:
				fmt.Println(elasticache.ErrCodeNodeQuotaForCustomerExceededFault, aerr.Error())
			case elasticache.ErrCodeCacheParameterGroupNotFoundFault:
				fmt.Println(elasticache.ErrCodeCacheParameterGroupNotFoundFault, aerr.Error())
			case elasticache.ErrCodeInvalidVPCNetworkStateFault:
				fmt.Println(elasticache.ErrCodeInvalidVPCNetworkStateFault, aerr.Error())
			case elasticache.ErrCodeTagQuotaPerResourceExceeded:
				fmt.Println(elasticache.ErrCodeTagQuotaPerResourceExceeded, aerr.Error())
			case elasticache.ErrCodeInvalidParameterValueException:
				fmt.Println(elasticache.ErrCodeInvalidParameterValueException, aerr.Error())
			case elasticache.ErrCodeInvalidParameterCombinationException:
				fmt.Println(elasticache.ErrCodeInvalidParameterCombinationException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
