package main

import (
	"context"
	"log"
	"os"

	"github.com/kaihendry/picofeed"
)

func main() {
	bucket := os.Getenv("BUCKET")
	if bucket == "" {
		bucket = "hendry.iki.fi"
		os.Setenv("BUCKET", bucket)
	}
	log.Printf("Using bucket: %s", bucket)
	ctx := context.Background()
	picofeed.Refresh(ctx)
}
