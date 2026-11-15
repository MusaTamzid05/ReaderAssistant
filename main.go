package main

import (
	"musa.io/bookreader/core"
	"log"
)
func main() {
	loader, err := core.MakeContextLoader("/home/musa/Downloads/build_an_ai_agent_from_scratch.pdf")

	if err != nil {
		log.Fatalln(err)
	}
	texts := loader.Load(124)


	for _, text := range texts {
		log.Println(text)

	}




}
