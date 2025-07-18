/*
Copyright © 2025 Paulo Klaudat
*/
package main

import (
	"fmt"
	"os"

	"github.com/pklaudat/azpimctl/cmd"
	"github.com/pklaudat/azpimctl/pkg/auth"
)

func main() {
	ctx, err := auth.GetContext()
	fmt.Printf("==> Default Subscription: %s\n", ctx.SubscriptionName)
	if err != nil {
		fmt.Printf("Failed to fetch context - login using Azure CLI first and choose a subscription.")
		os.Exit(1)
	}
	cmd.Execute()
}
