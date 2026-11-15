package main

import (
	"musa.io/bookreader/pdf"
	"log"
)
func main() {
	parser, err := pdf.NewParser("/home/musa/Downloads/build_an_ai_agent_from_scratch.pdf")

	if err != nil {
		log.Fatalln(err)
	}
	text, _ := parser.Extract(124)

	log.Println(text)



}
