The different header fields of udp protocol are as follows:

Source Port: Its value is 53. This specifies the port number of the source, which in this case is a DNS server (port 53 is typically used for DNS).

Destination Port: Its value is 53505. This specifies the port number of the destination, which is a randomly chosen high-numbered port on the client machine.

Length: Its value is 118. This specifies the length of the UDP header and the UDP payload, in bytes.

Checksum: Its value is 0x90db. This is used for error-checking of the header and data. The checksum is calculated over the entire UDP datagram.
