package cmd

import (
	"fmt"
	"log"

	"github.com/m7shapan/njson"
	"github.com/spf13/cobra"
)

// authorsCmd represents the authors command
var authorsCmd = &cobra.Command{
	Use:   "authors",
	Short: "Get author lists",
	Long: `This command fetch list of quote authors from Go Quotes API`,
	Run: func(cmd *cobra.Command, args []string) {
		
		url := "https://goquotes-api.herokuapp.com/api/v1/all/authors"
		responseBytes := getData(url)
		authors := Authors{}
		if err := njson.Unmarshal(responseBytes, &authors); err != nil {
			log.Fatalf("Could not read response - %v", err)
		}

		fmt.Printf("\nTotal Authors = %v\n", authors.Totals)
		
		count := 1
		for idx, val := range authors.Name {
			fmt.Printf("\n------Authors %v------\n\n",count)
			fmt.Printf("Name: %v\n", val)
			fmt.Printf("Total Qoute: %v\n", authors.TotalQoute[idx])
			count++
		}
	},
}

func init() {
	rootCmd.AddCommand(authorsCmd)
}

type Authors struct {
	Totals 		string		`njson:"count"`
	Name 		[]string	`njson:"authors.#.name"`
	TotalQoute 	[]string	`njson:"authors.#.count"`
}
