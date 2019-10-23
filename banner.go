package banner

var DefaultFont = "big"

type Banner struct {
	text string
	font *Font
}

func New(text string) Banner {
	return Banner{
		text: text,
		font: LoadFont(DefaultFont),
	}
}

func (self Banner) Font(name string) Banner {
	return Banner{
		font: LoadFont(name),
	}
}

func (self Banner) String() string {

	return ""
}
