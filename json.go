package main

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"log"
)

type Address struct {
	street  string `json:"street"`
	city    string `json:"city"`
	country string `json:"country"`
}

type User struct {
	username  string   `json:"username"`
	height    int      `json:"height"`
	address   Address  `json:"address"`
	interests []string `json:"interests"`
}

func main() {
	address := Address{
		street:  "Main 1b/23",
		city:    "Big",
		country: "Important",
	}
	user := User{
		username: "Martin",
		height:   171,
		address:  address,
		interests: []string{
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
