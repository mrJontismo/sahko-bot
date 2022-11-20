package price

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Prices struct {
    Min		string	`json:"min"`
    Max 	string	`json:"max"`
    Now 	string 	`json:"now"`
	Avg		string	`json:"avg"`
	Avg_28	string	`json:"avg_28"`
}

func GetPrice() Prices {
	response, err := http.PostForm("https://sahko.tk/api.php", url.Values{"mode": {"get_prices"}})

	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)

	var priceData Prices
	json.Unmarshal(content, &priceData)
	
	return priceData
}