package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "http-debug",
	Short: "Debug your http calls",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		serve(":8080")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
