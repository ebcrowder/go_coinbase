package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fabioberger/coinbase-go"
)

func main() {
	c := coinbase.ApiKeyClient(os.Getenv("COINBASE_KEY"), os.Getenv("COINBASE_SECRET"))

	user, err := c.GetUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name: ", user.Name)
	fmt.Println("Email: ", user.Email)

	balance, err := c.GetBalance()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current balance is: %f BTC", balance)
}
