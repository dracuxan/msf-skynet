/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dracuxan/msf-rpc-client/msf"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// sessionListCmd represents the sessionList command
var sessionListCmd = &cobra.Command{
	Use:   "sessionList",
	Short: "This command will list all ongoing sessions if any",
	Long: `This sessionList command will make a RPC call to MSFConsole to get all the ongoing 
	sessions and display none if there are no ongoing session`,
	Run: func(cmd *cobra.Command, args []string) {
		path := fmt.Sprintf(".env")
		if err := godotenv.Load(path); err != nil {
			log.Fatalf("Unable to load env file: %v", err)
		}

		host := os.Getenv("MSFHOST")
		pass := os.Getenv("MSFPASS")
		user := "msf"

		if host == "" || pass == "" {
			log.Fatalln("Missing env variables MSFHOST or MSFPASS")
		}

		msf, err := msf.New(host, user, pass)
		if err != nil {
			log.Panicln(err)
		}
		defer msf.Logout()

		sessions, err := msf.SessionList()
		if err != nil {
			log.Panicln(err)
		}

		if len(sessions) == 0 {
			fmt.Println("No ongoing sessions.")
			os.Exit(0)
		}

		organizedList, err := json.MarshalIndent(sessions, "", " ")
		if err != nil {
			log.Panicln(err)
		}
		fmt.Println("Current Sessions:")
		fmt.Println(string(organizedList))
	},
}

func init() {
	rootCmd.AddCommand(sessionListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sessionListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sessionListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
