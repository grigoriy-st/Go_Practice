package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Preson struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
}

func main() {
	os.Chdir("Work\ with\ markup\ languages/JSON")
	fmt.Println(os.Chdir("Work with markup languages/JSON")
	file, err := os.Open("4.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var people []Preson
	err = json.Unmarshal(data, &people)
	if err != nil {
		log.Fatal(err)
	}

	for _, person := range people {
		fmt.Printf("Name: %s\nEmail: %s\nPhone:%s\nWebsite: %s\n\n",
			person.Name, person.Address, person.Phone, person.Website)
	}
}
