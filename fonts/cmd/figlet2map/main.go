package main

import (
	"fmt"
)

// TODO: Write a simple script to take a figlet font, and embed the characters
// directly into a go map. This will enable us to have access to all the
// existing fonts created without requiring extra files, enable them to be
// imported and used in a variety of ways, but we will support each font being a
// separate package so we can limit dramatically how much we memory is required
// to get a very basic font to print a banner.
const (
	magicSequence = "flf2"
	fontExtension = ".flf"
)

var (
	delimiters          = [3]string{"@", "#", "$"}
	hardblanksBlacklist = [2]byte{'a', '2'}
)

type Font struct {
	Name         string
	Filename     string
	Delimter     string
	CharacterMap map[string]string
}

func main() {
	fmt.Println("figlet2go")
	fmt.Println("=========")

	// TODO:
	// 1) Expect an argument, if missing output helptext.
	// 2) Ensure file exists, output error if fails os.Stat
	// 3) parse each character and load into map. verify data as its being added
	// 4) Confirm we can determine the delimeter. Even if its not directly
	// parseable, it could be calculated by reading each line of the file and
	// seeing which character is across the most lines (or something very close to
	// this)
	// line by line by ensuring ascii only and not a delimeter
}

func (self *Font) GoSource() (output string) {
	output += `package ` + fontName + `\n\n`
	output += `var CharacterMap = map[string]string{`
	for character, ascii := range self.CharacterMap {
		output += `"` + character + `":`
		otuput += "`" + ascii + "`,"
	}
	output += `}`
	return output
}
