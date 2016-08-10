package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"

	"github.com/tealeg/xlsx"
)

var keyPriceUrl = "http://www.trade.tf/"

func main() {
	keyPrice := getKeyPrice()
	fmt.Printf("Key price: %f\n", keyPrice)

	excelFileName := os.Args[1]
	fmt.Printf("Processing excel file: %s\n", excelFileName)

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("Error: can't open excel file: %s, msg: %s\n", excelFileName, err)
		os.Exit(1)
	}

	sheet := xlFile.Sheets[0]
	//sheet.Rows[1].Cells[8].Value = strconv.FormatFloat(keyPrice, 'f', -1, 64)
	//sheet.Rows[1].Cells[8].Value = "1"
	xlFile.Save(excelFileName + ".new.xlsx")

	for rowNumber, row := range sheet.Rows {
		if len(row.Cells) < 6 {
			break
		}
		itemUrl := row.Cells[5].Value
		itemUrl, _ = url.QueryUnescape(itemUrl)
		fmt.Printf("Cell num: %d\tUrl: %s\n", rowNumber+1, itemUrl)
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
