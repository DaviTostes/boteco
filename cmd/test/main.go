package main

import (
	"boteco/internal/config"
	"boteco/internal/gen"
	"fmt"
	"log/slog"
	"os"
	"time"
)

func main() {
	c, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	g, err := gen.InitGenkit(c.Gemini.ApiKey)
	if err != nil {
		panic(err)
	}

	prompt := "When is the next brazil game in the world cup?"

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	})))

	stream := gen.GenerateStream(g, gen.BuildSystemPrompt(time.Now()), prompt, gen.Tools, nil)
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
