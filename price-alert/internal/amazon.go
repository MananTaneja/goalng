package priceAlert

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func Amazon(pageUrl string) (string, error) {
	fmt.Println("Searching for product on Amazon.in")

	client := &http.Client{Timeout: 30 * time.Second}

	req, _ := http.NewRequest("GET", pageUrl, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Android 12; Mobile; rv:109.0) Gecko/113.0 Firefox/113.0")
	req.Header.Set("Cookie", "cookie_key=cookie_value;")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error occured")
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		fmt.Println("Error occured while parsing with goquery")
	}

	var price string

	// TODO: make a class for all variations and read this from a config instead
	doc.Find(".reinventPricePriceToPayMargin > span .a-price-whole").Each(func(i int, s *goquery.Selection) {
		possiblePrice := s.Text()

		price = possiblePrice
		fmt.Println(possiblePrice)
	})

	if price == "" {
		fmt.Println("Unable to fetch price")
		html, err := doc.Html()

		if err != nil {
			fmt.Println("Error parsing html from document")
		}

		err = os.WriteFile("hello.html", []byte(html), 0644)

		if err != nil {
			fmt.Println("Error writing file")
		}
	}

	return price, err

}
