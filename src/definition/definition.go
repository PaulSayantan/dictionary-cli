//   Author:  Sayantan Paul
package definition

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

//-------------------    DEFINITIONS SECTION    -----------------------//

func Definition(word string) error {

	if word == "" {
		return errors.New("You didn't type any word. Please type any.")
	}

	URL := "https://www.dictionary.com/browse/"

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

	def(doc, word)
	return nil
}

func def(doc *goquery.Document, word string) {

	if doc.Find("h2").HasClass("spell-suggestions-subtitle css-6gthty e19m0k9k0") {
		doc.Find(".css-1w0dr93").Each(
			func(i int, selection *goquery.Selection) {
				fmt.Println(selection.Text())
			})
		doc.Find(".css-6gthty").EachWithBreak(
			func(i int, selection *goquery.Selection) bool {
				fmt.Println(selection.Text())
				return false
			})
		doc.Find(".css-wms8ca").EachWithBreak(
			func(i int, selection *goquery.Selection) bool {
				fmt.Println(selection.Text(), "\n")
				return false
			})
		doc.Find(".css-ohz4fb").Each(
			func(i int, selection *goquery.Selection) {
				fmt.Println(selection.Text())
			})
	} else {
		fmt.Println("Definitions of ", word, "\n---------------------------------------------\n")
		doc.Find(".css-1p89gle").Each(
			func(i int, def *goquery.Selection) {
				example := def.Find(".luna-example.italic").Text()
				defs := def.Text()
				defs = strings.Replace(defs, example, "", -1)
				if example == "" {
					fmt.Println("Definition :: ", defs, "\n\n-----")
				} else {
					fmt.Println("Definition :: ", defs, "\t\tExample :: ", example, "\n\n-----")
				}

			})
	}

}
