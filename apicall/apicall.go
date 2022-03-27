package apicall

import (
	"ctracker/pb"
	"ctracker/redisdb"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"google.golang.org/protobuf/proto"

	"github.com/go-redis/redis"
)

type CoinIds []struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func UpdateFromCoinGeckoApi(command string) (CoinIds, error) {
	url := fmt.Sprintf("http://api.coingecko.com/api/v3/%v", command)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var coinIds CoinIds
	err = json.Unmarshal(body, &coinIds)
	if err != nil {
		fmt.Println(err)
	}
	return coinIds, err
}

func GetCoinList(coinIds CoinIds) *pb.CoinListResponse {
	response := &pb.CoinListResponse{}
	coins := response.CoinList
	for _, v := range coinIds {
		coins = append(coins, &pb.CoinIds{
			Id:     v.ID,
			Symbol: v.Symbol,
			Name:   v.Name,
		})

	}
	response.CoinList = coins
	return response
}

func CoinListRoutine(command string) (*pb.CoinListResponse, error) {
	coinListFromRedis, err := redisdb.GetByteArray(command)
	if err != nil {
		if err == redis.Nil {
			body, err := UpdateFromCoinGeckoApi(command)
			if err != nil {
				fmt.Print(err)
				return nil, err
			} else {
				toStore := GetCoinList(body)
				marshalledToStore, err := proto.Marshal(toStore)
				if err != nil {
					fmt.Println(err)
					return nil, err
				}
				err = redisdb.SetByteArray(command, marshalledToStore)
				if err != nil {
					fmt.Println(err)
					return nil, err
				}
			}
		}
	}
	var response pb.CoinListResponse
	proto.Unmarshal(coinListFromRedis, &response)
	return &response, nil
}
