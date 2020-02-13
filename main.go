package main

import (
	"log"
	"sync"

	"github.com/xuren87/zhuixu/model"
)

const iv = 790

func main() {

	site, err := model.RetrieveBooks()

	if err != nil {
		log.Fatal(err)
	}

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(site.PlayList))

	for index, p := range site.PlayList {
		if index < iv {
			continue
		}
		model.DownloadFile(p)
	}
}
