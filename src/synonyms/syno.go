//   Author:  Sayantan Paul
package synonyms

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

//-------------------    SYNONYMS SECTION    -----------------------//

func Synonyms(word string) error {

	URL := "https://www.collinsdictionary.com/dictionary/english-thesaurus/"

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

	syno(doc)

	return nil
}

func syno(doc *goquery.Document) {
	fmt.Println("-----------SYNONYMS--------------")
	doc.Find(".cdet .form.type-syn .orth").Each(
		func(i int, selection *goquery.Selection) {
			fmt.Printf("\n%d : %s", i+1, selection.Text())
		})
	if doc.Find("div").HasClass("cB") {
		doc.Find(".cB").Find("h1").EachWithBreak(
			func(i int, selection *goquery.Selection) bool {
				fmt.Println(selection.Text())
				return false
			})

		doc.Find(".suggested_words").Each(
			func(i int, selection *goquery.Selection) {
				fmt.Println(selection.Text())
			})
	} else {
		fmt.Println("-----------SYNONYMS--------------")
		doc.Find(".cdet .form.type-syn .orth").Each(
			func(i int, selection *goquery.Selection) {
				fmt.Printf("\n%d : %s", i+1, selection.Text())
			})
	}
}
