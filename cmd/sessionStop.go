package cmd

import (
	"github.com/dracuxan/msf-skynet/handlers"
	"github.com/spf13/cobra"
)

var sessionStopCmd = &cobra.Command{
	Use:   "kill [session-id]",
	Short: "Stop a Metasploit session",
	Long:  `Stops an active Metasploit session by its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		handlers.StopSession(id)
	},
}

func init() {
	sessCmd.AddCommand(sessionStopCmd)
}
