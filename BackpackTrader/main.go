package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
)

var keyPriceUrl = "http://www.trade.tf/"

func main() {
	keyPrice := getKeyPrice()
	fmt.Printf("Key price: %f\n", keyPrice)

	fileName := os.Args[1]
	fmt.Printf("Processing csv file: %s\n", fileName)

	fh, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: can't open file: %s, msg: %s\n", fileName, err)
		os.Exit(1)
	}

	csvReader := csv.NewReader(fh)
	records, err := csvReader.ReadAll()

	if err != nil {
		fmt.Printf("Error: can't parse file: %s, msg: %s\n", fileName, err)
		os.Exit(1)
	}

	fmt.Println(len(records))
	for _, row := range records {
		url1, _ := url.QueryUnescape(row[0])
		url2, _ := url.QueryUnescape(row[1])

		fmt.Printf("url: %s  url2: %s\n", url1, url2)
	}
}

func getKeyPrice() float64 {
	resp, err := http.Get(keyPriceUrl)
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
	//	                                <a href="/mybots/buyers/Mann%20Co.%20Supply%20Crate%20Key" class="btn btn-info">Sell it</a> for 19.66 ref<br/>
	re := regexp.MustCompile("Key.+Sell it</a> for ([0-9.]+) ref")
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
