// Go - REST API, standard library only
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

var input = map[string]any{
	"url": "https://example-store.com",
	"resource": "products",
	"limit": 100,
}

func main() {
	token := os.Getenv("APIFY_TOKEN")
	body, _ := json.Marshal(input)
	url := fmt.Sprintf("https://api.apify.com/v2/acts/%s/runs?token=%s", actor, token)
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var started struct {
		Data struct{ ID, DefaultDatasetID string } `json:"data"`
	}
	raw, _ := io.ReadAll(resp.Body)
	json.Unmarshal(raw, &started)

	itemsURL := fmt.Sprintf("https://api.apify.com/v2/datasets/%s/items?token=%s&limit=5",
		started.Data.DefaultDatasetID, token)
	items, _ := http.Get(itemsURL)
	defer items.Body.Close()
	out, _ := io.ReadAll(items.Body)
	fmt.Println(string(out))
}
