package main

import (
	"boteco/internal/gen"
	"fmt"
)

func main() {
	g, err := gen.InitGenkit()
	if err != nil {
		panic(err)
	}

	prompt := "How to handle channels in go"

	stream := gen.GenerateStream(g, gen.SystemPrompt, prompt, gen.Tools, nil, nil)
	for result, err := range stream {
		if err != nil {
			panic(err)
		}

		if result.Done {
			fmt.Println("\nDone!")
			break
		}

		fmt.Print(result.Chunk.Text())
	}
}
