/*
Copyright © 2022 afif-dev https://github.com/afif-dev

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "qoutes-cli",
	Short: "Get Qoutes by Go Quotes",
	Long: `
################################################################	

░██████╗░░█████╗░██╗░░░██╗████████╗███████╗  ░█████╗░██╗░░░░░██╗
██╔═══██╗██╔══██╗██║░░░██║╚══██╔══╝██╔════╝  ██╔══██╗██║░░░░░██║
██║██╗██║██║░░██║██║░░░██║░░░██║░░░█████╗░░  ██║░░╚═╝██║░░░░░██║
╚██████╔╝██║░░██║██║░░░██║░░░██║░░░██╔══╝░░  ██║░░██╗██║░░░░░██║
░╚═██╔═╝░╚█████╔╝╚██████╔╝░░░██║░░░███████╗  ╚█████╔╝███████╗██║
░░░╚═╝░░░░╚════╝░░╚═════╝░░░░╚═╝░░░╚══════╝  ░╚════╝░╚══════╝╚═╝
----------------------------------------------------------------
Get Qoutes by Go Quotes API
Site: https://goquotes.docs.apiary.io/

Develop by afif-dev https://github.com/afif-dev
################################################################
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.qoutes-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getData(url string) []byte {
	
	res, err := http.Get(url)
	res.Header.Add("Accept", "application/json")
	res.Header.Add("User-Agent", "Qoutes CLI")

	if err != nil {
		log.Println(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatalln(err)
	}
	
	return body
}
