### Prerequisites

- Go installed on your machine (version 1.16 or later).

### Running the Server

1. Open a terminal window.
2. Navigate to the `Server` directory:
    ```sh
    cd Server
    ```
3. Run the server:
    ```sh
    go run main.go
    ```

### Running the Client

1. Open another terminal window.
2. Navigate to the `Client` directory:
    ```sh
    cd Client
    ```
3. Run the client:
    ```sh
    go run main.go
    ```

4. Enter a message when prompted and press Enter.

## How It Works

- **Server**: Listens on TCP port 8888, accepts a connection, reads a message from the client, and sends a response.
- **Client**: Connects to the server on localhost:8888, sends a message, and prints the server's response.
