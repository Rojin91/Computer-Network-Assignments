Q) What is the value of each of the header fields ? Explain why the value is what it is ?
Version: The version of the IP protocol being used. For IPv4, this value is always 4.

Header Length (IHL): Specifies the length of the IP header. The minimum value is 5 (20 bytes), and the maximum is 15 (60 bytes).

Differentiated Services (DSCP): This field is used for quality of service (QoS) features. It helps in classifying and managing network traffic.

Explicit Congestion Notification (ECN): This field is used for network congestion management without dropping packets.

Total Length: Specifies the total size of the packet, including both the header and the data. The size ranges from 20 bytes to 65,535 bytes.

Identification: A unique identifier for each packet. It is used to help reassemble fragmented packets.

Flags: Three bits used to control fragmentation. The important ones are:
    * DF (Don't Fragment): If set, the packet cannot be fragmented.
    * MF (More Fragments): Indicates if there are more fragments.
    * 
Fragment Offset: Indicates the position of the fragment in the original packet. It helps in reassembling the fragments in the correct order.

Time to Live (TTL): Limits the packet's lifetime to prevent it from circulating indefinitely. It is decreased by one by each router the packet passes through.

Protocol: Indicates the protocol used in the data portion of the packet (e.g., TCP, UDP).

Header Checksum: Used for error-checking the header. If the checksum is incorrect, the packet is discarded.

Source Address: The IP address of the sender.

Destination Address: The IP address of the receiver.

Options: Optional field for additional functionalities. It is not commonly used and can increase the header size.
