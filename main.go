package main

func main() {
	server := CreateAPISever(":5173")
	server.Run()
}
