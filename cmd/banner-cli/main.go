package main

import (
	"fmt"
	"os"
	"strings"

	banner "github.com/multiverse-os/banner"
)

func main() {
	message := "test"
	if 1 < len(os.Args) {
		message = strings.Join(os.Args[1:], " ")
	}

	fmt.Println("banner-cli")
	fmt.Println("==========")

	banner := banner.New(message).Font("big")
	fmt.Println("banner:\n" + banner.String())
}
