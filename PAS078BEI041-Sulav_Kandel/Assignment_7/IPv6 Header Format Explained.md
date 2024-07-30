### **IPv6 Header Format Explained**

**Question 1: How long is an IPv6 header?**

* An IPv6 header is 40 bytes (320 bits) long.

**Question 2: What are the different fields in the header?** The IPv6 header consists of the following fields:

1. **Version (4 bits)**  
1. **Traffic Class (8 bits)**  
1. **Flow Label (20 bits)**  
1. **Payload Length (16 bits)**  
1. **Next Header (8 bits)**  
1. **Hop Limit (8 bits)**  
1. **Source Address (128 bits)**  
1. **Destination Address (128 bits)**

**Question 3: What is the purpose of each header field?**

1. **Version (4 bits)**  
   * Indicates the IP version, which is 6 for IPv6.  
1. **Traffic Class (8 bits)**  
   * Used for QoS (Quality of Service) and priority settings.  
1. **Flow Label (20 bits)**  
   * Identifies traffic flows that require special handling.  
1. **Payload Length (16 bits)**  
   * Specifies the length of the payload (the data) in bytes following the header.  
1. **Next Header (8 bits)**  
   * Identifies the type of header immediately following the IPv6 header, often indicating the transport layer protocol (e.g., TCP, UDP).  
1. **Hop Limit (8 bits)**  
   * Specifies the maximum number of hops a packet can take before being discarded (similar to TTL in IPv4).  
1. **Source Address (128 bits)**  
   * The IPv6 address of the sender.  
1. **Destination Address (128 bits)**  
   * The IPv6 address of the intended recipient.

### 

### 

### 

### **Visual Representation of IPv6 Header**

![IPv6 Header](https://github.com/Rojin91/Computer-Network-Assignments/blob/main/PAS078BEI041-Sulav_Kandel/Assignment_7/IPv6%20Header.jpg)


**Question 4: In wireshark locate an IPv6 packet and discuss the header present.**

![IPv6 packet](https://github.com/Rojin91/Computer-Network-Assignments/blob/main/PAS078BEI041-Sulav_Kandel/Assignment_7/IPv6%20packet.jpg)


