The header fields of ipv4 are as follows:
Version: Its value is 4. This specifies the version of the IP protocol.
Header Length: Its value is 20 bytes (5). This specifies the length of the IP header in 32-bit words.
Differentiated Services Field (DSCP): Its value is 0x00 (DSCP: CS0, ECN: Not-ECT). This is used for packet classification for Quality of Service (QoS).
Total Length: Its value is 41. This specifies the entire packet size, including the header and data, in bytes.
Identification: Its value is 0x89cd (35277). This is used for uniquely identifying fragments of an original IP datagram.
Flags: Its value is 0x2 (Don't Fragment). This indicates that the Don't Fragment (DF) flag is set, meaning the packet should not be fragmented.
Fragment Offset: Its value is 0. This indicates the position of the fragment in the original datagram, and since fragmentation is not used, it is set to 0.
Time to Live (TTL): Its value is 128. This specifies the maximum number of hops a packet can traverse before being discarded.
Protocol: Its value is 6 (TCP). This indicates the protocol used in the data portion of the IP datagram.
Header Checksum: Its value is 0x0000 (validation disabled). This is used for error-checking of the header, but validation is disabled in this capture.
Source Address: Its value is 192.168.18.7. This specifies the IP address of the sender.
Destination Address: Its value is 40.99.34.226. This specifies the IP address of the receiver.
