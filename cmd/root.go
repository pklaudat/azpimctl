/*
Copyright © 2025 Paulo Klaudat
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "azpimctl",
	Short: "AzPIMCtl - Azure PIM Command Line Tool",
	Long:  `AzPIMCtl is a command-line tool for managing Azure Privileged Identity Management (PIM) roles.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

var completionCmd = &cobra.Command{
	Use:    "completion",
	Hidden: true,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.Printf("")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(completionCmd)

}
