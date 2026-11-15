package core

import (
	"musa.io/bookreader/pdf"
	"log"
)

type ContextLoader struct {
	parser *pdf.Parser
}

func MakeContextLoader(path string) (*ContextLoader, error) {
	parser, err := pdf.NewParser(path)

	if err != nil {
		return nil, err
	}

	return &ContextLoader{parser: parser}, nil
}

func (c *ContextLoader) Load(pageNumber int) [] string {
	texts := []string{}

	// want to read not just the page number
	// we want content, so i will offset
	// to get a range of data

	startPageNumber := max(pageNumber - ContextLoaderOffset, 1)
	endPageNumber := pageNumber + ContextLoaderOffset

	if endPageNumber > c.parser.TotalPageNumber() {
		endPageNumber = c.parser.TotalPageNumber()
	}

	for currentPageNumber := startPageNumber; currentPageNumber <= endPageNumber; currentPageNumber += 1 {
		currentText, err := c.parser.Extract(currentPageNumber)

		if err != nil {
			log.Println("Error getting ", currentPageNumber, " ", err)
		}

		texts = append(texts, currentText)
	}


	return texts
}

