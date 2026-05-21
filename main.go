package main

import (
	"boteco/internal/db"
	"boteco/internal/gen"
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"

	"github.com/firebase/genkit/go/ai"
)

func main() {
	err := db.Connect()
	if err != nil {
		panic(err)
	}

	g, err := gen.InitGenkit()
	if err != nil {
		panic(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		<-ctx.Done()
		db.DB.Close()
		fmt.Println("\nBye!")
		os.Exit(0)
	}()

	reader := bufio.NewReader(os.Stdin)
	var messages []*ai.Message
	for {
		fmt.Print("\033[32m>\033[0m ")
		prompt, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			panic(err)
		}

		var resp string
		stream := gen.Generate(g, gen.SystemPrompt, prompt, gen.Tools, nil, messages)
		for result, err := range stream {
			if err != nil {
				panic(err)
			}
			if result.Done {
				resp = result.Response.Text()
				break
			}

			fmt.Print(result.Chunk.Text())
		}

		fmt.Println("")

		messages = append(messages, &ai.Message{
			Role: "user",
			Content: []*ai.Part{
				{Text: prompt},
			},
		})

		messages = append(messages, &ai.Message{
			Role: "model",
			Content: []*ai.Part{
				{Text: resp},
			},
		})
	}
}
