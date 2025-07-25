package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dracuxan/msf-skynet/msf"
)

func GetSessions() {
	cfg := msf.Config{}
	cfg.New()

	msf, err := msf.New(cfg.MsfHost, cfg.MsfUser, cfg.MsfPass)
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
