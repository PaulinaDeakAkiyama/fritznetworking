# fritznetworking

Simple home-network automation project for a Raspberry Pi 3 connected to a FRITZ!Box. NEW and not done yet, the read me just outlines structure for now.

## Goals

- Run Pi-hole as local DNS filtering and logging.
- Use Pi-hole as the gateway for all home devices.
- Automate timed parental controls via FRITZ!Box API (TR-064).
- Capture lightweight traffic metadata for debugging.
- Store debugging data in a simple SQL database.
- Connect and manage ExpressVPN on the Pi.

## Architecture

1. Devices use Pi-hole on the Raspberry Pi for DNS and as the gateway.
2. Go scheduler service applies time-based parental rules on FRITZ!Box via TR-064.
3. Go sniffer service captures flow metadata and writes to SQL.
4. ExpressVPN runs on the Pi for secure outbound/remote connectivity.

## Components

- Rasberry Pi 3:
    - connected to router with ethernet for speed and reliablity
    - eth0 as upstream, wlan0 as downstream
    - the primary and only DNS resolver
    - primary gateway

- FRITZ!Box integration:
	- Use TR-064 services for host and parental control actions.
	- Main target service: X_AVM-DE_HostFilter (parental controls).

- dnsmasq (runs on Pi):
	- Handles DHCP for managed devices on wlan0 subnet.
	- Hands out Pi as default gateway and DNS server to all connecting devices.

- DNS filtering and logs (Pi-hole):
	- Per-device DNS queries and blocklist enforcement.
	- Domain-level visibility for most parental debugging.

- Firewall:
    - Makes Rasberry Pi the mandatory gateway    
    - DNS lock

- Packet sniffer (Go):
	- Lightweight capture using interface filters.
	- Store only metadata: timestamp, src/dst IP, resolved domain, ASN, organization, dst port, protocol, bytes/packets, geolocation, service label.

- SQL storage:
	- Start with SQLite for low overhead on Raspberry Pi 3.
    - migrate old data to a different device

- VPN (ExpressVPN):
	- Run ExpressVPN client on the Pi.
	- Keep routing rules explicit so DNS + parental controls still behave as intended.

- Health checkup:
    - (only if nessessary) check cpu usage or destination ip for dynamic offloading using linux ip route to manage gateway.
    - check SD card or external storage usage.
    - check SQLite database file size and row count.

## Runtime Services on Pi

- pihole-FTL (Pi-hole)
- fritz-scheduler (Go)
- fritz-sniffer (Go)
- expressvpn daemon/client
- Health checkup

## Repository Structure

Current folders:

- cmd/
	- blockpage/: optional local status/block page app
	- logger/: logging or data export utility
	- scheduler/: timed parental-control runner
	- test.go: scratch/test entrypoint
- internal/
	- fritzapi/: FRITZ!Box client and API logic
	- config.go: app config loader
	- config.yaml: local config values
	- models.go: SQL and domain models
	- store.go: database access layer
	- info.txt: notes
- web/: static files or minimal UI


## Step by step configuration

1. **Physical wiring**
   - Connect Pi Ethernet port to a LAN port on FRITZBox.
   - eth0: upstream toward FRITZBox and internet.
   - wlan0: local Wi-Fi access point for managed devices.

2. **Configure Pi Wi-Fi as access point**
   - Create a new SSID for managed devices (example: HomeKids).
   - Protect with WPA2 password.
   - Managed devices connect to this SSID, not directly to FRITZBox Wi-Fi.

3. **Give wlan0 a static IP**
   - Pi wlan0 IP: 10.10.10.1
   - This is the fixed address that becomes the gateway for managed devices.

4. **Run DHCP on Pi for that subnet**
   - Hand out client IPs in range 10.10.10.x.
   - Set DHCP default gateway = 10.10.10.1 (the Pi).
   - Set DHCP DNS server = Pi-hole on the Pi (10.10.10.1).

5. **Enable routing on Pi**
   - Turn on IPv4 forwarding so traffic can pass between wlan0 and eth0.

6. **Add firewall and NAT rules on Pi**
   - Allow forwarding from wlan0 to eth0.
   - Allow return traffic from eth0 back to wlan0.
   - Masquerade/NAT outbound traffic from wlan0 via eth0.
   - DNS lock rules:
     - Allow DNS only to Pi-hole (10.10.10.1).
     - Block external DNS on port 53.
     - Block DNS-over-TLS on port 853.

7. **FRITZBox role**
   - FRITZBox sees Pi as one LAN client for upstream internet access.
   - Managed devices are behind Pi and hidden from FRITZBox.

8. **Validate**
   - On a managed device, confirm:
     - IP is in Pi subnet (10.10.10.x).
     - Gateway is Pi (10.10.10.1).
     - DNS is Pi-hole.
   - Confirm internet works.
   - Confirm blocked domains fail.
   - Confirm flow logs appear in SQL table.

