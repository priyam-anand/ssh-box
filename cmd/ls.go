package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var all bool

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all registered servers",
	Run: func(cmd *cobra.Command, args []string) {
		if len(servers) == 0 {
			fmt.Println("No servers registered")
			return
		}

		for tag, server := range servers {
			if all {
				fmt.Printf("Tag: %s, IP: %s, User: %s\n", tag, server.IP, server.User)
			} else {
				fmt.Println(tag)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().BoolVarP(&all, "all", "a", false, "Show detailed server info")
}
