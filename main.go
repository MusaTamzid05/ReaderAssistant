package main

import (
	"musa.io/bookreader/core"
	"musa.io/bookreader/llm"
	"log"
	"fmt"
)
func main() {
	loader, err := core.MakeContextLoader("/home/musa/Downloads/build_an_ai_agent_from_scratch.pdf")

	if err != nil {
		log.Fatalln(err)
	}
	texts := loader.Load(137)

	document := ""
	for _, text := range texts {
		document += text
	}

	ollama := llm.MakeOllama()
	response, err := ollama.Ask("explain the func load_mcp_tools, please show the func with comments ", document)

	if err != nil {
		log.Fatalln(err)
	}


	fmt.Println(response)




}
