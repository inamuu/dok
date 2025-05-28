package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dok",
	Short: "Dok is a tool for managing Docker containers",
	Long: `Dok is a command-line tool that simplifies the management of Docker containers.
It provides a set of commands to create, start, stop, and remove containers easily.`,
	Run: func(cmd *cobra.Command, args []string) {
		mainMenu()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
