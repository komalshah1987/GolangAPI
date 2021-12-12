package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/komalshah1987/GolangAPI/model"
)

var spool = []model.Article{}

// GetData will get the data from file as marshal json
func InitializeData() error {
	f, err := os.Open("./moovie")
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&spool)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
