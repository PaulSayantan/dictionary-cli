//	Author:	Sayantan Paul
package main

import (
	"antonyms"
	"definition"
	"github.com/spf13/cobra"
	"log"
	"os"
	"synonyms"
	"wordoftheday"
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
			err := wordoftheday.WordOfDay()
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
		Use:     "synonyms",
		Short:   "returns synonyms of the given word",
		Example: "dict synonyms [word] \n\tdict synonyms satire",
		RunE: func(cmd *cobra.Command, args []string) error {
			word := args[0]
			err := synonyms.Synonyms(word)
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
		Use:     "antonyms",
		Short:   "returns antonyms of the given word",
		Example: "dict antonyms [word] \n\tdict antonyms sad",
		RunE: func(cmd *cobra.Command, args []string) error {
			word := args[0]
			err := antonyms.Antonyms(word)
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
			err := definition.Definition(word)
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
