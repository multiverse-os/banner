package text

type Font struct {
	Name       string
	Characters map[string]string
}

func LoadFont(name string) *Font {
	return BigText
}

type Banner struct {
	Font *Font
}

func New() {}

func (self *Font) Banner(text string) {}
