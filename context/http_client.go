package context

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func Client() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequest("GET", "http://localhost:8082", nil)
	if err != nil {
		log.Fatalf("failed to create request: %v", err)
	}

	req = req.WithContext(ctx)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("failed to make request: %v", err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
