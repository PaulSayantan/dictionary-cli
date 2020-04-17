//   Author:  Sayantan Paul
package antonyms

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

//-------------------    ANTONYMS SECTION    -----------------------//

func Antonyms(word string) error {

	URL := "https://www.synonyms.com/synonyms/"

	resp, err := http.Get(URL + word)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	anto(doc)

	return nil
}

func anto(doc *goquery.Document) {

	if doc.Find("div").HasClass("not-found-message") {
		doc.Find(".not-found-message").Find("h2").EachWithBreak(
			func(i int, selection *goquery.Selection) bool {
				fmt.Println(selection.Text())
				return false
			})
		fmt.Println("You have entered a wrong word.")
	} else {
		fmt.Println("--------------ANTONYMS---------------")
		doc.Find(".antonyms-page .card.type-antonyms .card-content .chip, " +
			".synonyms-page .card.type-antonyms .card-content .chip").Each(
			func(i int, selection *goquery.Selection) {
				fmt.Println(selection.Text())
			})
	}
}
