package parser

import (
	"encoding/json"
	"fmt"
	"github.com/mertozler/currency-cli/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	client = &http.Client{
		Timeout: 30 * time.Second,
	}
)

func GetExchange() (*model.Exchange, error) {
	resp, err := client.Get("https://api.genelpara.com/embed/doviz.json")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	model := model.Exchange{}
	if err := json.Unmarshal(body, &model); err != nil {
		fmt.Println("Can not unmarshal JSON")
		return nil, nil
	}

	return &model, nil
}
