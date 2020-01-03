package main

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"log"
)

// you need to use capitalized field names for json.Marshall to work
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type User struct {
	Username  string   `json:"username"`
	Height    int      `json:"height"`
	Address   Address  `json:"address"`
	Interests []string `json:"interests"`
}

func main() {
	address := Address{
		Street:  "Main 1b/23",
		City:    "Big",
		Country: "Important",
	}
	user := User{
		Username: "Martin",
		Height:   171,
		Address:  address,
		Interests: []string{
			"singin",
			"dancing",
		},
	}

	log.Println("Initial user object")
	log.Println("")
	spew.Dump(&user)
	log.Println("")

	// marshall json
	data, err := json.Marshal(&user)
	if err != nil {
		panic(err.Error())
	}

	log.Println("Marshalled byte array")
	log.Println("")
	log.Println(data)
	log.Println("")

	// unmarshall json into interface{}
	var object interface{}
	err = json.Unmarshal(data, &object)
	if err != nil {
		panic(err.Error())
	}

	log.Println("UnMarshalled to interface{} user object")
	log.Println("")
	spew.Dump(object)
	log.Println("")

	log.Println("--------------------")
	data2 := []byte(`{"name":"Marshal", "id": 13, "interests":["fire", "firefighter"], "address":{"city":"Adventure Bay"}}`)
	log.Println("Marshalled byte array")
	log.Println("")
	log.Println(data2)
	log.Println("")

	err = json.Unmarshal(data2, &object)
	if err != nil {
		panic(err.Error())
	}

	log.Println("UnMarshalled to interface{} user object")
	log.Println("")
	spew.Dump(object)
	log.Println("")
}
