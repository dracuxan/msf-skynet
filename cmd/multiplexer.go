package cmd

import (
	"log"
	"strings"

	"github.com/dracuxan/msf-skynet/msf"
	"github.com/spf13/cobra"
)

var multiplexerCmd = &cobra.Command{
	Use: "multiplexer",

	Short: "Run a reverse proxy to route HTTP requests to Metasploit listeners based on Host headers.",

	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			log.Fatalf("Error reading port flag: %v", err)
		}
		mappings, err := cmd.Flags().GetStringSlice("mapping")
		if err != nil {
			log.Fatalf("Error reading mapping flag: %v", err)
		}

		// Convert mappings to a map
		hostProxy := make(map[string]string)
		for _, m := range mappings {
			parts := strings.SplitN(m, "=", 2)
			if len(parts) != 2 {
				log.Fatalf("Invalid mapping format: %s. Expected format: hostname=url", m)
			}
			hostProxy[parts[0]] = parts[1]
		}

		msf.RunMultiplexer(port, hostProxy)
	},
}

func init() {
	rootCmd.AddCommand(multiplexerCmd)
	multiplexerCmd.Flags().Int("port", 8080, "Port for the reverse proxy to listen on")
	multiplexerCmd.Flags().StringSlice("mapping", []string{}, "Hostname to URL mappings (e.g., attacker1.com=http://10.0.1.20:10080)")
}
