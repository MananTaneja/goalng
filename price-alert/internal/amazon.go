package priceAlert

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Amazon(url string) {
	fmt.Println("Searching for product on Amazon.in")

	client := &http.Client{Timeout: 30 * time.Second}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Android 12; Mobile; rv:109.0) Gecko/113.0 Firefox/113.0")
	req.Header.Set("Cookie", "cookie_key=cookie_value;")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error occured")
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

}
