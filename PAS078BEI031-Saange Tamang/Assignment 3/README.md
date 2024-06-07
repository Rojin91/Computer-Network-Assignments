Task Involved
- Simulating DHPC server on packet tracer
- Tracking the DHPC packets on your local network using WIRESHARK

WIRESHARK
Objective: 
-Capturing DHCP packets to analyze how devices on our network obtain IP addresses.

1. Initial setup
-Opening Wireshark and selecting the correct network interface.

2. Filtering Traffic
-Applying the bootp /dhcp filter to isolate DHCP packets from other network traffic.

2. Capturing Traffic
-Showingthe live capture process and explain how to trigger DHCP events  to generate relevant traffic.

{ At background 1) Windows(Powershell)
                - ipconfig /release
                - ipconfig /renew
                - ipconfig /release
                2) MacOS(zsh)
                - sudo ipconfig set en0 BOOTP
                  sudo ipconfig set en0 DHCP
}

4. Stopping and Saving:
-Detailing the steps to stop the capture and saving the results for later analysis.
