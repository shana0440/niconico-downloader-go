package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	Account  string
	Password string

	rootCmd = &cobra.Command{
		Use:   "nico",
		Short: "A nico nico downloader",
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&Account, "account", "a", "", "user account")
	rootCmd.PersistentFlags().StringVarP(&Password, "password", "p", "", "account password")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
