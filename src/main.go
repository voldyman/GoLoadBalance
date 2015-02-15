package main

func main() {
	backends := NewBackends()
	backends.Add("localhost:8888",
		"localhost:8881",
		"localhost:8882",
		"localhost:8883")

	listenPort := 8050
	startServer(listenPort, backends)
}
