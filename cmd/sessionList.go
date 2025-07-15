/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/dracuxan/msf-skynet/handlers"
	"github.com/spf13/cobra"
)

// sessionListCmd represents the sessionList command
var sessionListCmd = &cobra.Command{
	Use:   "sessionList",
	Short: "This command will list all ongoing sessions if any",
	Long: `This sessionList command will make a RPC call to MSFConsole to get all the ongoing 
	sessions and display none if there are no ongoing session`,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.GetSessions()
	},
}

func init() {
	rootCmd.AddCommand(sessionListCmd)
}
