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

// func ServeApi() {
// 	http.HandleFunc("/pokemon/", GetPokemon)

// 	log.Fatal(http.ListenAndServe(":5000", nil))

// }

// func GetPokemon(w http.ResponseWriter, r *http.Request) {
// 	s := r.URL.RequestURI()
// 	fmt.Println(s)
// 	path := strings.SplitAfter(s, "/")
// 	pokemon := path[2]
// 	multiplePokemon := strings.Split(pokemon, ",")

// 	fmt.Println(path)
// 	var wg sync.WaitGroup
// 	wg.Add(len(multiplePokemon))
// 	ch := make(chan string, len(multiplePokemon))
// 	for _, s := range multiplePokemon {
// 		// fmt.Printf("%v - ",s)
// 		go PokemonRoutine(s, &wg, ch)
// 	}
// 	wg.Wait()
// 	close(ch)
// 	var responses = make([]Monsters.Pokemon, len(multiplePokemon))
// 	counter := 0
// 	for pokemon2 := range ch {
// 		prettyResponse := Monsters.Pokemon{}
// 		err := json.Unmarshal([]byte(pokemon2), &prettyResponse)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		responses[counter] = prettyResponse
// 		counter++
// 	}
// 	output := ""
// 	for i, s := range responses {
// 		// fmt.Fprintln(w, i, s.Name, s.Abilities)
// 		output = fmt.Sprintf("%v %v %v %v \n \n", output, i, s.Name, s.Abilities)
// 	}
// 	_, err := w.Write([]byte(output))
// 	if err == nil {
// 		fmt.Println(err)
// 	}
// }

// func PokemonRoutine(pokemon string, wg *sync.WaitGroup, ch chan string) {
// 	// prettyResponse := Monsters.Pokemon{}
// 	pokemon2, err := redisdb.Get(pokemon)
// 	if err != nil {
// 		if err == redis.Nil {
// 			body, err := UpdateFromApi(pokemon)
// 			if err != nil {
// 				fmt.Print(err)
// 			} else {
// 				pokemon2 = string(body)
// 				err = redisdb.Set(pokemon, pokemon2)
// 				if err != nil {
// 					fmt.Println(err)
// 				}
// 			}
// 		}
// 	}
// 	ch <- pokemon2
// 	wg.Done()
// }
func UpdateFromApi(command string) (CoinIds, error) {
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

func GetCoinList(coinIds CoinIds) *pb.CoinResponse {
	response := &pb.CoinResponse{}
	coins := response.CoinList
	// fmt.Printf("Coins in GetCoinList: \n \n \n %v", coins)
	for _, v := range coinIds {
		// fmt.Printf("ID in loop: \n \n \n %v %v %v", v.ID, v.Symbol, v.Name)
		coins = append(coins, &pb.CoinIds{
			Id:    v.ID,
			Token: v.Symbol,
			Name:  v.Name,
		})
	}
	response.CoinList = coins
	// fmt.Printf("Response in GetCoinList: \n \n \n %v", response)
	return response
}

func CoinListRoutine(command string) (*pb.CoinResponse, error) {
	// prettyResponse := Monsters.Pokemon{}
	coinListFromRedis, err := redisdb.GetByteArray(command)
	if err != nil {
		if err == redis.Nil {
			body, err := UpdateFromApi(command)
			if err != nil {
				fmt.Print(err)
				return nil, err
			} else {
				toStore := GetCoinList(body)
				marshalledToStore, err := proto.Marshal(toStore)
				fmt.Printf("Value of %v \n \n \n", toStore)
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
	var response pb.CoinResponse
	proto.Unmarshal(coinListFromRedis, &response)
	return &response, nil
}
