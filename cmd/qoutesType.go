package cmd

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/m7shapan/njson"
	"github.com/spf13/cobra"
)

// qoutesTypeCmd represents the qoutesType command
var qoutesTypeCmd = &cobra.Command{
	Use:   "qoutesType",
	Short: "Get quotes by type of author or tag",
	Long: `This command fetch quotes by type of author or tag from Go Quotes API`,
	Run: func(cmd *cobra.Command, args []string) {
		qouteType, _ := cmd.Flags().GetString("type")
		qouteTypeVal, _ := cmd.Flags().GetString("value")
		saveQoute, _ := cmd.Flags().GetBool("save")

		url := "https://goquotes-api.herokuapp.com/api/v1/all?type=" + url.QueryEscape(qouteType) + "&val=" + url.QueryEscape(qouteTypeVal)

		responseBytes := getData(url)
		qouteByType := QouteByType{}
		if err := njson.Unmarshal(responseBytes, &qouteByType); err != nil {
			log.Fatalf("Could not read response - %v", err)
		}

		fmt.Printf("\nTotal Qoutes = %v\n", qouteByType.Totals)
		fmt.Printf("Qoute Type = %v\n", strings.Title(qouteType))
		fmt.Printf("Qoute Value = %v\n", qouteTypeVal)

		count := 1
		for idx, val := range qouteByType.Text {
			fmt.Printf("\n------Qoute %v------\n\n",count)
			fmt.Printf("Text: \n%v\n", val)
			fmt.Printf("Author: %v\n", qouteByType.Author[idx])
			fmt.Printf("Tag: %v\n", qouteByType.Tag[idx])
			count++
		}

		// save qoute in text file
		if saveQoute {
			sv_count := 1
			f, _ := os.Create("random_qoutes-"+qouteType+"-"+qouteTypeVal+".text")
			f.WriteString("###################\nQoutes by Type\n###################\n\n")
			f.WriteString("Total Qoutes = "+ qouteByType.Totals +"\n")
			f.WriteString("Qoute Type = "+ strings.Title(qouteType) +"\n")
			f.WriteString("Qoute Value = "+ qouteTypeVal +"\n")
			for idx, val := range qouteByType.Text {
				f.WriteString("\n------Qoute "+ strconv.Itoa(sv_count) +"------\n\n")
				f.WriteString("Text: \n"+ val +"\n")
				f.WriteString("Author: "+ qouteByType.Author[idx] +"\n" )
				f.WriteString("Tag: "+ qouteByType.Tag[idx] +"\n")
				sv_count++
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(qoutesTypeCmd)
	qoutesTypeCmd.Flags().StringP("type", "t", "tag", "Filter quotes by type (author|tag)")
	qoutesTypeCmd.Flags().StringP("value", "v", "motivational", "Filter quotes by value of type")
	qoutesTypeCmd.Flags().BoolP("save", "s", false, "save quotes in text file")
}

type QouteByType struct {
	Totals 	string		`njson:"count"`
	Text 	[]string	`njson:"quotes.#.text"`
	Author 	[]string	`njson:"quotes.#.author"`
	Tag		[]string	`njson:"quotes.#.tag"`
}
