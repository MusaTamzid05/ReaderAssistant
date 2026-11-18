package core

import (
	"os/exec"
	"os"
	"fmt"
	"strconv"
	"bufio"
	"musa.io/bookreader/llm"
)

type Repl struct {
	loader *ContextLoader

}

func NewRepl(path string) (*Repl, error) {
	loader, err := MakeContextLoader(path)

	if err != nil {
		return nil, err
	}

	repl := Repl{loader: loader}
	return &repl, nil
}

func (r *Repl) Start() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	reader := bufio.NewReader(os.Stdin)
	ollama := llm.MakeOllama()
	running := true
	var pageNumber int
	var userInput string
	var err error

	for running {
		fmt.Println("Enter Page number : ")
		fmt.Scanln(&userInput)

		pageNumber, err = strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Invalid page num ", userInput)
			continue
		}

		fmt.Println("Question: ")
		userInput, _ := reader.ReadString('\n')

		if userInput == "exit" {
			running = false
			continue
		}


		texts := r.loader.Load(pageNumber)

		document := ""
		for _, text := range texts {
			document += text
		}


		response, err := ollama.Ask(userInput, document)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(response)


	}


}
