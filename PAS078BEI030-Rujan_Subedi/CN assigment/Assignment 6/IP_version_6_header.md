Q. How long is an IPv6 header?
ans: The length of the IPv6 header is 40 bytes which is always the same.

Q. What are the different fields in the header? What is the purpose of each header field?
ans: According to the packet captured in Wireshark, the values of header fields are given as follows:
  1. Version (4 bits): The first field is 0110, which indicates the protocol version. For IPv6, this value is always 6.
  2. Traffic Class (8 bits): This field is 0000 0000, which is used for priority and quality of service (QoS). In this packet,        it is set to 0, meaning default priority.
  3. Flow Label (20 bits): The flow label is 0000 0000 0010 1010 1011 0010, which is used for labeling sequences of packets for       special handling. In this packet, the flow label is 0x2ab20.
  4. Payload Length (16 bits): The payload length is 1448 bytes, which specifies the length of the payload (the data following        the header).
  5. Next Header (8 bits): The value is 44, which indicates that the next header is a Fragment Header. The Fragment Header is        used for fragmenting packets.
  6. Hop Limit (8 bits): The hop limit is 64, which specifies the maximum number of hops the packet can traverse. Each router        that forwards the packet decrements this value by one. When it reaches 0, the packet is discarded.
  7. Source Address (128 bits): The source address is 2405:ac0:1207:9014::1. This is the IPv6 address of the sender.
  8. Destination Address (128 bits): The destination address is 2600:6c46:6d00:42dc:d0b2:b0e1:516:fe89. This is the IPv6 
     address of the receiver.
