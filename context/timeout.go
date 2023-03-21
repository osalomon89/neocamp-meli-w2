package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const shortDuration10 = 10000 * time.Millisecond

func mainfdsf() {
	// Set timeout in context
	ctx, _ := context.WithTimeout(context.Background(), shortDuration10)

	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com/", nil)

	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}

	fmt.Println("Response received, status code:", res.StatusCode)
}
