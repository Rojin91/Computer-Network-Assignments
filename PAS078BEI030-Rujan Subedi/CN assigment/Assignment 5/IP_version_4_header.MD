Q) What is the value of each of the header fields? Explain why the value is what it is.
ans:
 According to the packet captured in Wireshark, the values of header fields are given as follows:
 1. Version: Its value is 4. It identifies the version of IP in use.
 2. Internet header length(IHL): Its value is 5(20 bytes). It specifies the size of the IP header. 
 3. Differentiated services code point(DSCP + ECN): Its value is 0x00. DSCP is used for specifying the priority of the packet, and ECN is used for congestion notification. Both are set to 0 in this case, indicating default handling.
 4. Total Length: Its value is 56. It indicates the total length of the IP packet (header + data) in bytes. In this case, the total length is 56 bytes.
 5. Identification: Its value is 0x287e (10366 in decimal). This field is used for uniquely identifying the fragments of an original IP packet.
 6. Flags: Its value is 0x0. The flags field has three bits, indicating control flags: reserved, don't fragment (DF), and more fragments (MF). Here, all are set to 0, meaning no fragmentation control flags are set.
 7. Fragment Offset: Its value is 0.  This indicates where in the datagram this fragment belongs. A value of 0 means this is either the first fragment or the packet is not fragmented.
 8. Time to Live (TTL): Its value is 128. : This field indicates the maximum time the packet is allowed to exist in the network before being discarded. It is set to 128 in this case.
 9. Protocol: Its value is 17. This field specifies the next-level protocol used in the data portion of the IP datagram. A value of 17 corresponds to UDP(User Datagram Protocol).
10. Header Checksum: Its value is 0x6cdc. This field is used for error-checking in the header.
11. Source address: Its value is 192.168.18.9. This is the IP address of the sender of the packet.
12. Destination Address: Its value is 192.168.18.1. This is the IP address of the intended receiver of the packet.
