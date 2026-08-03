package searchAgent

import (
	"context"
	"fmt"
	"sistemaTenis/internal/product"

	"time"

	ollama "github.com/prathyushnallamothu/ollamago"
)

func Init() {
	tenis := product.Tenis{
		Name:     "Luka 5 \"Luka Lifestyle\"Big Kids' Basketball Shoes",
		Price:    "$105",
		Category: "Basketball",
		Status:   "pending",
	}

	prompt := generatePrompt(&tenis)

	client := ollama.NewClient(
		ollama.WithTimeout(time.Minute * 10),
	)

	resp, err := client.Chat(context.Background(), ollama.ChatRequest{
		Model:    "qwen3",
		Messages: prompt,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Message.Content)
}
