// Go - REST API, standard library only. Reads input.json.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const actor = "data_foundry~woocommerce-scraper"

func main() {
	token := os.Getenv("APIFY_TOKEN")

	input, err := os.ReadFile("input.json")
	if err != nil {
		panic(err)
	}

	runURL := fmt.Sprintf("https://api.apify.com/v2/acts/%s/runs?token=%s", actor, token)
	resp, err := http.Post(runURL, "application/json", bytes.NewReader(input))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var started struct {
		Data struct {
			DefaultDatasetID string `json:"defaultDatasetId"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&started); err != nil {
		panic(err)
	}

	itemsURL := fmt.Sprintf(
		"https://api.apify.com/v2/datasets/%s/items?token=%s&limit=5",
		started.Data.DefaultDatasetID, token,
	)
	items, err := http.Get(itemsURL)
	if err != nil {
		panic(err)
	}
	defer items.Body.Close()

	out, _ := io.ReadAll(items.Body)
	fmt.Println(string(out))
}
