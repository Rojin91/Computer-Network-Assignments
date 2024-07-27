package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    // Update the server address to localhost (assuming server runs on the same machine)
    serverAddress := "localhost:8888"

    // 1. Connect to the server
    conn, err := net.Dial("tcp", serverAddress)
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        return
    }
    defer conn.Close() // Ensure the connection is closed when done

    // 2. Read message from standard input
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter message: ")
    message, _ := reader.ReadString('\n')

    // 3. Send the message to the server
    _, err = conn.Write([]byte(message))
    if err != nil {
        fmt.Println("Error sending message:", err)
        return
    }

    // 4. Receive response from the server
    reply := make([]byte, 1024) // Buffer for the server's response
    _, err = conn.Read(reply)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return
    }

    // 5. Print the server's response
    fmt.Println("Server reply:", string(reply))
}
