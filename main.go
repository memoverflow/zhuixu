package main

import (
	"book/model"
	"fmt"
	"log"
)

func main() {

	sites, err := model.RetrieveBooks()

	if err != nil {
		log.Fatal(err)
	}

	for _, site := range sites {

		fmt.Println(site.SiteName)

		for _, book := range site.Books {
			fmt.Println(book.Name)
		}

	}

}
