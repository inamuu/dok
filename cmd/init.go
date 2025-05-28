package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Dok configuration",
	Long: `The init command sets up the Dok configuration files and directories.
It creates a default configuration file and necessary directories for Dok to function properly.`,
	Run: func(cmd *cobra.Command, args []string) {
		home := os.Getenv("HOME")
		configFile := filepath.Join(home, ".dok.config")

		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			content := `[Commands]
/bin/sh
/bin/bash
ls -la
ps
`
			os.WriteFile(configFile, []byte(content), 0644)
			fmt.Println("Initialized configuration file at", configFile)
		} else {
			fmt.Printf("Configuration file already exists at %s\n", configFile)
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
