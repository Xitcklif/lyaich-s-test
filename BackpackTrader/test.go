package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

var keyPriceUrlRub = "http://steamcommunity.com/market/listings/440/Mann%20Co.%20Supply%20Crate%20Key"

func main() {
	keyPrice := getKeyPriceRub()
	fmt.Printf("Key price: %f\n", keyPrice)
	fmt.Scanln()
}

func getKeyPriceRub() float64 {
	resp, err := http.Get(keyPriceUrlRub)
	if err != nil {
		fmt.Printf("Can't get key price: %s\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("Can't read body from key price page: %s\n", err)
		os.Exit(1)
	}
	//продается, начальная цена: <span class="market_commodity_orders_header_promote">156,55 pуб.</span></div>
	re := regexp.MustCompile("начальная цена:.+>([0-9,]+)  pуб.</span></div>")
	priceStrings := re.FindStringSubmatch(string(b))
	if len(priceStrings) == 0 {
		fmt.Printf("Can't find price on the page\n")
		os.Exit(1)
	}
	price, err := strconv.ParseFloat(priceStrings[1], 64)
	if err != nil {
		fmt.Printf("Can't parse price string: %s\n", price)
		os.Exit(1)
	}

	return price
}
