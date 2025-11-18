package cmd

import (
	"github.com/dracuxan/msf-skynet/handlers"
	"github.com/spf13/cobra"
)

var sessionListCmd = &cobra.Command{
	Use:   "list",
	Short: "This command will list all ongoing sessions if any",
	Long: `This sessionList command will make a RPC call to MSFConsole to get all the ongoing 
	sessions and display none if there are no ongoing session`,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.GetSessions()
	},
}

func init() {
	sessCmd.AddCommand(sessionListCmd)
}
