package exchange

import (
	"encoding/json"
	"fmt"
	"github.com/mertozler/currency-cli/model"
	"io"
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
		log.Println(err)
	}
	defer resp.Body.Close()

	model, err := parseBody(resp.Body)

	if err != nil {
		return nil, err
	}
	return model, nil
}

func parseBody(data io.ReadCloser) (*model.Exchange, error) {
	body, err := ioutil.ReadAll(data)
	if err != nil {
		fmt.Println("Error while parsing body: ", err)
		return nil, err
	}

	model := model.Exchange{}
	if err := json.Unmarshal(body, &model); err != nil {
		fmt.Println("Can not unmarshal JSON")
		return nil, nil
	}
	return &model, nil
}