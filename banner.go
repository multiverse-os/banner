package banner

type Banner struct {
	text string
	font *Font
}

func New(text string) Banner {
	return Banner{
		text: text,
		font: LoadFont("big"),
	}
}

func (self Banner) Font(name string) Banner {
	self.font = LoadFont(name)
	return self
}

func (self Banner) String() string {
	return self.font.WriteString(self.text)
}
