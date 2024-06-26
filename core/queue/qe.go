package queue

import (
	"encoding/json"
	"log"
	"os"
	"test_task/models"
)

func Get(tName string) *[]models.Action {
	b, err := os.ReadFile(tName)
	if err != nil {
		log.Println(err)
	}
	var data models.Data

	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Fatal(err, data)
	}

	return &data.Actions
}

func ToFile(name string, data *[]models.Action) error {
	var res models.Data
	res.Actions = *data
	b, err := json.Marshal(&res)
	if err != nil {
		return err
	}

	err = record(name, &b)
	return err
}

func record(name string, data *[]byte) error {
	err := os.WriteFile(name, *data, 0666)
	return err
}
