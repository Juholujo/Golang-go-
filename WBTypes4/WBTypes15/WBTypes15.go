package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	req, err := http.NewRequest("GET", "https://example11111rrr1.com", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
}
