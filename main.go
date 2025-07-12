package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dracuxan/msf-rpc-client/msf"
	"github.com/joho/godotenv"
)

func main() {
	path := fmt.Sprintf(".env")
	if err := godotenv.Load(path); err != nil {
		log.Fatalf("Unable to load env file: %v", err)
	}

	host := os.Getenv("MSFHOST")
	pass := os.Getenv("MSFPASS")
	user := "msf"

	if host == "" || pass == "" {
		log.Fatalln("Missing required environment variables MSFHOST or MSFPASS")
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

	better_looks, err := json.MarshalIndent(sessions, "", " ")
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("Current Sessions:")
	fmt.Println(string(better_looks))
}
