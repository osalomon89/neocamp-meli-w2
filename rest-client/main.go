package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_go-core/pkg/rusty"
)

const port = ":8080"

type Item struct {
	Code        string  `json:"code"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

type Response struct {
	Items []Item `json:"items"`
}

type ClientError struct {
	Message    string
	StatusCode int
}

func main() {
	r := gin.Default()

	r.GET("/me/v1/items", getItems)
	r.GET("/me/v1/items/:id", getItemByID)

	log.Println("Server listening on port", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalln(err)
	}
}

func getItems(ctx *gin.Context) {
	resp, err := http.Get("http://localhost:8000/items")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		errorMessage := fmt.Sprintf("unexpected response status %v", resp.StatusCode)
		errorDetail := fmt.Sprintf("error getting items from API, %s", errorMessage)

		fmt.Println(errorDetail)

		ctx.JSON(http.StatusServiceUnavailable, ClientError{Message: errorDetail, StatusCode: resp.StatusCode})
		return
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result Response
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"data":  "error unmarshalling data",
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func getItemByID(ctx *gin.Context) {
	id := ctx.Param("id")

	t := http.Transport{
		IdleConnTimeout:     5 * time.Second,
		MaxConnsPerHost:     100,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
	}

	httpClient := &http.Client{
		Timeout:   3 * time.Second,
		Transport: &t,
	}

	endpoint, err := rusty.NewEndpoint(httpClient, "http://localhost:8000/items/{id}")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	resp, err := endpoint.Get(ctx, rusty.WithParam("id", id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		errorMessage := fmt.Sprintf("unexpected response status %v", resp.StatusCode)
		errorDetail := fmt.Sprintf("error getting books from API, %s", errorMessage)

		fmt.Println(errorDetail)

		ctx.JSON(http.StatusServiceUnavailable, ClientError{Message: errorDetail, StatusCode: resp.StatusCode})
		return
	}

	var result Item
	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"data":  "error unmarshalling data",
		})
	}

	ctx.JSON(http.StatusOK, result)
}
