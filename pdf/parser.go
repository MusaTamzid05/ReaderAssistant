package pdf

import (
	"github.com/ledongthuc/pdf"
	"os"
)

type Parser struct {
	fPointer *os.File
	reader *pdf.Reader
}

func NewParser(path string) (*Parser, error) {
	f, reader,  err := pdf.Open(path)

	if err != nil {
		return nil, err
	}

	parser := &Parser{reader: reader, fPointer: f}
	return parser, nil

}

func (p *Parser) Extract(pageNumber int) (string, error) {
	page := p.reader.Page(pageNumber)
	text, err := page.GetPlainText(nil)

	return text, err
}

func (p *Parser) TotalPageNumber() int {
	return p.reader.NumPage()
}
