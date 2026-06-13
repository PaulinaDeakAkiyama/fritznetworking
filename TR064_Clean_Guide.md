# FRITZ TR-064 Clean Guide (Beginner Friendly)

This is a simplified reference based on your saved official document.

## 1) What curl --digest means

curl --digest tells curl to use HTTP Digest Authentication.

- It is not an import.
- It is not a library you add to your code.
- It is a command line option for curl.

Digest auth is a challenge-response login method. The password is not sent as plain text in the request body. Instead, the client and server exchange hash-based values.

Should you use it?

- Yes, for quick testing from terminal: very convenient.
- For scripts/apps, native language support is often better.

In PowerShell, you can do the same thing without curl by using Invoke-WebRequest or Invoke-RestMethod with a credential and Authentication Digest (on supported versions).

## 2) Can you do TR-064 without curl?

Yes.

### Option A: PowerShell native

Use Invoke-WebRequest with Digest authentication and SOAP XML body.

### Option B: Your Go project

Use net/http and implement or reuse a Digest auth helper, then POST SOAP envelopes.

### Option C: Manual content-level SOAP authentication

The FRITZ document also supports SOAP header auth (nonce, realm, auth hash). Useful if you must avoid HTTP Digest and follow the SOAP auth flow in section 10.7.

## 3) Core discovery endpoints

TR-064 is discovered dynamically. Do not hardcode everything.

1. SSDP M-SEARCH
   - ST: urn:dslforum-org:device:InternetGatewayDevice:1
2. Read root device description from LOCATION response header
   - Example: http://192.168.178.1:49000/tr64desc.xml
3. Parse service list from XML
   - For each service, read:
     - serviceType
     - controlURL
     - SCPDURL

Manual discovery mentioned in the FRITZ document:

- http://192.168.178.1:49000/tr64desc.xml
- https://192.168.178.1:49443/tr64desc.xml

Useful action for secure port:

- DeviceInfo:GetSecurityPort (returns HTTPS TR-064 port, commonly 49443)

## 4) Authentication methods in FRITZ TR-064

## 4.1 Default: HTTP Digest Authentication

Use configured FRITZ username + password.

Common behavior:

- 401 Unauthorized: not authenticated
- 606 Action not authorized: authenticated but missing rights
- 866/867/868: second-factor flow required/blocked/busy for protected actions

Notes from doc:

- Too many unauthenticated tries can return 503 instead of 401.
- Some actions are available without authentication (rights marked with -).

## 4.2 Content-Level SOAP Authentication (optional)

SOAP challenge-response in headers.

Digest formula from document:

- secret = MD5(uid:realm:pwd)
- response = MD5(secret:nonce)

Flow:

1. Send InitChallenge header
2. Receive Challenge with Nonce + Realm
3. Send ClientAuth with computed Auth hash
4. Receive NextChallenge and normal SOAP response

## 4.3 URL Session ID for URL-based resources

For URL GET/POST resources (phonebook images, fax, TAM content), use:

- DeviceConfig:X_AVM-DE_CreateUrlSID

## 5) Transactions (important for config changes)

To avoid parallel-write conflicts between clients:

1. DeviceConfig:X_GenerateUUID
2. DeviceConfig:ConfigurationStarted
3. One or more protected actions
4. DeviceConfig:ConfigurationFinished

Timeout: transaction expires after ~45 seconds if not finished.

## 6) Endpoint model (how to think about all endpoints)

TR-064 does not use one single REST-style endpoint list. Instead:

- One control endpoint per service instance (controlURL)
- SOAPAction selects the action within that service

Pattern:

- POST http://<box-ip>:<port><controlURL>
- Header SOAPAction: "<serviceType>#<ActionName>"

Example:

- URL: /upnp/control/deviceinfo
- SOAPAction: urn:dslforum-org:service:DeviceInfo:1#GetSecurityPort

## 7) Complete service catalog from the document

These are the services listed in section 6.1.x. Treat these as your possible TR-064 service endpoints (actual controlURL comes from tr64desc.xml).

1. DeviceInfo
2. DeviceConfig
3. Layer3Forwarding
4. LANConfigSecurity
5. ManagementServer
6. Time
7. UserInterface
8. LANHostConfigManagement
9. LANEthernetInterfaceConfig
10. Hosts
11. WANCommonInterfaceConfig
12. WANDSLInterfaceConfig
13. WANDSLLinkConfig
14. WANEthernetLinkConfig
15. WANPPPConnection
16. WANIPConnection
17. WLANConfiguration
18. X_VoIP
19. X_AVM-DE_Storage
20. X_AVM-DE_WebDAVClient
21. X_AVM-DE_UPnP
22. X_AVM-DE_OnTel
23. X_AVM-DE_TAM
24. X_AVM-DE_RemoteAccess
25. X_AVM-DE_MyFritz
26. X_AVM-DE_Speedtest
27. X_AVM-DE_AppSetup
28. X_AVM-DE_Homeplug
29. X_AVM-DE_Homeauto
30. X_AVM-DE_Dect
31. X_AVM-DE_Filelinks
32. X_AVM-DE_USPController
33. X_AVM-DE_Auth
34. X_AVM-DE_HostFilter
35. X_AVM-DE_Media
36. X_AVM-DE_WANMobileConnection
37. X_AVM-DE_WANFiber

## 8) Device discovery actions (most useful for beginners)

Primary service: Hosts

- GetHostNumberOfEntries
- GetGenericHostEntry
- GetSpecificHostEntry
- X_AVM-DE_GetHostListPath
- X_AVM-DE_GetMeshListPath

Typical approach:

1. Call GetHostNumberOfEntries
2. Loop index 0..N-1 with GetGenericHostEntry
3. Collect IP, MAC, host name, active status

## 9) Rights model quick reference

Rights symbols from the document:

- C: Configuration
- A: Restricted config for apps
- P: Phone
- N: NAS
- H: Home automation
- -: No rights

If action needs rights and user lacks them, call fails with 606.

## 10) Practical recommendations

1. Use local router IP directly if fritz.box DNS is wrong in your setup.
2. Discover services dynamically from tr64desc.xml, do not hardcode control URLs.
3. Start with HTTP Digest auth, then move to content-level auth only if needed.
4. For inventory of connected devices, start with Hosts service actions.
5. For config writes, use transactions.

## 11) Minimal test checklist

1. Ping router IP
2. Fetch tr64desc.xml
3. Call DeviceInfo:GetSecurityPort
4. Authenticate with Digest on a protected action
5. Enumerate hosts via Hosts service

---

If you want, next step can be a second file with copy-paste PowerShell examples for each major action (discover, auth test, list hosts, and parse XML output).
