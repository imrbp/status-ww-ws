package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"microservice/entity"
)

func ReadJson(filePath string, data *entity.Data) {

	// defer wg.Done()

	byteValue, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byteValue, &data)

	if err != nil {
		panic(err)
	}
}

func WriteJson(filePath string, data *entity.Data) {
	// defer wg.Done()
	data.Status.Water = int(GenerateRand(1, 100))
	data.Status.Wind = int(GenerateRand(1, 100))

	content, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("data.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
