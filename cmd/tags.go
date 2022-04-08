package cmd

import (
	"fmt"
	"log"

	"github.com/m7shapan/njson"
	"github.com/spf13/cobra"
)

// tagsCmd represents the tags command
var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Get tags list",
	Long: `This command fetch list of quote tags from Go Quotes API`,
	Run: func(cmd *cobra.Command, args []string) {
		
		url := "https://goquotes-api.herokuapp.com/api/v1/all/tags"
		responseBytes := getData(url)
		tags := Tags{}
		if err := njson.Unmarshal(responseBytes, &tags); err != nil {
			log.Fatalf("Could not read response - %v", err)
		}

		fmt.Printf("\nTotal Tags = %v\n", tags.Totals)

		count := 1
		for idx, val := range tags.Name {
			fmt.Printf("\n------Tags %v------\n\n",count)
			fmt.Printf("Name: %v\n", val)
			fmt.Printf("Total Qoute: %v\n", tags.TotalQoute[idx])
			count++
		}

		
	},
}

func init() {
	rootCmd.AddCommand(tagsCmd)
}

type Tags struct {
	Totals 		string		`njson:"count"`
	Name 		[]string	`njson:"tags.#.name"`
	TotalQoute 	[]string	`njson:"tags.#.count"`
}
