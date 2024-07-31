## Questions
- In the programming language of your choice write a web application that allows you to upload a file.
- Capture a traffic in wireshark while uploading a file. Follow the TCP stream and explain conncection initialization , connection maintainance and conncection termination.

## Tasks Done
#### 1. Made a simple PNG uploader in Go
- Captured the TCP Streams while uploading the file in `WIRESHARK`
- Understood the TCP `Connection-Initialization` , `Connection-Maintenance` and `Connection-Termination`
- Captured the TCP Packets in `WIRESHARK`
      
#### 2. Made a simple TCPClient to understand the TCP in `GO` for `FUN`
- To run TCPClient in your local machine, just read the `README` in TCPClient directory

## TCP Packet Analysis

This is my explanation of the TCP packet capture file and the analysis of the connection initialization, maintenance, and termination processes. The capture file was analyzed using Wireshark, and the key steps and packets involved in these processes are outlined below.Please note that here I have only highlighted the request-response cycle of client and server with SYN number.Here `SYN`, `SYN-ACK` ,`ACK` have the major role in depicting the `TCP connection`.There might be some technical errors since I'm a noobie.Just ignore the errors.

## Overview

The TCP packet capture file contains packets exchanged between a client and a server. There are mainly 3 steps involved
in TCP.
1. Connection Initialization
2. Connection Maintenance
3. Connection Termination

## 1. Connection Initialization

This is inital step where a TCP connection is established between a client and a server. It involves a handshake procedure known as the **three-way handshake**.

**Packets Involved:**
- **SYN Packet**: The client (source port 63209) initiates the connection by sending a SYN packet to the server (destination port 8080).
  - `No. 38 - 63209 → 8080 [SYN] Seq=0 Win=65535 Len=0 MSS=16324 WS=64 TSval=3524908965 TSecr=0 SACK_PERM`

- **SYN-ACK Packet**: The server responds with a SYN-ACK packet, acknowledging the SYN request and sending its own SYN to the client.
  - `No. 39 - 8080 → 63209 [SYN, ACK] Seq=0 Ack=1 Win=65535 Len=0 MSS=16324 WS=64 TSval=65867666 TSecr=3524908965 SACK_PERM`

- **ACK Packet**: The client acknowledges the server’s response with an ACK packet.
  - `No. 40 - 63209 → 8080 [ACK] Seq=1 Ack=1 Win=407744 Len=0 TSval=3524908965 TSecr=65867666`

## 2. Connection Maintenance

Once the connection is established, data can be transferred between the client and server. This phase includes sending and receiving data packets, maintaining the connection’s state.

**Packets Involved:**
- **Data Transfer**: Actual data being transferred between the client and server.
  - `No. 41 - 8080 → 63209 [PSH, ACK] Seq=1 Ack=1 Win=407744 Len=937 TSval=65867667 TSecr=3524908965 [TCP segment of a reassembled PDU]`
  - `No. 42 - 63209 → 8080 [ACK] Seq=1 Ack=937 Win=407744 Len=0 TSval=3524908965 TSecr=65867667`
  - `No. 45 - 8080 → 63209 [PSH, ACK] Seq=862 Ack=1 Win=406912 Len=16388 TSval=65867668 TSecr=3524908965 [TCP segment of a reassembled PDU]`
  - `No. 46 - 63209 → 8080 [ACK] Seq=1 Ack=17246 Win=407744 Len=0 TSval=3524908966 TSecr=65867668`
  - `No. 50 - 8080 → 63209 [PSH, ACK] Seq=49942 Ack=1 Win=407744 Len=148 TSval=65867670 TSecr=3524908968 [TCP segment of a reassembled PDU]`
  - `No. 51 - 63209 → 8080 [ACK] Seq=1 Ack=4508 Win=407744 Len=0 TSval=3524908969 TSecr=65867670`

- **HTTP Requests and Responses**: Exchange of HTTP requests and responses.
  - `No. 56 - POST /upload HTTP/1.1 (PNG)`
  - `No. 58 - HTTP/1.1 200 OK (text/plain)`

## 3. Connection Termination

When the data transfer is complete, the connection is closed  using a termination procedure known as the **four-way handshake**.

**Packets Involved:**
- **FIN Packet**: One side initiates the connection termination by sending a FIN (Finish) packet.
  - `No. 80 - 63209 → 8080 [FIN, ACK] Seq=60221 Ack=334 Win=407424 Len=0 TSval=3524912415 TSecr=658680876`

- **ACK Packet**: The other side acknowledges the FIN packet with an ACK.
  - `No. 81 - 8080 → 63209 [ACK] Seq=334 Ack=60222 Win=347520 Len=0 TSval=658681316 TSecr=3524912415`

- **FIN Packet**: The other side sends its own FIN packet to complete the termination process.
  - `No. 82 - 8080 → 63209 [FIN, ACK] Seq=334 Ack=60222 Win=347520 Len=0 TSval=658681316 TSecr=3524912415`

- **Final ACK**: The initial side acknowledges the final FIN packet.
  - `No. 83 - 63209 → 8080 [ACK] Seq=60222 Ack=335 Win=407424 Len=0 TSval=3524912415 TSecr=658681316`

  ### Note: The sequence (SEQ) and acknowledgment (ACK) numbers in TCP packets are incremented accordingly to ensure reliable and ordered data transmission.This is the underlying principle of TCP.

      
      
 

