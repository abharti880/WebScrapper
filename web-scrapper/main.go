package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	fName := "webdata.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("could not create file, err :%q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	collector := colly.NewCollector(
		colly.AllowedDomains("amazon.in"),
	)

	collector.OnHTML("titleSection", func(element *colly.HTMLElement) {
		writer.Write([]string{
			element.ChildText("span"),
		})
	})

	collector.OnHTML("acrCustomerReviewLink", func(element *colly.HTMLElement) {
		writer.Write([]string{
			element.ChildText("span"),
		})
	})

	collector.OnHTML("price", func(element *colly.HTMLElement) {
		writer.Write([]string{
			element.ChildText("priceblock_ourprice"),
		})
	})

	collector.OnHTML(".feature-bullets li", func(element *colly.HTMLElement) {
		writer.Write([]string{
			element.ChildText("ul"),
		})
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting ", request.URL.String())
	})

	collector.Visit("https://www.amazon.in/New-Apple-iPhone-12-128GB/dp/B08L5TNJHG/ref=sr_1_7?dchild=1&keywords=iphone&qid=1622718161&sr=8-7&th=1")

	log.Printf("scrapping complete")
	log.Println(collector)
}
