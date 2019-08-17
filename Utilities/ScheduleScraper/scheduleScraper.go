package scraperutil

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetExams Scrapes the url for exams
func GetExams(url string) []string {
	// Make HTTP request
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	var examsList []string

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find(".commonCell").Each(func(i int, s *goquery.Selection) {
		if len(s.Text()) != 0 && strings.Contains(strings.ToLower(s.Text()), "salstentamen") {
			exams := s.Siblings().Text()
			exams = strings.TrimPrefix(exams, "A")
			examsList = append(examsList, exams)
		}
	})
	if len(examsList) == 0 {
		examsList = []string{}
	}
	return examsList
}
