package main

import (
	"fmt"

	banner "github.com/multiverse-os/banner"
)

func main() {
	fmt.Println("banner-cli")
	fmt.Println("==========")

	banner := banner.New("test").Font("big")
	fmt.Println("banner:\n" + banner.String())
}
