package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dracuxan/msf-skynet/msf"
	"github.com/joho/godotenv"
)

func GetSessions() {
	if err := godotenv.Load(msf.Path); err != nil {
		log.Fatalf("Unable to load config file: %v", err)
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
}
