package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerScore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	score PlayerScore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, p.score.GetPlayerScore(player))
}
