package tools

import (
	"warbler/internal/db"
	"encoding/json"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

type Memory struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
}

func CreateMemoryTool(g *genkit.Genkit) *ai.ToolDef[Memory, string] {
	return genkit.DefineTool(g, "create_memory",
		"Create a memory with description",
		func(ctx *ai.ToolContext, input Memory) (string, error) {
			_, err := db.DB.Exec("INSERT INTO memories(description) VALUES(?)", input.Description)
			if err != nil {
				return "", err
			}

			return "Memory created", nil
		},
	)
}

type DeleteMemory struct {
	ID uint `json:"id"`
}

func DeleteMemoryTool(g *genkit.Genkit) *ai.ToolDef[DeleteMemory, string] {
	return genkit.DefineTool(g, "delete_memory",
		"Delete a memory",
		func(ctx *ai.ToolContext, input DeleteMemory) (string, error) {
			_, err := db.DB.Exec("DELETE FROM memories WHERE id = ?", input.ID)
			if err != nil {
				return "", err
			}

			return "Memory created", nil
		},
	)
}

type FetchMemoriesInput struct{}

func FetchMemories(g *genkit.Genkit) *ai.ToolDef[FetchMemoriesInput, string] {
	return genkit.DefineTool(g, "fetch_memories",
		"Fetch all Memories registered",
		func(ctx *ai.ToolContext, input FetchMemoriesInput) (string, error) {
			rows, err := db.DB.Query("SELECT * FROM memories")
			if err != nil {
				return "", err
			}
			defer rows.Close()

			memories := []Memory{}
			for rows.Next() {
				var e Memory
				if err := rows.Scan(&e.ID, &e.Description); err != nil {
					return "", err
				}
				memories = append(memories, e)
			}

			json, err := json.Marshal(memories)
			if err != nil {
				return "", err
			}

			return string(json), nil
		},
	)
}
