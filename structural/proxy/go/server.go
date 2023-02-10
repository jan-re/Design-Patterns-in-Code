package main

// It is crucial that the application and the proxy share an interface.
type server interface {
	handleRequest(string, string) (int, string)
}
