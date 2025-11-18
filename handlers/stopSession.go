package handlers

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dracuxan/msf-skynet/msf"
)

func StopSession(sessionID string) {
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

	id, err := parseSessionID(sessionID)
	if err != nil {
		log.Fatalf("Failed to stop session %s: %v", sessionID, err)
	}

	if _, exists := sessions[id]; !exists {
		fmt.Printf("Session %s does not exist.\n", sessionID)
		os.Exit(1)
	}

	if ok, err := msf.SessionStop(sessionID); !ok || err != nil {
		log.Fatalf("Failed to stop session %s: %v", sessionID, err)
	}

	fmt.Printf("Session %s stopped successfully.\n", sessionID)
}

func parseSessionID(s string) (uint32, error) {
	id, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(id), nil
}
