package apicall

import (
	"ctracker/Monsters"
	"ctracker/redisdb"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/go-redis/redis"
)

func ServeApi() {
	http.HandleFunc("/pokemon/", GetPokemon)

	log.Fatal(http.ListenAndServe(":5000", nil))

}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	s := r.URL.RequestURI()
	fmt.Println(s)
	path := strings.SplitAfter(s, "/")
	pokemon := path[2]
	multiplePokemon := strings.Split(pokemon, ",")

	fmt.Println(path)
	var wg sync.WaitGroup
	wg.Add(len(multiplePokemon))
	ch := make(chan string, len(multiplePokemon))
	for _, s := range multiplePokemon {
		// fmt.Printf("%v - ",s)
		go PokemonRoutine(s, &wg, ch)
	}
	wg.Wait()
	close(ch)
	var responses = make([]Monsters.Pokemon, len(multiplePokemon))
	counter := 0
	for pokemon2 := range ch {
		prettyResponse := Monsters.Pokemon{}
		err := json.Unmarshal([]byte(pokemon2), &prettyResponse)
		if err != nil {
			fmt.Println(err)
		}
		responses[counter] = prettyResponse
		counter++
	}
	output := ""
	for i, s := range responses {
		// fmt.Fprintln(w, i, s.Name, s.Abilities)
		output = fmt.Sprintf("%v %v %v %v \n \n", output, i, s.Name, s.Abilities)
	}
	_, err := w.Write([]byte(output))
	if err == nil {
		fmt.Println(err)
	}
}
func PokemonRoutine(pokemon string, wg *sync.WaitGroup, ch chan string) {
	// prettyResponse := Monsters.Pokemon{}
	pokemon2, err := redisdb.Get(pokemon)
	if err != nil {
		if err == redis.Nil {
			body, err := UpdateFromApi(pokemon)
			if err != nil {
				fmt.Print(err)
			} else {
				pokemon2 = string(body)
				err = redisdb.Set(pokemon, pokemon2)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	ch <- pokemon2
	wg.Done()
}
func UpdateFromApi(pokemon string) ([]byte, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", pokemon)
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
	return body, err
}
