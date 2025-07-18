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
	pimClient := pim.NewPimClient(pim.PIM_DEFAULT_SCOPE)
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
	listCmd.Flags().StringP("subscription", "s", "", "The subscription display name or id")
	listCmd.Flags().StringP("role", "r", "", "The role display name to filter")

}
