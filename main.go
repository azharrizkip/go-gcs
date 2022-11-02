package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	gcs "cloud.google.com/go/storage"

	"test/storage"

	"google.golang.org/api/option"
)

var creds string = ``

var gcpAccessID string = ""
var gcpPrivateKey string = ""

func main() {
	gcsclient, err := gcs.NewClient(context.Background(), option.WithCredentialsJSON([]byte(creds)))
	if err != nil {
		panic(err)
	}

	image, err := os.ReadFile("cat.png")

	q := bytes.NewReader(image)

	gcs := storage.NewGCSAdapter(gcsclient, gcpAccessID, gcpPrivateKey)
	err = gcs.PutObject(context.Background(), "project-testing", "testing/cat.png", q, "image/png", nil)
	if err != nil {
		panic(err)
	}

	url, err := gcs.SignURL(context.Background(), http.MethodGet, "project-testing", "testing/cat.png", time.Second*30)
	if err != nil {
		panic(err)
	}

	fmt.Println(url)
}
