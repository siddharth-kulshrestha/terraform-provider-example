package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spaceapegames/terraform-provider-example/api/server"
)

func main() {
	seed := flag.String("seed", "", "a file location with some data in JSON form to seed the server content")
	port := flag.String("port", "3001", "port to execute the request....")
	flag.Parse()

	items := map[string]server.Item{}

	if *seed != "" {
		seedData, err := ioutil.ReadFile(*seed)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(seedData, &items)
		if err != nil {
			log.Fatal(err)
		}
	}
	addr := fmt.Sprintf("localhost:%s", *port)
	log.Printf("Starting server atL %s \n", addr)
	itemService := server.NewService(addr, items)
	err := itemService.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
