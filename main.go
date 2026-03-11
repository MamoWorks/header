package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Record struct {
	Time    string              `json:"time"`
	Method  string              `json:"method"`
	Path    string              `json:"path"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
}

var (
	records []Record
	mu      sync.Mutex
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		rec := Record{
			Time:    time.Now().Format("2006-01-02 15:04:05"),
			Method:  r.Method,
			Path:    r.URL.String(),
			Headers: r.Header,
			Body:    string(body),
		}

		mu.Lock()
		records = append(records, rec)
		if len(records) > 10 {
			records = records[len(records)-10:]
		}
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "recorded"})
		return
	}

	// GET or any other method: return latest records
	mu.Lock()
	data := make([]Record, len(records))
	copy(data, records)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(data)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on :8899")
	if err := http.ListenAndServe(":8899", nil); err != nil {
		fmt.Println("Error:", err)
	}
}
