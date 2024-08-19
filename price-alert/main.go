package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	priceAlert "priceAlert/internal"
)

type RequestBody struct {
	PageUrl string `json:"pageUrl"`
}

type Response struct {
	Price string `json:"price"`
}

func amazonHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("inside the amazon api handler")

		var body RequestBody

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			fmt.Println("unable to parse body")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		price, err := priceAlert.Amazon(body.PageUrl)

		if err != nil {
			fmt.Println("Error while parsing price from the given url")
		}

		w.Header().Set("Content-Type", "application/json")

		response := Response{
			Price: price,
		}

		json.NewEncoder(w).Encode(response)

	}

}

func main() {
	fmt.Println("Running")

	// priceAlert.Amazon("https://www.amazon.in/Stellar-Sports-Taotei-Helmet-Certified/dp/B0D2TN3MCY/ref=sr_1_7?crid=XN6SC58V6AKL&dib=eyJ2IjoiMSJ9.UY8E04-AxmyzSrxh9bLemRc3xTAQ2SXQ5-IeRVVg0UmyR0ZalntNNyarENdnA5jXYmEEQHYseG7RRiFGOFmcfhS0HEtiudrLSzoSPbBZd46LoS2UcHxAcOK8kSlm1JZ7Nwhh-8h-BL52T3JWH2uEdoCjHlAaaQQK1gGMcv-6p1Dt6PvOvakhHzFTs3rN_jIqX1RQ-Aq1oEmmfi7g7UHkkgBhxV8LztgfE_Yyoiquko9yt-fAKgTjLG9QNvLOHy9dXD8P0lzNO2zUvDDRqMPXL9DSCi3QC9O_LDCz4x3TIlM.YYtp0yEA052RL8iLGlYF8u9oJPj5Yx1TadIcRRVqCxY&dib_tag=se&keywords=smk%2Bhelmets&qid=1723396933&sprefix=smk%2Bhelmen%2Caps%2C232&sr=8-7&th=1")

	// priceAlert.Amazon("https://www.flipkart.com/smk-stellar-motorbike-helmet/p/itme3f2cf69da28f?pid=HLMGZNDN39JBW9FS&lid=LSTHLMGZNDN39JBW9FSZ14XK9&marketplace=FLIPKART&q=smk+helmet&store=1mt%2Fztf%2Fiv8&srno=s_1_7&otracker=search&otracker1=search&fm=organic&iid=df0c8f6f-1d39-4641-a93e-c30f4816a843.HLMGZNDN39JBW9FS.SEARCH&ppt=hp&ppn=homepage&ssid=xxunzh68xyubn5s01723397493020&qH=553fa1d8a9ce7b60")

	http.HandleFunc("/api/v1/amazon", amazonHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))

}
