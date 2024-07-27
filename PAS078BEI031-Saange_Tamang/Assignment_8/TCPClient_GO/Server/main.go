package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    // Listen on TCP port 8888 on all available interfaces
    ln, err := net.Listen("tcp", ":8888")
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    defer ln.Close()

    fmt.Println("Server is listening on port 8888...")

    // Accept a single connection
    conn, err := ln.Accept()
    if err != nil {
        fmt.Println("Error accepting connection:", err)
        return
    }
    defer conn.Close()

    // Handle the connection
    handleConnection(conn)
}

func handleConnection(conn net.Conn) {
    reader := bufio.NewReader(conn)

    // Read message from the client
    message, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading message:", err)
        return
    }

    fmt.Printf("Received message: %s", message)

    // Send a response back to the client
    response := "Message received\n"
    _, err = conn.Write([]byte(response))
    if err != nil {
        fmt.Println("Error sending response:", err)
        return
    }
}
