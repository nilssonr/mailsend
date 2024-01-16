package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mailsend",
	Short: "Simple email sender",
	Long:  "mailsend - Simple email sender written by Robin Nilsson (robin.ingemar.nilsson@gmail.com)",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
