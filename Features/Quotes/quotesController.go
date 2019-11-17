package quotes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	httputils "github.com/femosc2/ia-discord-bot-2/Utilities/Http"
	randomgensutils "github.com/femosc2/ia-discord-bot-2/Utilities/RandomGenerators"
	config "github.com/femosc2/ia-discord-bot-2/config"
	viewmodels "github.com/femosc2/ia-discord-bot-2/viewmodels"
)

func getQuotesList() []byte {
	response, err := http.Get(config.Firebase + "/Quotes.json")
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
	return body
}

// GetQuote Get a random quote
func GetQuote() string {

	quoteList := []viewmodels.QuoteViewModel{}
	jsonErr := json.Unmarshal([]byte(getQuotesList()), &quoteList)

	if jsonErr != nil {
		fmt.Println("json Error")
		fmt.Println(jsonErr)
	}
	// fmt.Println(string(quoteList.Quote.Text + " -" + quoteList.Quote.Who))
	randInt := randomgensutils.RandIntFromRange(len(quoteList))

	return quoteList[randInt].Text + " -" + quoteList[randInt].Who
}

// PostQuote Allows a user to post a new quote
func PostQuote(Text string, Who string, Reporter string) string {
	requestBody, requestErr := json.Marshal(viewmodels.QuoteViewModel{
		Text,
		Who,
		Reporter,
	})
	if requestErr != nil {
		log.Fatalln(requestErr)
	}

	quoteList := []viewmodels.QuoteViewModel{}
	jsonErr := json.Unmarshal(getQuotesList(), &quoteList)
	if jsonErr != nil {
		fmt.Println("json Error")
		fmt.Println(jsonErr)
	}
	httputils.PutRequest(config.Firebase+"/Quotes/"+strconv.Itoa((len(quoteList)))+".json", bytes.NewBuffer(requestBody))

	return "Citatet finns nu i databasen! Tack för ditt bidrag!"
}
