package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rmTag string

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a server from the registry",
	Run: func(cmd *cobra.Command, args []string) {
		if rmTag == "" {
			fmt.Println("Error: --tag is required")
			os.Exit(1)
		}

		if _, exists := servers[rmTag]; !exists {
			fmt.Printf("Error: No server found with tag %s\n", rmTag)
			os.Exit(1)
		}

		delete(servers, rmTag)
		saveServers()

		fmt.Printf("Server %s removed\n", rmTag)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
	rmCmd.Flags().StringVar(&rmTag, "tag", "", "Tag of the server to remove (required)")
	rmCmd.MarkFlagRequired("tag")
}
