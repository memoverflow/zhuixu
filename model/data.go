package model

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Chapter struct {
	Link  string
	Title string
	Text  string
}

type Book struct {
	Name     string `json:"name"`
	Author   string `json:"author"`
	Path     string `json:"path"`
	Reslover string `json:"reslover"`
	Pool     int    `json:"pool"`
	Cover    string
	Chapters []Chapter
}

type Site struct {
	SiteName string `json:"sitename"`
	Site     string `json:"site"`
	Resolver string `json:"resolver"`
	Books    []Book `json:"books"`
}

func RetrieveBooks() ([]*Site, error) {
	file, err := os.Open(dataFile)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var sites []*Site

	err = json.NewDecoder(file).Decode(&sites)

	return sites, err
}
