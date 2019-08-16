package quotes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	utilities "github.com/femosc2/ia-discord-bot-2/Utilities/RandomGenerators"
	config "github.com/femosc2/ia-discord-bot-2/config"
	viewmodels "github.com/femosc2/ia-discord-bot-2/viewmodels"
)

// GetQuote Get a random quote
func GetQuote() string {
	quotes := config.Firebase + "/Quotes.json"
	return getJSON(quotes)
}

func getJSON(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("response Error")
		// return err
	}
	defer response.Body.Close()
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		fmt.Println("read Error")
		fmt.Println(readErr)
	}
	quoteList := []viewmodels.QuoteViewModel{}
	jsonErr := json.Unmarshal(body, &quoteList)

	if jsonErr != nil {
		fmt.Println("json Error")
		fmt.Println(jsonErr)
	}
	// fmt.Println(string(quoteList.Quote.Text + " -" + quoteList.Quote.Who))
	randInt := utilities.RandIntFromRange(len(quoteList))

	return quoteList[randInt].Text + " -" + quoteList[randInt].Who
}
