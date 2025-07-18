/*
Copyright © 2025 Paulo Klaudat
*/
package cmd

import (
	"strconv"

	"github.com/pklaudat/azpimctl/pkg/pim"
	"github.com/spf13/cobra"
)

func activateRole(cmd *cobra.Command, args []string) {
	pimClient := pim.NewPimClient(pim.PIM_DEFAULT_SCOPE)
	var resourceID string
	if len(args) > 0 {
		resourceID = args[0]
	}
	var hours int
	if len(args) > 1 {
		// Try to parse the hours argument from args[1]
		var err error
		hours, err = strconv.Atoi(args[1])
		if err != nil {
			hours = 8 // fallback to default if parsing fails
		}
	} else {
		hours = 8 // default to 8 hours if not provided
	}
	pimClient.ActivateElegibleRole(resourceID, "", strconv.Itoa(hours))
}

var activateCmd = &cobra.Command{
	Use:     "activate",
	Short:   "Activate a PIM role for the given resource context.",
	Long:    `Activate roles for the given scope using Privileged Identity Management APIs.`,
	Aliases: []string{"up", "apply", "a"},
	Run:     activateRole,
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
