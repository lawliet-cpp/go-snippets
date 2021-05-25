package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/kevinburke/twilio-go"
	"github.com/spf13/cobra"
)

var URL string
var ssid string = "TWILLO SSID"
var authtoken string = "TWILLO AUTHTOKEN"
var client = twilio.NewClient(ssid, authtoken, nil)

var startCmd = &cobra.Command{
	Use:     "website monitor",
	Short:   "start -u [websiteURL]",
	Example: "start -u [websiteURL]",
	PreRun: func(cmd *cobra.Command, args []string) {
		if (URL) == "" {
			fmt.Println("The URL flag is required")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		for {
			response, err := http.Get(URL)
			if err != nil {
				client.Messages.SendMessage("YOUR TWILLO NUMBER", "+YOUR NUMBER", "Your website is down", nil)

				os.Exit(1)
			}
			if response.StatusCode != 200 {
				fmt.Println("Website Down")
				fmt.Println(response.StatusCode)
				client.Messages.SendMessage("YOUR TWILLO NUMBER", "+YOUR NUMBER", "Your website is down", nil)

			}
			time.Sleep(10 * time.Second)
		}

	},
}

func init() {

	startCmd.Flags().StringVarP(&URL, "u", "u", "", "website URL")
	rootCmd.AddCommand(startCmd)
}
