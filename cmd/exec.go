package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var execTag string

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "SSH into a server by tag",
	Run: func(cmd *cobra.Command, args []string) {
		if execTag == "" {
			fmt.Println("Error: --tag is required")
			os.Exit(1)
		}

		server, exists := servers[execTag]
		if !exists {
			fmt.Printf("Error: No server found with tag %s\n", execTag)
			os.Exit(1)
		}

		sshCommand := exec.Command("ssh", "-i", server.KeyPath, fmt.Sprintf("%s@%s", server.User, server.IP))
		sshCommand.Stdout = os.Stdout
		sshCommand.Stderr = os.Stderr
		sshCommand.Stdin = os.Stdin

		if err := sshCommand.Run(); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
	execCmd.Flags().StringVar(&execTag, "tag", "", "Tag of the server to SSH into (required)")
	execCmd.MarkFlagRequired("tag")
}
