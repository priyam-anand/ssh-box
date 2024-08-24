package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var keyPath, serverIP, tag, user string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a server to the registry",
	Run: func(cmd *cobra.Command, args []string) {
		if keyPath == "" || serverIP == "" || tag == "" {
			fmt.Println("Error: --key, --server, and --tag are required")
			os.Exit(1)
		}

		servers[tag] = Server{
			KeyPath: keyPath,
			IP:      serverIP,
			Tag:     tag,
			User:    user,
		}

		saveServers()

		fmt.Printf("Server %s registered with IP %s\n", tag, serverIP)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVar(&keyPath, "key", "", "Path to SSH key (required)")
	addCmd.Flags().StringVar(&serverIP, "server", "", "Server IP address (required)")
	addCmd.Flags().StringVar(&tag, "tag", "", "Tag to identify the server (required)")
	addCmd.Flags().StringVar(&user, "user", "ubuntu", "User to SSH into the server (optional)")
	addCmd.MarkFlagRequired("key")
	addCmd.MarkFlagRequired("server")
	addCmd.MarkFlagRequired("tag")
}
