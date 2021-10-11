package cmd

import (
	"log"

	"github.com/shana0440/niconico-downloader-go/pkg/auth"
	"github.com/spf13/cobra"
)

var (
	downloadVideoCmd = &cobra.Command{
		Use:   "video",
		Short: "A downloader that used to download niconico video.",

		Run: func(cmd *cobra.Command, args []string) {
			session := auth.Login(Account, Password)
			log.Println(session)
		},
	}
)

func init() {
	rootCmd.AddCommand(downloadVideoCmd)
}
