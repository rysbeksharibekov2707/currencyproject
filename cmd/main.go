package main

import (
	"currencyrate/internal/fetcher"
	"currencyrate/internal/parser"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res := fetcher.GetHTML("https://mig.kz/") // поставь сюда реальный URL

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	rates := parser.ParseRates(doc)

	for _, r := range rates {
		fmt.Printf("%s: Покупка — %s | Продажа — %s (%s)\n", r.Name, r.Buy, r.Sell, r.Link)
	}
}
