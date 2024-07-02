# IPv6 Header Information

## Question 1: How long is an IPv6 header?
IPv6 header is always fixed at 40 bytes (320 bits).

## Question 2: What are the different fields in the header?
+------------------+--------------------+----------------------+
| Version (4 bits) | Traffic Class (8 bits) | Flow Label (20 bits) |
+------------------+--------------------+----------------------+
| Payload Length (16 bits) | Next Header (8 bits) | Hop Limit (8 bits) |
+-------------------------------------------------------------+
| Source Address (128 bits) |
+-------------------------------------------------------------+
| Destination Address (128 bits) |
+-------------------------------------------------------------+

## Question 3: What is the purpose of each header field.

1. **Version (4 bits)**:
   - Indicates the IP version number. For IPv6, this field is always set to 6.

2. **Traffic Class (8 bits)**:
   - Used for packet classification and prioritization, helping in managing Quality of Service (QoS).

3. **Flow Label (20 bits)**:
   - Used to identify packet flows that require special handling by routers.

4. **Payload Length (16 bits)**:
   - Indicates the length of the payload, i.e., the data carried after the IPv6 header. It does not include the length of the header itself.

5. **Next Header (8 bits)**:
   - Identifies the type of header immediately following the IPv6 header. It indicates whether the next header is a transport layer protocol, an extension header, or an upper layer protocol.

6. **Hop Limit (8 bits)**:
   - Specifies the maximum number of hops (routers) that the packet can pass through. Each router decreases the value by one, and the packet is discarded when the value reaches zero.

7. **Source Address (128 bits)**:
   - The IPv6 address of the originator of the packet.

8. **Destination Address (128 bits)**:
   - The IPv6 address of the intended recipient of the packet.

##Question 4: In wireshark locate an IPv6 packet and discuss the header present.
<img src - 'IPv6 Packets-Wireshark.png'>
