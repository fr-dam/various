package main

import (
	//"bytes"
	"context"
	"fmt"
	"io/ioutil"

	"cloud.google.com/go/storage"
)

func main() {
	ctx := context.Background()
	client, _ := storage.NewClient(ctx)
	bucketName := fmt.Sprintf("%s%s", "fr-869ti53tfja6byvdnqpws6jimka", "-org-promo-reports")
	bh := client.Bucket(bucketName)
	reader, err := bh.Object("1638380785").NewReader(ctx)
	if err != nil {
		fmt.Println(err)
	}
	//var object bytes.Buffer
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf))
}
