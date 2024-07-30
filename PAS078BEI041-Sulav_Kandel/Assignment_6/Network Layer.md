Network Layer: Routing Protocols

1) **The IP Header**  
     
* It is a Datagram(connectionless)service.  
* It will tell the information about how the packet will move.  
* Total 160 bits (20 byte).  
* Payload (Data coming from the transport layer.)= 0 – 65515 Bytes  
* Datagram= Header size \+ Payload= 65535 Bytes= 2^16

**Visual Representation**

Here’s a visual representation of how the header length affects the overall structure:

![IPV4](https://github.com/Rojin91/Computer-Network-Assignments/blob/main/PAS078BEI041-Sulav_Kandel/Assignment_6/IPV4.jpg)


Now, let’s discuss all the fields.
![IPV4](https://github.com/Rojin91/Computer-Network-Assignments/blob/main/PAS078BEI041-Sulav_Kandel/Assignment_6/IPV4_fields.jpg)


1) **Version (4 Bits):** 

* Indicates the version of the IP protocol.   
* For IPv4, this value is 4\.  
* Here, 0100= 4= TPv4. For IPv6 its, 0101\.  
    
    
1) **Header Length (4 bits):** 

   * Overall length of the header.   
   * Its value may vary from 0000 (0) \- 1111(15).   
   * The length of the header is calculated by multiplying the value with 4\. For example, In the above fig 0101= 5 \= 20 bytes.  
   * here the minimum length of the header must be 20 byte so starting is always from 0101\.

1) **Differentiated Services (DSCP) (8 bits):**

* Originally intended to specify the priority of the packet.  
* Now used for Differentiated Services (DiffServ) for Quality of Service (QoS).  
* It helps in classifying and managing network traffic.

1) **Total Length (16 bits):**  
* Specifies the entire packet size, including header and data.  
* Maximum value is 65,535 bytes.

1) **Identification (16 bits):**  
* Used for uniquely identifying the fragments of an original IP datagram.  
* It is used to help reassemble fragmented packets.

1) **Flags (3 bits):**

* Control or identify fragments.  
* First bit: reserved (must be 0).  
* Second bit: Don't Fragment (DF).  
* Third bit: More Fragments (MF).  
    
1) **Fragment Offset (13 bits):**

* Indicates the position of the fragment in the original datagram.  
* Measured in 8-byte units.

1) **Time to Live (TTL) (8 bits):**

* Limits the lifespan of the packet, to prevent it from circulating indefinitely.  
* Decremented by each router the packet passes through.  
* Packet is discarded when TTL reaches 0\.

1) **Protocol (8 bits):**  
* Indicates the protocol used in the data portion of the IP datagram.  
* Examples: 1 for ICMP, 6 for TCP, 17 for UDP.

1) **Header Checksum (16 bits):**  
* Used for error-checking of the header.  
* Recalculated at each hop due to changes in TTL.

1) **Source Address (32 bits):**  
* IP address of the sender.

1) **Destination Address (32 bits):**  
* IP address of the recipient.

1) **Options (variable length, optional):**  
* Allows for additional options.  
* Not commonly used.  
* Length must be a multiple of 32 bits.

1) **Padding (variable length, optional):**  
* Ensures the header length is a multiple of 32 bits.


  


  




                             
