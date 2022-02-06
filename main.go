package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/joho/godotenv"
)

func main() {
	// .envから環境変数読み込み
	_ = godotenv.Load()

	ctx := context.Background()
	c, err := vision.NewImageAnnotatorClient(ctx)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer c.Close()

	file, _ := os.Open("sample.png")
	defer file.Close()

	image, _ := vision.NewImageFromReader(file)

	texts, err := c.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, text := range texts {
		fmt.Println(text.Description)
	}
}
