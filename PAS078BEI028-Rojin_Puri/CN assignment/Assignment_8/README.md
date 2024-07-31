## Questions
- In the programming language of your choice, write a web application that allows you to upload a file.
- Capture the traffic in Wireshark while uploading a file. Follow the TCP stream and explain connection initialization, connection maintenance, and connection termination.

## Tasks Done by John Doe
#### 1. Made a simple PNG uploader in Go
- Captured the TCP Streams while uploading the file in `WIRESHARK`
- Understood the TCP `Connection-Initialization`, `Connection-Maintenance`, and `Connection-Termination`
- Captured the TCP Packets in `WIRESHARK`
      
#### 2. Made a simple TCPClient to understand the TCP in `GO` for `FUN`
- To run TCPClient in your local machine, just read the `README` in the TCPClient directory

## TCP Packet Analysis

Th explanation of the TCP packet capture file and the analysis of the connection initialization, maintenance, and termination processes. The capture file was analyzed using Wireshark, and the key steps and packets involved in these processes are outlined below. Please note that here John has only highlighted the request-response cycle of the client and server with SYN number. Here `SYN`, `SYN-ACK`, and `ACK` have the major role in depicting the `TCP connection`.

## Overview

The TCP packet capture file contains packets exchanged between a client and a server. There are mainly 3 steps involved in TCP:
1. Connection Initialization
2. Connection Maintenance
3. Connection Termination

## 1. Connection Initialization

This is the initial step where a TCP connection is established between a client and a server. It involves a handshake procedure known as the **three-way handshake**.

**Packets Involved:**
- **SYN Packet**: The client (source port 50919) initiates the connection by sending a SYN packet to the server (destination port 9000).
  - `No. 58 - 50919 → 9000 [SYN] Seq=0 Win=65535 Len=0 MSS=16324 WS=64 TSval=769422296 TSecr=0 SACK_PERM`

- **SYN-ACK Packet**: The server responds with a SYN-ACK packet, acknowledging the SYN request and sending its own SYN to the client.
  - `No. 59 - 9000 → 50919 [SYN, ACK] Seq=0 Ack=1 Win=65535 Len=0 MSS=16324 WS=64 TSval=1395474443 TSecr=769422296 SACK_PERM`

- **ACK Packet**: The client acknowledges the server’s response with an ACK packet.
  - `No. 60 - 50919 → 9000 [ACK] Seq=1 Ack=1 Win=407744 Len=0 TSval=769422296 TSecr=1395474443`

## 2. Connection Maintenance

Once the connection is established, data can be transferred between the client and server. This phase includes sending and receiving data packets, maintaining the connection’s state.

**Packets Involved:**
- **Data Transfer**: Actual data being transferred between the client and server.
  - `No. 61 - 9000 → 50919 [PSH, ACK] Seq=1 Ack=1 Win=407744 Len=937 TSval=769422296 TSecr=1395474443 [TCP segment of a reassembled PDU]`
  - `No. 62 - 50919 → 9000 [ACK] Seq=1 Ack=937 Win=407744 Len=0 TSval=769422296 TSecr=1395474443`
  - `No. 63 - 9000 → 50919 [PSH, ACK] Seq=862 Ack=1 Win=406912 Len=16388 TSval=769422296 TSecr=1395474443 [TCP segment of a reassembled PDU]`
  - `No. 64 - 50919 → 9000 [ACK] Seq=1 Ack=17246 Win=407744 Len=0 TSval=769422296 TSecr=1395474443`
  - `No. 65 - 9000 → 50919 [PSH, ACK] Seq=49942 Ack=1 Win=407744 Len=148 TSval=769422296 TSecr=1395474443 [TCP segment of a reassembled PDU]`
  - `No. 66 - 50919 → 9000 [ACK] Seq=1 Ack=4508 Win=407744 Len=0 TSval=769422296 TSecr=1395474443`

- **HTTP Requests and Responses**: Exchange of HTTP requests and responses.
  - `No. 67 - POST /upload HTTP/1.1 (PNG)`
  - `No. 68 - HTTP/1.1 200 OK (text/plain)`

## 3. Connection Termination

When the data transfer is complete, the connection is closed using a termination procedure known as the **four-way handshake**.

**Packets Involved:**
- **FIN Packet**: One side initiates the connection termination by sending a FIN (Finish) packet.
  - `No. 94 - 50919 → 9000 [FIN, ACK] Seq=26411 Ack=334 Win=407424 Len=0 TSval=769426843 TSecr=1395474890`

- **ACK Packet**: The other side acknowledges the FIN packet with an ACK.
  - `No. 95 - 9000 → 50919 [ACK] Seq=334 Ack=26411 Win=381376 Len=0 TSval=1395474890 TSecr=769426843`

- **FIN Packet**: The other side sends its own FIN packet to complete the termination process.
  - `No. 96 - 9000 → 50919 [FIN, ACK] Seq=334 Ack=26411 Win=381376 Len=0 TSval=1395474890 TSecr=769426843`

- **Final ACK**: The initial side acknowledges the final FIN packet.
  - `No. 97 - 50919 → 9000 [ACK] Seq=26411 Ack=335 Win=407424 Len=0 TSval=769426843 TSecr=1395474890`

