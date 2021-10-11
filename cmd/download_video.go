package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	downloadVideoCmd = &cobra.Command{
		Use:   "video",
		Short: "A downloader that used to download niconico video.",

		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Hello", Account, Password)
		},
	}
)

func init() {
	rootCmd.AddCommand(downloadVideoCmd)
}
