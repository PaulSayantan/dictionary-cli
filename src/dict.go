package main

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	cmd := &cobra.Command{
		Use:   "dict",
		Short: "a simple cli dictionary",
		Long: "------------------------------------------------------------------------------------\n" +
			"A cli monolingual dictionary. \n" +
			"This dictionary helps to get \n" +
			"1.)	the definition of any word or phrase. \n" +
			"2.)	get to know a new word everyday. \n" +
			"3.)	get synonyms and antonyms of any word. \n\n" +
			"This is a cli-based dictionary made using GoLang.\n" +
			"Please ensure that you have internet connection enabled in your system before running the cli.\n" +
			"Also put the .exe filepath in the system environmental variables\n" +
			"------------------------------------------------------------------------------------\n",
		SilenceUsage: true,
	}
	cmd.AddCommand(scraper())
	cmd.AddCommand(wordofday())
	cmd.AddCommand(synonym())
	cmd.AddCommand(antonym())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

//-------------------    SECTION CMD CALLERS    -----------------------//

func wordofday() *cobra.Command {
	return &cobra.Command{
		Use:     "wod",
		Short:   "Today's Word Of The Day",
		Example: "dict wod",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := wordOfDay()
			if err != nil {
				log.Fatal(err)
			}
			//s := "Thank you"
			//speak := exec.Command("espeak", s)
			//if err := speak.Run(); err != nil {
			//	log.Fatal(err)
			//}
			return nil
		},
	}
}

func synonym() *cobra.Command {
	return &cobra.Command{
		Use:     "synonym",
		Short:   "returns synonyms of the given word",
		Example: "dict synonym [word] \n\tdict synonym satire",
		RunE: func(cmd *cobra.Command, args []string) error {
			word := args[0]
			err := synonyms(word)
			if err != nil {
				log.Fatal(err)
			}
			//s := "Search Complete. Thank you"
			//speak := exec.Command("espeak", s)
			//if err := speak.Run(); err != nil {
			//	log.Fatal(err)
			//}
			return nil
		},
	}
}

func antonym() *cobra.Command {
	return &cobra.Command{
		Use:     "antonym",
		Short:   "returns antonyms of the given word",
		Example: "dict antonym [word] \n\tdict antonym sad",
		RunE: func(cmd *cobra.Command, args []string) error {
			word := args[0]
			err := antonyms(word)
			if err != nil {
				log.Fatal(err)
			}
			//s := "Search Complete. Thank you"
			//speak := exec.Command("espeak", s)
			//if err := speak.Run(); err != nil {
			//	log.Fatal(err)
			//}
			return nil
		},
	}
}

func scraper() *cobra.Command {
	return &cobra.Command{
		Use: "search",
		Short: "returns every definitions of a word or phrase.\n" +
			"\t      When entering a phrase please provide hyphen(-) in between.\n",
		Example: "dict search [word]/[phrase-with-hyphen]\n\tdict search fall\tdict search fallen-timbers",
		RunE: func(cmd *cobra.Command, args []string) error {
			word := args[0]
			err := definition(word)
			if err != nil {
				log.Fatal(err)
			}
			//s := "Search Complete. Thank you"
			//speak := exec.Command("espeak", s)
			//if err := speak.Run(); err != nil {
			//	log.Fatal(err)
			//}
			return nil
		},
	}
}

//-------------------    WORD OF THE DAY SECTION    -----------------------//

func wordOfDay() error {

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

//-------------------    SYNONYMS SECTION    -----------------------//

func synonyms(word string) error {

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

//-------------------    DEFINITIONS SECTION    -----------------------//

func definition(word string) error {

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

	if doc.Find("div").HasClass("css-1urpfgu e16867sm0") {
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

//-------------------    ANTONYMS SECTION    -----------------------//

func antonyms(word string) error {

	URL := "https://www.synonym.com/synonyms/"

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
		doc.Find(".antonym-page .card.type-antonym .card-content .chip, " +
			".synonym-page .card.type-antonym .card-content .chip").Each(
			func(i int, selection *goquery.Selection) {
				fmt.Println(selection.Text())
			})
	}
}
