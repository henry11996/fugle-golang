package main

import (
	"log"

	"github.com/henry11996/fugle-golang/fugle"
)

func main() {
	options := &fugle.ClientOption{
		ApiToken: "your_token",
	}

	client, _ := fugle.NewFugleClient(*options)

	data, _ := client.Meta("2330", false)

	log.Print(data.APIVersion)
}
