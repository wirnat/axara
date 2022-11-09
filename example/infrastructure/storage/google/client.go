package google

import (
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/option"
)

func InitClient() (*storage.Client, error) {
	ctx := context.Background()
	return storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
}
