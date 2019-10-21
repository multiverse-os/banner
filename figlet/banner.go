package figletbanner

import (
	"strings"
)

// TODO: A way to import new font files will be useful

const defaultFont = "standard"

type Banner struct {
	Text   string
	Lines  []string
	Height int
	Width  int
}

func (self Banner) String() (output string) {
	for _, line := range self.Lines {
		output += line + "\n"
	}
	return output
}

func New(fontName, text string) Banner {
	font := loadFont(fontName)
	banner := Banner{
		Text: text,
	}
	for line := 0; line < font.height; line++ {
		var lineText string
		for _, textCharacter := range banner.Text {
			// TODO: Is this just checking if its ascii? Because we could just use the
			// standard library that supplies that functionality
			if byte(textCharacter) < ' ' || byte(textCharacter) > '~' {
				textCharacter = '?'
			}
			i := textCharacter - 32
			lineText += strings.Replace(font.letters[i][line], string(font.hardblank), " ", -1)
		}
		// TODO: This is a bit messy, should fix this
		if line < font.baseline || len(strings.TrimSpace(lineText)) > 0 {
			if !(line == 0 && len(strings.TrimSpace(lineText)) == 0) {
				banner.Lines = append(banner.Lines, lineText)
			}
		}
	}
	if len(banner.Lines) != 0 {
		banner.Width = len(banner.Lines[0])
	} else {
		banner.Width = 0
	}
	return banner
}
