package reporting

import (
	"context"

	"cloud.google.com/go/storage"
)

const (
    bucketName = "a-bucket"
)

type StorageClient struct {
	ctx    context.Context
	client *storage.Client
}

type StorageClienter interface {
    GetBucketWriter(ctx context.Context, name string) *storage.Writer
	GetBucketReader(ctx context.Context, name string) (*storage.Reader, error)
	GetBucketObjectIterator(ctx context.Context) *storage.ObjectIterator
}

func (s StorageClient) GetBucketWriter(ctx context.Context, name string) *storage.Writer {
	return s.client.Bucket(bucketName).Object(name).NewWriter(ctx)
}

func (s StorageClient) GetBucketReader(ctx context.Context, name string) (*storage.Reader, error) {
	return s.client.Bucket(bucketName).Object(name).NewReader(ctx)
}

func (s StorageClient) GetBucketObjectIterator(ctx context.Context) *storage.ObjectIterator {
	return s.client.Bucket(bucketName).Objects(ctx, &storage.Query{})
}

