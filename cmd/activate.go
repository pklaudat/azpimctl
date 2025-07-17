/*
Copyright © 2025 Paulo Klaudat
*/
package cmd

import (
	"fmt"

	"github.com/pklaudat/azpimctl/pkg/pim"
	"github.com/spf13/cobra"
)

func activateRole(cmd *cobra.Command, args []string) {
	fmt.Printf("Activate the PIM role for")
	pimClient := pim.NewPimClient(pim.PIM_DEFAULT_SCOPE)
	var resourceID string
	if len(args) > 0 {
		resourceID = args[0]
	}
	pimClient.ActivateElegibleRole(resourceID)
}

var activateCmd = &cobra.Command{
	Use:   "up",
	Short: "Activate a PIM role for the given resource context.",
	Long:  `Activate a specific role for the given scope via PIM.`,
	Run:   activateRole,
}

func init() {
	rootCmd.AddCommand(activateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// activateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// activateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
