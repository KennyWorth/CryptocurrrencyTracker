package apicall

import (
	"ctracker/pb"
	db "ctracker/redisdb"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"google.golang.org/protobuf/proto"
)

//struct must match the incoming response data to successfully marshal
type PriceData struct {
	Prices       [][]float64 `json:"prices"`
	MarketCaps   [][]float64 `json:"market_caps"`
	TotalVolumes [][]float64 `json:"total_volumes"`
}

//broke out api call to separate function returning byte array which can be parsed
func ApiCall(command string) ([]byte, error) {
	url := fmt.Sprintf("http://api.coingecko.com/api/v3/%v", command)
	fmt.Printf("URL in UpdateFromCoinGecko: %v\n", url)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Printf("error creating request::: %v\n", err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("error invoking command::: %v\n", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error reading body::: %v\n", err)
	}
	return body, err
}

//name should be changed to parsepricedata
func UpdatePrice(body []byte) PriceData {
	var priceData PriceData
	err := json.Unmarshal(body, &priceData)
	if err != nil {
		fmt.Println(err, " --unmarshaling body data to priceData in UpdatePrice--")
	}
	return priceData
}

func GetMarketPrice(marketPrices PriceData) *pb.MarketPriceResponse {
	response := &pb.MarketPriceResponse{}
	price := response.Prices
	for _, priceArray := range marketPrices.Prices {
		price = append(price, &pb.PricePoints{
			Prices:    float64(priceArray[1]),
			Timestamp: int64(priceArray[0]),
		})
	}
	response.Prices = price
	return response
}

func MarketPriceRoutine(urlCommand string, coinName string) (*pb.MarketPriceResponse, error) {
	var priceResponse = &pb.MarketPriceResponse{}
	//TODO: implement cache value marshal and return,
	storedExpirationString, err := db.Get(coinName)
	if err != nil && err == redis.Nil {
		fmt.Println("Line 73")
		return UpdateCache(urlCommand, coinName)
	}
	if err == nil {
		fmt.Println("Line 76")
		storedTimestamp, _ := time.Parse(time.RFC1123, storedExpirationString)
		fmt.Println("Time stamp Expired? :", storedTimestamp.Before(time.Now()), storedTimestamp, "Current Time", time.Now(), "$###", storedExpirationString)
		if storedTimestamp.Before(time.Now()) {
			fmt.Println("TimeStamp expired, getting refreshed Data")
			return UpdateCache(urlCommand, coinName)

		} else {
			data, _ := db.GetByteArray(coinName + ":Prices")
			priceResponse, err = parseCachedPriceData(data), nil
			return priceResponse, err
		}
	}
	return priceResponse, nil
}

func UpdateCache(urlCommand, coinName string) (*pb.MarketPriceResponse, error) {
	fmt.Println("--no cache value--")
	body, _ := ApiCall(urlCommand)
	priceData := UpdatePrice(body)
	priceResponse := GetMarketPrice(priceData)
	toStorePrice, err := proto.Marshal(priceResponse)
	if err != nil {
		fmt.Printf("2 + %v", err)
		return nil, err
	}
	expirationTime := time.Now().Add(time.Hour).Format(time.RFC1123)
	fmt.Printf("Expiration Time: %v", expirationTime)
	err = db.Set(coinName, expirationTime)
	if err != nil {
		return nil, err
	}
	err = db.SetByteArray(coinName+":Prices", toStorePrice)
	if err != nil {
		return nil, err
	}
	return priceResponse, nil
}

func parseCachedPriceData(data []byte) *pb.MarketPriceResponse {
	var response = &pb.MarketPriceResponse{}
	proto.Unmarshal(data, response)
	return response
}
