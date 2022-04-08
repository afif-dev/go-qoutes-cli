package cmd

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/m7shapan/njson"
	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get random quote",
	Long: `This command fetch random quote from Go Quotes API`,
	Run: func(cmd *cobra.Command, args []string) {
		countQoute, _ := cmd.Flags().GetString("total")
		saveQoute, _ := cmd.Flags().GetBool("save")
		
		url := "https://goquotes-api.herokuapp.com/api/v1/random?count=" + url.QueryEscape(countQoute)
		responseBytes := getData(url)
		qoute := Qoute{}
		if err := njson.Unmarshal(responseBytes, &qoute); err != nil {
			log.Fatalf("Could not read response - %v", err)
		}
		// fmt.Printf("%+v\n", qoute)
		
		count := 1
		for idx, val := range qoute.Text {
			fmt.Printf("\n------Qoute %v------\n\n",count)
			fmt.Printf("Text: \n%v\n", val)
			fmt.Printf("Author: %v\n", qoute.Author[idx])
			fmt.Printf("Tag: %v\n", qoute.Tag[idx])
			count++
		}
		
		// save qoute in text file
		if saveQoute {
			sv_count := 1
			f, _ := os.Create("random_qoutes.text")
			f.WriteString("###################\nRandom Qoutes\n###################\n\n")
			for idx, val := range qoute.Text {
				f.WriteString("\n------Qoute "+ strconv.Itoa(sv_count) +"------\n\n")
				f.WriteString("Text: \n"+ val +"\n")
				f.WriteString("Author: "+ qoute.Author[idx] +"\n" )
				f.WriteString("Tag: "+ qoute.Tag[idx] +"\n")
				sv_count++
			}
		}
	},

}

func init() {
	rootCmd.AddCommand(randomCmd)
	randomCmd.Flags().StringP("total", "t", "1", "Number of quotes to return")
	randomCmd.Flags().BoolP("save", "s", false, "save quotes in text file")
}

type Qoute struct {
	Text 	[]string	`njson:"quotes.#.text"`
	Author 	[]string	`njson:"quotes.#.author"`
	Tag		[]string	`njson:"quotes.#.tag"`
}
