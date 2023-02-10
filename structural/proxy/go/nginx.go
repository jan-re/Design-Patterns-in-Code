package main

type nginx struct {
	application       *application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func newNginxProxy() nginx {
	return nginx{
		application:       &application{},
		maxAllowedRequest: 2,
		rateLimiter:       map[string]int{},
	}
}

// Provides initial handling of the method where ratelimiting is enforced.
// Afterwards, passes the request to the real application underneath.
func (n *nginx) handleRequest(url, method string) (int, string) {
	if !n.isAllowedByRateLimiter(url) {
		return 429, "Too many requests"
	}

	return n.application.handleRequest(url, method)
}

func (n *nginx) isAllowedByRateLimiter(url string) bool {
	n.rateLimiter[url] = n.rateLimiter[url] + 1

	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}

	return true
}
