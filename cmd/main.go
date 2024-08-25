package main

import (
	"fmt"
	"log"

	appServer "github.com/Kinesocks/cat_facts_service/server"
)

func main() {

	server := appServer.NewServer()
	fmt.Println("Server is running...")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
