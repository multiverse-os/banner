package main

import (
	"fmt"
	"os"

	color "github.com/multiverse-os/cli/text/ansi/color"
	banner "github.com/multiverse-os/cli/text/ascii/banner"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(color.Silver("[error] failed to provide arguments, try again using the following:\n"))
		fmt.Println(color.White("  banner-cli [font] [text]\n"))
	} else {
		fmt.Println(color.Silver("banner-cli: ASCII banner example"))
		fmt.Println(color.Gray("================================"))
		fmt.Println(banner.New(args[1], args[2]).String())
	}
}
