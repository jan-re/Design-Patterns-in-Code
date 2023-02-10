package main

type application struct {
}

func (a *application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "OK"
	} else if url == "/create/user" && method == "POST" {
		return 201, "User created"
	}

	return 404, "Not found"
}
