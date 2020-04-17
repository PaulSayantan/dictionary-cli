//   Author:  Sayantan Paul
package wordoftheday

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

//-------------------    WORD OF THE DAY SECTION    -----------------------//

func WordOfDay() error {

	URL := "https://www.dictionary.com/e/word-of-the-day/"

	resp, err := http.Get(URL)
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

	wod(doc)

	return nil
}

func wod(doc *goquery.Document) {
	var word string
	doc.Find(".wotd-item-wrapper-content").First().Each(
		func(i int, data *goquery.Selection) {
			data.Find(".wotd-item-headword__word h1").EachWithBreak(
				func(i int, selection *goquery.Selection) bool {
					word = selection.Text()
					return false
				})

			fmt.Println("Today's Word Of the Day: ", word, "\n---------------------------------------------")
			data.Find(".wotd-item-headword__pos p").Each(
				func(i int, selection *goquery.Selection) {
					fmt.Println(selection.Next().Text())

				})

		})

}
