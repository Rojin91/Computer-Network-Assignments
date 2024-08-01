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

This is my explanation of the TCP packet capture file and the analysis of the connection initialization, maintenance, and termination processes. The capture file was analyzed using Wireshark, and the key steps and packets involved in these processes are outlined below.Please note that here I have only highlighted the request-response cycle of client and server with SYN number.Here `SYN`, `SYN-ACK` ,`ACK` have the major role in depicting the `TCP connection`
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
  - ` No. 43 - 64956 → 8080 [SYN] Seq=0 Win=65535 Len=0 MSS=16324 WS=64 TSval=3357812407 TSecr=0 SACK_PERM`

- **SYN-ACK Packet**: The server responds with a SYN-ACK packet, acknowledging the SYN request and sending its own SYN to the client.
  - `No. 44 - 8080 → 64956 [SYN, ACK] Seq=0 Ack=1 Win=65535 Len=0 MSS=16324 WS=64 TSval=1015158397 TSecr=3357812407 SACK_PERM`

- **ACK Packet**: The client acknowledges the server’s response with an ACK packet.
  - `No. 45 - 64956 → 8080 [ACK] Seq=1 Ack=1 Win=407744 Len=0 TSval=3357812407 TSecr=1015158397`

## 2. Connection Maintenance

Once the connection is established, data can be transferred between the client and server. This phase includes sending and receiving data packets, maintaining the connection’s state.

**Packets Involved:**
- **Data Transfer**: Actual data being transferred between the client and server.
  - `No. 47 - 64956 → 8080 [PSH, ACK] Seq=1 Ack=1 Win=407744 Len=873 TSval=3357812407 TSecr=1015158397 [TCP segment of a reassembled PDU]`
  - `No. 48 - 8080 → 64956 [ACK] Seq=873 Ack=874 Win=407744 Len=0 TSval=1015158398 TSecr=3357812407`
  - `No. 49 - 8080 → 64956 [PSH, ACK] Seq=874 Ack=874 Win=407744 Len=512 TSval=1015158398 TSecr=3357812407 [TCP segment of a reassembled PDU]`
  - `No. 50 - 64956 → 8080 [ACK] Seq=874 Ack=1386 Win=407744 Len=0 TSval=3357812408 TSecr=1015158398`
  - `No. 51 - 64956 → 8080 [PSH, ACK] Seq=1386 Ack=1386 Win=407744 Len=512 TSval=3357812408 TSecr=1015158398 [TCP segment of a reassembled PDU]`

- **HTTP Requests and Responses**: Exchange of HTTP requests and responses.
  - `No. 127 - 64956 → 8080 [POST /upload HTTP/1.1 (PNG)]`

## 3. Connection Termination

When the data transfer is complete, the connection is closed  using a termination procedure known as the **four-way handshake**.

**Packets Involved:**
- **FIN Packet**: One side initiates the connection termination by sending a FIN (Finish) packet.
  - `No. 129 - 64956 → 8080 [FIN, ACK] Seq=462664 Ack=467072 Win=407744 Len=0 TSval=3357812407 TSecr=1015158397`

- **ACK Packet**: The other side acknowledges the FIN packet with an ACK.
  - `No. 130 - 8080 → 64956 [ACK] Seq=467072 Ack=462665 Win=292288 Len=0 TSval=1015158401 TSecr=3357812407`

- **FIN Packet**: The other side sends its own FIN packet to complete the termination process.
  - `No. 131 - 8080 → 64956 [FIN, ACK] Seq=467072 Ack=462665 Win=292288 Len=0 TSval=1015158401 TSecr=3357812407`

- **Final ACK**: The initial side acknowledges the final FIN packet.
  - `No. 132 - 64956 → 8080 [ACK] Seq=462665 Ack=467073 Win=407744 Len=0 TSval=3357812407 TSecr=1015158401`

      
