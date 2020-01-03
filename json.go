package main

import (
	"github.com/davecgh/go-spew/spew"
)

type Address struct {
	street  string `json:"street,omitempty"`
	city    string `json:"city,omitempty"`
	country string `json:"country,omitempty"`
}

type User struct {
	username  string   `json:"username,omitempty"`
	height    int      `json:"height,omitempty"`
	address   Address  `json:"address,omitempty"`
	interests []string `json:"interests,omitempty"`
}

func main() {
	address := Address{
		street: "Main 1b/23",
		city: "Big",
		country: "Important",
	}
	user := User{
		username: "Martin",
		height: 171,
		address: address,
		interests: []string{
			"singin",
			"dancing",
		},
	}

	spew.Dump(&user)
}
