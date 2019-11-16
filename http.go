package main

import (
	"fmt"
	"net/http"
	"net"

	ratelimit "github.com/tingShean/projectRL/ratelimit"
)

type visitor struct {
	limiter *ratelimit.Limiter
}

var visitors = make(map[string]*visitor)

// server
func main() {
	fmt.Println("server start")

	http.HandleFunc("/", Hello_world)

	http.ListenAndServe(":9000", nil)
}

// if ip not exists
func set_visitor(ip string) *ratelimit.Limiter {
	limiter := ratelimit.NewRateLimiter(1)
	visitors[ip] = &visitor{limiter}
	return limiter
}

// add limit if ip exists
func get_visitor(ip string) *ratelimit.Limiter {
	visitor, exists := visitors[ip]
	if !exists {
		return set_visitor(ip)
	}

	// add limit
	visitor.limiter.AddLimit()
	return visitor.limiter
}

func Hello_world(w http.ResponseWriter, req *http.Request) {
	res := "hello world %v rate limit %v"

	// get ip
	ip := req.Header.Get("X-Forwarded-For")

	// can not get ip
	if len(ip) == 0 {
		ip, _, _ = net.SplitHostPort(req.RemoteAddr)
		// if ip always empty, set default
		if len(ip) == 0 {
			ip = "0.0.0.0"
		}
	}

	// calculate
	lim := get_visitor(ip)

	// header
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if lim.GetLimit() >= ratelimit.Limit(61) {
		fmt.Fprintf(w, res, ip, "Error")
	} else {
		fmt.Fprintf(w, res, ip, lim.GetLimit())
	}
}
