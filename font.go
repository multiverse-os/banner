package banner

import (
	"strings"

	big "github.com/multiverse-os/cli/framework/text/banner/fonts/big"
	giant "github.com/multiverse-os/cli/framework/text/banner/fonts/giant"
)

type Character struct {
	Lines []string
	//Height int
	//Width int
}

type Font struct {
	Characters map[string]Character
}

func LoadFont(name string) *Font {
	var characters []string
	switch strings.ToLower(name) {
	case "giant":
		characters
	default:
	}
	for _, character := range big.Characters {

	}
}

func (self *Font) CharacterLines(character string) []string {
	return strings.Split(self.Characters[character], "\n")
}
