package main

//go run main.go client.go server.go
func main() {
	go StartServer()
	StartClient()
}
