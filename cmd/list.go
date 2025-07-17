/*
Copyright © 2025 Paulo Klaudat
*/
package cmd

import (
	"fmt"

	"github.com/pklaudat/azpimctl/pkg/pim"
	"github.com/spf13/cobra"
)

func listCommand(cmd *cobra.Command, args []string) {
	fmt.Printf("Listing PIM roles for user: \n")
	scope := "abc"
	pimClient := pim.NewPimClient(scope)
	pimClient.ListElegibleRoles()

}

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List PIM role assignments and their scopes - Resource based roles.",
	Long:    `List command retrieves and display all PIM Resource Role Assignments and their scopes for the authenticated user.`,
	Aliases: []string{"ls"},
	Run:     listCommand,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pimCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pimCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
