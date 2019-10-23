package banner

import (
	"fmt"
	"strings"

	big "github.com/multiverse-os/cli/framework/text/banner/fonts/big"
	giant "github.com/multiverse-os/cli/framework/text/banner/fonts/giant"
)

//type Character struct {
//	Lines  []string
//	Height int
//	Width  int
//}

type Font struct {
	Characters       map[string][]string
	characterSpacing int
	//Characters       map[string]Character
	//paddingLeft      int
	//paddingRight     int
}

func LoadFont(name string) *Font {
	switch strings.ToLower(name) {
	// TODO: Need to do analysis to see if this method will load every condition
	// in the switch. If so, need to hand off this specific step to thing
	// calling lib
	case "giant":
		return &Font{
			characterSpacing: 1,
			Characters:       giant.Characters,
		}
	default:
		return &Font{
			characterSpacing: 1,
			Characters:       big.Characters,
		}
	}
	//for _, character := range big.Characters {
	//	font.Characters = append(font.Characters, Character{
	//		Lines:  character,
	//		Height: len(character),
	//		Width:  len(character[0]),
	//	})
	//}
}

//func (self *Font) PaddingLeft(whitespace int) *Font {
//	self.paddingLeft = whitespace
//	return self
//}
//
//func (self *Font) PaddingRight(whitespace int) *Font {
//	self.paddingRight = whitespace
//	return self
//}

func (self *Font) CharacterSpacing(whitespace int) *Font {
	self.characterSpacing = whitespace
	return self
}

func (self *Font) WriteString(text string) string {
	outputLines := []string{}
	fmt.Println("text:", text)
	for _, character := range text {
		fontCharacter := self.Characters[string(character)]
		for index, line := range fontCharacter {
			if len(outputLines) <= index {
				outputLines = append(outputLines, "")
			}
			outputLines[index] += line + strings.Repeat(" ", self.characterSpacing)
		}
	}
	return strings.Join(outputLines, "\n")
}
