UDP Header

Q) What is the size of the UDP header?
ans: The size of the UDP header is 8 bytes long.

Q) What are the values of different header fields in the UDP header and why are they so?
ans: According to the packet captured in Wireshark, the values of header fields are given as follows:
1.  Source port: Its value is 7866. This indicates the port number on the source device that sent the packet.
2.  Destination port: Its value is 6881. This indicates the port number on the destination device that is intended to receive the packet.
3.  Length: Its value is 109. This field specifies the length of the UDP header and data in bytes.
4.  Checksum: Its value is 0xe098. This field is used to verify the integrity of the UDP header and data.

TCP Header

Q)What is the size of the TCP header?
ans: The size of the TCP header can vary, but the minimum size is 20 bytes. This is the size of the header without any options.

Q)What are the different fields in the TCP header?
ans: According to the packet captured in Wireshark, the values of header fields are given as follows:
1. Source port: Its value is 44521. The port number from which the packet was sent.
2. Destination port: Its value is 6881. The port number to which the packet is addressed.
3. Sequence number: Its value is 0. This is because the initial sequence number of the SYN packet is set to 0.
4. Acknowledgment number: Its value is also 0. This is because the acknowledgment number is not yet relevant until the connection is established.
5. Data offset(Header Length): Its value is 10. This field specifies the length of the TCP header and data in bytes.
6. Flags: Its value is 0x002. This flag indicates that this is a SYN packet that initiates a connection.
7. Window Size: Its value is 65535. It specifies how much data the sender is willing to accept.
8. Checksum: Its value is 0x0d84 but is unverified in the packet. This is used for error-checking the header and data.
9. Urgent pointer: Its value is 0. This is because the URG flag is not set.
10. Options: It includes additional TCP options:
    -Maximum Segment Size(MSS)
    -Selective Acknowledgment (SACK) permitted
    -Timestamps
    -No-Operation (NOP)
    -Window scale

Detailed breakdown of options:
1. Maximum Segment Size: It specifies the largest amount of data, in bytes, that the device is willing to receive in a single TCP segment.
2. SACK Permitted: It indicates that the device supports selective acknowledgment.
3. Timestamps: It is used for round-trip time measurement and PAWS (Protection Against Wrapped Sequence numbers).
4. No-Operation (NOP): It is used to align fields to a 32-bit boundary.
5. Window Scale: It is used to scale the window size field, allowing for a larger range of values.


