package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprintf(w, GetPlayerScore(player))
}

func GetPlayerScore(player string) string {
	if player == "Pepper" {
		return "20"
	}

	if player == "Rob" {
		return "10"
	}
	return ""
}
