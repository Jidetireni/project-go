package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

var (
	visitCounts = make(map[string]int)

	mu sync.Mutex
)

func count(w http.ResponseWriter, r *http.Request) {
	clientIP := r.RemoteAddr

	clientIP = strings.Split(clientIP, ":")[0]
	mu.Lock()

	visitCounts[clientIP]++

	visitCount := visitCounts[clientIP]

	mu.Unlock()

	fmt.Fprintf()

}
