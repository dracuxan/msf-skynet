package msf

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

func RunMultiplexer(port int, hostProxy map[string]string) {
	// Create proxies map
	proxies := make(map[string]*httputil.ReverseProxy)
	for host, targetURL := range hostProxy {
		remote, err := url.Parse(targetURL)
		if err != nil {
			log.Fatalf("Unable to parse proxy target %s: %v", targetURL, err)
		}
		proxies[host] = httputil.NewSingleHostReverseProxy(remote)
	}

	// Set up router
	r := mux.NewRouter()
	for host, proxy := range proxies {
		r.Host(host).Handler(proxy)
	}

	// Start server
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting reverse proxy on port %d", port)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
