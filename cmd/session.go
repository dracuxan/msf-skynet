package cmd

import "github.com/spf13/cobra"

var sessCmd = &cobra.Command{
	Use:   "session",
	Short: "Session shortcuts",
}

func init() { rootCmd.AddCommand(sessCmd) }
