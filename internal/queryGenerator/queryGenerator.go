package queryGenerator

import (
	"context"
	"fmt"
	"sistemaTenis/internal/product"
	"sistemaTenis/internal/promptGenerator"
	"time"

	ollama "github.com/prathyushnallamothu/ollamago"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Init() {

	tenis := product.GetShoesByParam(&bson.M{
		"url": 1,
		"_id": 1,
	})

	prompt := promptGenerator.GeneratePrompt(&tenis)

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
