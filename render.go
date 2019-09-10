package banner

import (
	"bufio"
	"bytes"
	"io"
	"path"
	"strconv"
	"strings"
)

var magicSequence = "flf2"
var fontExtension = ".flf"
var charDelimiters = [3]string{"@", "#", "$"}
var hardblanksBlacklist = [2]byte{'a', '2'}

type font struct {
	name      string
	height    int
	baseline  int
	hardblank byte
	letters   [][]string
}

func newFont(name string) (font font) {
	font.name = name
	if len(name) < 1 {
		font.name = defaultFont
	}
	fontBytes, err := Asset(path.Join("fonts", font.name+fontExtension))
	if err != nil {
		panic(err)
	}
	fontBytesReader := bytes.NewReader(fontBytes)
	scanner := bufio.NewScanner(fontBytesReader)
	font.setAttributes(scanner)
	font.setLetters(scanner)
	return font
}

func newFontFromReader(reader io.Reader) (font font) {
	scanner := bufio.NewScanner(reader)
	font.setAttributes(scanner)
	font.setLetters(scanner)
	return font
}

func (self *font) setAttributes(scanner *bufio.Scanner) {
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, magicSequence) {
			self.height = getHeight(text)
			fontData := strings.Fields(text)[2]
			self.baseline, _ = strconv.Atoi(fontData)
			self.hardblank = getHardblank(text)
			break
		}
	}
}

func (self *font) setLetters(scanner *bufio.Scanner) {
	self.letters = append(self.letters, make([]string, self.height, self.height)) //TODO: set spaces from flf
	for i := range self.letters[0] {
		self.letters[0][i] = "  "
	} //TODO: set spaces from flf
	letterIndex := 0
	for scanner.Scan() {
		text, cutLength, letterIndexInc := scanner.Text(), 1, 0
		if lastCharLine(text, self.height) {
			self.letters = append(self.letters, []string{})
			letterIndexInc = 1
			if self.height > 1 {
				cutLength = 2
			}
		}
		if letterIndex > 0 {
			appendText := ""
			if len(text) > 1 {
				appendText = text[:len(text)-cutLength]
			}
			self.letters[letterIndex] = append(self.letters[letterIndex], appendText)
		}
		letterIndex += letterIndexInc
	}
}

func getHeight(text string) int {
	fontData := strings.Fields(text)[1]
	height, _ := strconv.Atoi(fontData)
	return height
}

// TODO: Thse function names are terrible, they are not decipherable for people
// not very aquainted with this development
func getHardblank(text string) byte {
	fontData := strings.Fields(text)[0]
	hardblank := fontData[len(fontData)-1]
	if hardblank == hardblanksBlacklist[0] || hardblank == hardblanksBlacklist[1] {
		return ' '
	} else {
		return hardblank
	}
}

func lastCharLine(text string, height int) bool {
	endOfLine, length := "  ", 2
	if height == 1 && len(text) > 0 {
		length = 1
	}
	if len(text) >= length {
		endOfLine = text[len(text)-length:]
	}
	return endOfLine == strings.Repeat(charDelimiters[0], length) ||
		endOfLine == strings.Repeat(charDelimiters[1], length) ||
		endOfLine == strings.Repeat(charDelimiters[2], length)
}
