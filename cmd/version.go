package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Dok",
	Long:  `All software has versions. This is Dok's version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Dok version %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
