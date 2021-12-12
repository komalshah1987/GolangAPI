package v1

import (
	"encoding/json"
	"fmt"
	"os"

	model "github.com/komalshah1987/GolangAPI/model"
)

type ArticleRepository struct{}

var lstArticle = []model.Article{}

func GetAllArticle() model.Articles {
	f, err := os.Open("./model/articles")
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&lstArticle)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
