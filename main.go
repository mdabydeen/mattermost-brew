package main

import (
	"fmt"
	"os"
)

func main() {
	serverURL := os.Getenv("MATTERMOST_SERVER")
	fmt.Print(serverURL)
}
