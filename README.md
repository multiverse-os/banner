<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">

## Multiverse OS: ASCII `banner` library
**URL** [multiverse-os.org](https://multiverse-os.org)

A simple library to use embedded figlet fonts to generate ASCII banners for
improved user interfaces.

### Example
A CLI program is included as an example and provide ASCII fonts from terminal.

``` go
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
```

Resulting in the following output:

```
banner-cli: ASCII banner example
================================
  _                  _   
 | |_    ___   ___  | |_ 
 | __|  / _ \ / __| | __|
 | |_  |  __/ \__ \ | |_ 
  \__|  \___| |___/  \__|
```
