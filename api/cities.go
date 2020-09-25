package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, currentTime)

	cities := []string{
		"hcm",
		"ha noi",
	}
	data, _ := json.Marshal(cities)

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
