package parser

import (
	"currencyrate/internal/model"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseRates(doc *goquery.Document) []model.CurrencyRate {
	var rates []model.CurrencyRate

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		currency := strings.TrimSpace(s.Find("td").Eq(0).Text())
		buy := strings.TrimSpace(s.Find("td").Eq(1).Text())
		sell := strings.TrimSpace(s.Find("td").Eq(2).Text())
		link, _ := s.Find("a").Attr("href")

		if currency != "" {
			rates = append(rates, model.CurrencyRate{
				Name: currency,
				Buy:  buy,
				Sell: sell,
				Link: link,
			})
		}
	})

	return rates
}
