FRITZ! TR-064 - First Steps © FRITZ! GMBH
FRITZ! TR-064 – First Steps
Supported by FRITZ!
Author: FRITZ! GmbH
Date: 2025-08-06
Content
 1 Introduction.......................................................................................................................9
 2 Device Discovery..............................................................................................................9
 2.1.1 Remark...............................................................................................................9
 3 Service Discovery.............................................................................................................9
 3.1 Manual Discovery......................................................................................................9
 4 Security.............................................................................................................................9
 4.1 Authentication..........................................................................................................10
 4.1.1 Remarks...........................................................................................................10
 4.1.2 Content Level Authentication...........................................................................10
 4.2 Encryption................................................................................................................11
 4.2.1 Remark.............................................................................................................11
 4.3 Two-Factor-Authentication......................................................................................12
 5 Discovery and Communication.......................................................................................12
 5.1 Device Description...................................................................................................13
 6 Actions.............................................................................................................................14
 6.1 Actions and User Rights..........................................................................................14
 6.1.1 Service DeviceInfo...........................................................................................16
 6.1.2 Service DeviceConfig.......................................................................................16
 6.1.3 Service Layer3Forwarding...............................................................................16
 6.1.4 Service LANConfigSecurity..............................................................................17
 6.1.5 Service ManagementServer.............................................................................17
 6.1.6 Service Time.....................................................................................................17
 6.1.7 Service UserInterface.......................................................................................18
 6.1.8 Service LANHostConfigManagement..............................................................18
 6.1.9 Service LANEthernetInterfaceConfig...............................................................18
 6.1.10 Service Hosts.................................................................................................19
 6.1.11 Service WANCommonInterfaceConfig...........................................................19
 6.1.12 Service WANDSLInterfaceConfig..................................................................20
 6.1.13 Service WANDSLLinkConfig..........................................................................20
 6.1.14 Service WANEthernetLinkConfig...................................................................20
 6.1.15 Service WANPPPConnection........................................................................20
 6.1.16 Service WANIPConnection............................................................................21
 6.1.17 Service WLANConfiguration...........................................................................22
 6.1.18 Service X_VoIP..............................................................................................23
 6.1.19 Service X_AVM-DE_Storage.........................................................................24
 6.1.20 Service X_AVM-DE_WebDAVClient..............................................................24
 6.1.21 Service X_AVM-DE_UPnP.............................................................................25
 6.1.22 Service X_AVM-DE_OnTel............................................................................25
Version: 59 1/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 6.1.23 Service X_AVM-DE_TAM..............................................................................26
 6.1.24 Service X_AVM-DE_RemoteAccess..............................................................26
 6.1.25 Service X_AVM-DE_MyFritz..........................................................................26
 6.1.26 Service X_AVM-DE_Speedtest......................................................................27
 6.1.27 Service X_AVM-DE_AppSetup......................................................................27
 6.1.28 Service X_AVM-DE_Homeplug.....................................................................27
 6.1.29 Service X_AVM-DE_Homeauto.....................................................................27
 6.1.30 Service X_AVM-DE_Dect...............................................................................28
 6.1.31 Service X_AVM-DE_Filelinks.........................................................................28
 6.1.32 Service X_AVM-DE_USPController...............................................................28
 6.1.33 Service X_AVM-DE_Auth...............................................................................28
 6.1.34 Service X_AVM-DE_HostFilter......................................................................29
 6.1.35 Service X_AVM-DE_Media............................................................................29
 6.1.36 Service X_AVM-DE_WANMobileConnection................................................29
 6.1.37 Service X_AVM-DE_WANFiber.....................................................................29
 7 State Variables................................................................................................................30
 7.1 Compatibility with clients.........................................................................................30
 8 Support for Multiple Clients and Parallel Access............................................................30
 9 Transactions....................................................................................................................31
 9.1 Usage of Transactions.............................................................................................31
 10 Examples......................................................................................................................31
 10.1 SSDP request from application.............................................................................31
 10.2 SSDP response from CPE....................................................................................31
 10.3 Service description request from application.........................................................32
 10.4 Service description response from CPE................................................................32
 10.5 SOAP Request with Action GetSecurityPort.........................................................33
 10.6 SOAP Response for Action GetSecurityPort.........................................................34
 10.7 Content Level Authentication.................................................................................35
 10.7.1 Initial Client Request......................................................................................35
 10.7.2 Server Response to Initial Client Request.....................................................35
 10.7.3 Client Request with Authentication................................................................36
 10.7.4 Server Response for successful Authentication............................................36
 11 Appendix.......................................................................................................................37
 11.1 References............................................................................................................37
Version: 59 2/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
History
Date Version Changes
2010-06-17 1 Initial version
2010-06-18 2 Minor changes
2010-06-21 3 4.1 Authorization, actions GetSecurityPort and DialNumber
2010-07-09 4 Added chapter 6 Actions
2010-08-31 5 Added chapter 7 State Variables
Added reference to UPnP specification
2011-02-25 6 Added chapter Transactions
2011-03-03 7 Added chapter Support for Multiple Clients and Parallel Access
with transactions or multiple session IDs
2011-12-05 8 Changed chapter Authorization to Authentication
Added sub chapter Content Level Authentication
2011-12-21 9 Changed examples for Content Level Authentication
2012-11-06 10 Added chapter 6.1 Actions and User Rights .
Change list of action allowed without authentication in chapter 4.1
Authentication .
2012-11-07 11 LanConfigSecurity:GetInfoEx () without authentication
ManagementServer:GetInfo () without authentication
Layer3Forwarding:SetDefaultConnectionService () with
configuration rights.
2012-12-03 12 Description of authentication for multi user configurations.
X_AVM-DE_GetAnonymousLogin ()
X_AVM-DE_GetCurrentUser ()
2013-01-23 13 Changed chapter 4.1.1 Remarks to clarify service behaviour for
unauthorized requests.
2013-03-25 14 Details for content level authentication.
2013-04-02 15 Spelling
Added actions in chapter 6.1.25 Service X_AVM-DE_MyFritz .
2013-08-20 16 Added action and changed right for all actions in chapter 6.1.10
Service Hosts .
2013-08-23 17 Changed right for GetInfo in chapter 6.1.25 Service X_AVMDE_MyFritz .
2013-08-26 18 Add action in chapter
6.1.17 Service WLANConfiguration :
X_AVM-DE_GetWLANExtInfo
Corrections 6.1.17 Service WLANConfiguration :
X_AVM_DE_SetStickSurfEnable → X_AVMVersion: 59 3/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
DE_SetStickSurfEnable
Corrections 6.1.18 Service X_VoIP :
X_AVM_DE_*VoIPAccount → X_AVM-DE_*VoIPAccount
X_AVM_DE_GetClient* → X_AVM-DE_GetClient*
X_AVM_DE_SetClient* → X_AVM-DE_SetClient*
X_AVM_DE_DeleteClient → X_AVM-DE_DeleteClient
X_AVM_DE_GetClients → X_AVM-DE_GetClients
X_AVM_DE_GetNumbers → X_AVM-DE_GetNumbers
X_AVM_DE_GetNumberOfNumbers → X_AVMDE_GetNumberOfNumbers
2014-05-26 19 Added actions to service Hosts in chapter 6.1.10 Service Hosts :
X_AVM-DE_GetAutoWakeOnLANByMACAddress,
X_AVM-DE_SetAutoWakeOnLANByMACAddress,
X_AVM-DE_SetHostNameByMACAddress,
X_AVM-DE_WakeOnLANByMACAddress
2015-02-06 20 Added service X_AVM-DE_Speedtest in chapter 6.1.26 Service
X_AVM-DE_Speedtest .
2015-03-17 21 Added missing services and actions to chapter 6.1 Actions and
User Rights . Change needed rights for many actions.
Sort actions per service in alphabetical order.
2015-06-02 22 UserInterface:GetInfo change Right C => C or A
Added Service: X_AVME_DE_Homeauto in chapter 6.1 Actions
and User Rights
2015-06-05 23 Added New Service: X_AVME_DE_Dect in chapter 6.1 Actions
and User Rights
2015-06-10 24 Added missing actions
UserInterface:X_AVM_DE_GetInfo
UserInterface:X_AVM_DE_SetConfig
2015-07-17 25 Added missing action
Service X_AVM-DE_AppSetup: SetAppMessageReceiver
2015-08-31 26 Change rights for UserInterface:GetInfo
2015-10-01 27 Add action SetFTPServerWAN to service 6.1.19 Service X_AVMDE_Storage .
2015-11-19 28 Corrected actions for service
DeviceConfig
Layer3Forwarding: Entries => Entry
UserInterface: X_AVM-DE_DoPrepareCGI
WANCommonInterfaceConfig: X_AVM-DE_GetOnlineMonitor
WLANConfiguration: X_AVM-DE_WLANGetNightControl =>
X_AVM-DE_GetNightControl
X_AVM-DE_OnTel: GetDeflection, GetDeflections,
GetInfoByIndex, GetNumberOfDeflections
X_AVM-DE_Filelink added
Version: 59 4/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
X_AVM-DE_Homeplug
X_AVM-DE_StorageService: SetUserConfig
2016-06-29 29 Changed required rights for action GetDDNSInfo
2017-01-09 30 Add description to chapter 4.1.2 Content Level Authentication .
Change several actions in chapter 6.1 Actions and User Rights :
Add - Hosts:X_AVM-DE_GetHostListPath
Add - Hosts:X_AVM-DE_GetSpecificHostEntyByIP
Add - WLANConfiguration:X_AVMDE_GetSpecificAssociatedDeviceInfoByIp
Add - X_AVM-DE_AppSetup:ResetEvent
Add - X_AVM-DE_FileLinks:GetFileLinkListPath
Add - X_AVM-DE_Storage:RequestFTPServerWAN
Add - X_AVM-DE_TAM:GetList
Add - X_VoIP:X_AVM-DE_GetClientByClientId
Add - X_VoIP:X_AVM-DE_SetClient4
Chg - Time:GetInfo
Chg - X_AVM-DE_AppSetup:GetInfo
Chg - X_AVM-DE_AppSetup:RegisterApp
Del - WANIPConnection:X_SetDNSServers
Del – WANPPPConnection:X_SetDNSServers
2017-06-02 31 Add PhoneRight and HomeautoRight for Actions
GetNumberOfDectEntries, GetGenericDectEntry,
GetSpecificDectEntry
2017-06-06 32 Change rights for Actions
- SetEnable
- X_AVM-DE_GetWPSInfo
- X_AVM-DE_SetWPSInfo
in WLANConfiguration
2017-11-22 33 Add Action SetEnable in Service RemoteAccess
2018-08-15 34 Updating to state of other tr064-service documents and adding
some Two-Factor-Authentication information. Changed referred IP
address. Added table for required rights for the service X_AVMDE_Auth. LanConfigSecurity:GetInfo need authentication and is
removed from list in section 4.1.
Add X_VoIP:X_AVM-DE_GetVoIPCommonAreaCode
Add X_VoIP:X_AVM-DE_GetVoIPCommonCountryCode
Add X_VoIP:X_AVM-DE_SetDelayedCallNotification
Add X_VoIP:X_AVM-DE_SetVoIPCommonAreaCode
Add X_VoIP:X_AVM-DE_SetVoIPCommonCountryCode
Add WANDSLInterfaceConfig:X_AVM-DE_GetDSLDiagnoseInfo
Add Hosts:X_AVM-DE_GetMeshListPath
Add X_AVM-DE_OnTel:GetPhonebookEntryUID
Version: 59 5/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
Add X_AVM-DE_OnTel:DeletePhonebookEntryUID
Add X_AVM-DE_AppSetup:GetAppMessageFilter
Add X_AVM-DE_AppSetup:SetAppMessageFilter
Add X_AVM-DE_AppSetup:SetAppVPNwithPFS
Add X_AVM-DE_Homeauto:GetInfo
Add X_AVM-DE_Dect:GetDectListPath
Add X_AVM-DE_Auth:GetInfo
Add X_AVM-DE_Auth:GetState
Add X_AVM-DE_Auth:SetConfig
Add WLANConfiguration:X_AVM-DE_SetWLANGlobalEnable
Del X_AVM-DE_GetGenericHostEntryExt
Del X_AVM-DE_GetSpecificHostEntryExt
2019-01-07 35 Changed needed rights for
WLANConfiguration:GetWLANExtInfo
X_AVM-DE_Storage:GetInfo
2019-01-15 36 Changed needed rights for UserInterface:X_AVM-DE_DoUpdate
2019-02-08 37 Add Service X_AVM-DE_AppSetup:GetAppRemoteInfo
2019-03-01 38 Changed needed rights for actions of WLANConfiguration service:
GetBeaconType
GetSecurityKeys
GetSSID
SetSecurityKeys
X_AVM-DE_GetWPSInfo
X_AVM-DE_SetWPSConfig
2019-12-12 39 New Service X_AVM-DE_HostFilter added.
2020-01-09 40 New Action added WLANConfiguration:X_AVMDE_GetWLANConnectionInfo
2020-02-27 41 New action added LANConfigSecurity:X_AVM-DE_GetUserList
2020-06-23 42 New actions added
X_AVM-DE_Hostfilter:DisallowWANAccessByIP,
X_AVM-DE_Hostfilter:GetWANAccessByIP
X_AVM-DE_OnTel:GetCallBarringEntry
X_AVM-DE_OnTel:GetCallBarringEntryByNum
X_AVM-DE_OnTel:GetCallBarringList
X_AVM-DE_OnTel:SetCallBarringEntry
X_AVM-DE_OnTel:DeleteCallBarringEntryUID
WANDSLInterfaceConfig:X_AVM-DE_GetDSLInfo
X_VoiP:X_AVM-DE_GetAlarmClock
X_VoiP:X_AVM-DE_GetNumberOfAlarmClocks
X_VoiP:X_AVM-DE_SetAlarmClockEnable
2021-01-20 43 New actions added DeviceConfig:X_AVMDE_GetSupportDataInfo, DeviceConfig:X_AVMDE_SendSupportData
Version: 59 6/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
2021-07-07 44 New actions added
WANCommonInterfaceConfig:WanCommonInterfaceConfig_GetAc
tiveProvider, X_VoIP:X_AVM-DE_GetVoIPStatus, Hosts:X_AVMDE_SetPrioritizationByIP
Changed needed rights for action of WANCommonInterfaceConfig:
GetCommonLinkProperties
2021-07-12 45 Added a new chapter describing the device description and some
contents of the UpnP root device description
2021-09-29 46 New action added DeviceConfig:X_AVM-DE_SetSupportDataEnable
New user rights for Hosts:X_AVM-DE_SetPrioritizationByIP: App, Phone
2021-11-22 47 New user rights for
DeviceConfig: X_AVM-DE_GetSupportDataEnable
Hosts: X_AVM-DE_GetFriendlyName and X_AVM-DE_SetFriendlyName
X_VoIP: X_AVM-DE_GetVoIPAccounts
X_AVM-DE_OnTel: DeleteDeflection and SetDeflection
X_AVM-DE_RemoteAccess: SetLetsEncryptEnable
X_AVM-DE_MyFritz: SetMyFRITZ
2022-01-19 48 New user rights for
X_AVM-DE_Speedtest: GetStatistics and ResetStatistics
2022-02-22 49 Changed rights for UserInterface:X_AVM-DE_GetInfo, Hosts:X_AVMDE_GetAutoWakeOnLANByMACAddress and X_AVMDE_SetAutoWakeOnLANByMACAddress
Required rights explained in subdocuments
2022-02-22 50 Added new service X_AVM-DE_Media with actions and user rights
2022-06-13 51 Added new service X_AVM-DE_USPController with actions and user
rights
2022-10-13 52 Added new service X_AVM-DE_WANMobileConnection with actions and
user rights
Added actions X_AVM-DE_GetInfo, X_AVM-DE_SetFriendlyNameByIP
and X_AVM-DE_SetFriendlyNameByMAC for service Hosts
Added actions X_AVM-DE_GetWLANDeviceListPath, and X_AVMDE_SetWPSEnable for service WLANConfiguration
Added action GetBoxSenderId for service X_AVM-DE_AppSetup
Marked action SetDefaultWEPKeyIndex and GetDefaultWEPKeyIndex
as removed
2022-11-10 53 Added new Actions GetAccessTechnology, SetAccessTechnology,
GetBandCapabilities, GetEnabledBandCapabilities and
SetEnabledBandCapabilities for Service X_AVMDE_WANMobileConnection
2022-11-18 54 Added new actions GetPreferredAccessTechnology and
SetPreferredAccessTechnology for Service X_AVMDE_WANMobileConnection
2023-09-11 55 Added new service X_AVM-DE_WANFiber with actions and user rights
Added actions GetInfo, GetInfoGPON and GetStatistics for service
X_AVM-DE_WANFiber
Version: 59 7/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
2024-02-15 56 Added action X_AVM-DE_GetDeviceLogPath for service DeviceInfo
2024-12-19 57 Added actions X_AVM-DE_SetDeviceClassUserByIP and X_AVMDE_SetDeviceClassUserByMAC for service Hosts
2025-03-25 58 Added actions GetHostEntryByIP, GetFilterProfileByID,
AddTicketTimeToHostEntryByIP, AddHostEntryToFilterProfile and
GetFilterProfiles to X_AVM-DE_HostFilter
2025-08-06 59 Changed AVM to FRITZ!
Version: 59 8/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 1 Introduction
This document gives at short overview about the TR-064 protocol implementation in
FRITZ! devices. Especially the discovery of IP addresses, ports, TR-064 devices and
services.
Next chapter explains the usage of the security options.
In addition some example messages between the management application and the CPE
are listed to see some details.
For details please refer the TR-064 document at
http://www.broadband-forum.org/technical/download/TR-064.pdf.
 2 Device Discovery
To be able to communicate with a TR-064 device from FRITZ!, the following steps have to
be done.
A complete description can be found in chapter 3.2.2 SSDP Search on Start-up of CPE
Management Application in the [TR064] specification.
 1 SSDP Discovery using M-SEARCH
 2 Retrieval of device description file(s) using HTTP GET
 2.1 The content of the SSDP discover response message has to be
parsed to get the valid URL for downloading the service description.
 3 Be aware of the lease time for every announcement. If an announcement is not resent prior to the expire of the lease, the service discovery has to be restarted by the
application.
 2.1.1 Remark
In the [TR064] specification chapter 3.2.1 SSDP Advertisement on Start-up of CPE
Device, the CPE MUST send broadcast SSDP advertisement (NOTIFY) messages.
FRITZ! devices do not do this, to avoid being shown in Microsoft Windows UPnP/ network
devices list by default.
 3 Service Discovery
For the individual services several separate description files can be retrieved.
The URLs are given in the central service description.
An URL can be combined by taking the IP address and the port from the SSDP discovery
response message (LOCATION header) and the value in a XML tag SCPDURL.
 3.1 Manual Discovery
If a TR-064 client tries to detect an FRITZ! TR-064 service manually, the following steps
can be used.
• DNS resolution of fritz.box
• HTTP GET request to http://192.168.178.1:49000/tr64desc.xml
• HTTPS GET request to https://192.168.178.1:49443/tr64desc.xml
 4 Security
TR-064 offers authorization and encryption to make safe communication. Furthermore
some actions with potential to be the target of fraud attacks are protected via Two-FactorAuthentication.
Version: 59 9/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 4.1 Authentication
In the [TR064] specification, chapter 4.2 Authentication, the default usernames and factory
default passwords are described.
FRITZ! CPEs have the following implementation.
• If an individual username and a password is configured via WebGUI or TR-064, a
configured username and password have to be used for password-protected
actions.
It is recommended to always use a configured username and password.
The following actions do not need authorization:
◦ DeviceInfo:GetSecurityPort
◦ WLANConfiguration:X_SetHighFrequencyBand
◦ LanConfigSecurity:X_AVM-DE_GetAnonymousLogin
• If only a password is configured via WebGUI or TR-064, a matching username is
used internally.
◦ It is not recommended to use only a password without a configured username.
◦ If no configured username is used the following action shows which internal user
is used instead:
LanConfigSecurity:X_AVM-DE_GetCurrentUser
• If no individual password is configured, TR-064 does not require a password or
username.
The following action is protected and not available without an individual password:
◦ DialNumber
The default authentication mechanism is HTTP authentication using digest (MD5 hashes).
Alternatively an authentication on content level is supported.
 4.1.1 Remarks
The username dslf-reset is not supported.
The TR-064 service may use HTTP status code 503 (service unavailable) instead of 401
(unauthorized) if too many unauthenticated requests from one client are made.
 4.1.2 Content Level Authentication
In addition to HTTP digest authentication which is supported by many libraries another
alternative authentication is available.
The Content Level Authentication is based on the idea of SOAP extension for basic and
digest authentication described in [SOAPAUTH01].
The client may either use the HTTP authentication mechanisms or use the content level
authentication for SOAP actions that require authorization.
To access URLs with HTTP GET or HTTP POST a TR-064 URL session ID has to be
retrieved first.
The TR-064 URL session ID is part of the URLs for phone book and call list.
For further requests like:
• phone book images,
• fax messages or
• answering machine messages
the SOAP action DeviceConfig::X_AVM-DE_CreateUrlSID () has to be invoked.
See the example for the authentication in the chapter 10.7 Content Level Authentication .
Version: 59 10/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
The client has to insert a SOAP header in the content to initiate the content level
authentication.
The server responses an initial request with content level authentication contains a nonce
value, the realm and a SOAP fault with a number and a text.
The client needs the nonce value, username, realm and password to calculate the digest
hash.
A client side nonce is supported as optional parameter, too.
The following calculation and example values are taken from the examples in chapter
10.7 Content Level Authentication .
Variable Value Remark
uid admin
pwd gurkensalat
realm F!Box SOAP-Auth Not writeable.
sn F758BE72FB999CEA
response b4f67585f22b0af7c4615db5a18faa14
The calculation is similar to chapter 4.3 Digest Computation in [SOAPAUTH01].
response = MD5( concat(secret, ":", sn) )
when only client authentication is required.
Where:
• secret (string) represents the hex encoding of the digest
secret = MD5( concat(uid, ":", realm, ":", pwd) )
• uid (string) represents the User ID.
• realm (string) represents the Realm name.
• pwd (string) represents the password.
• sn (string) represents the nonce serving as the challenge from the server to the
client.
 4.2 Encryption
The TR-064 specification chapter 4.4 Encryption describes the possibilities for encryption.
To make it more easy for applications the port can be retrieved using a special SOAP
action.
The service type urn:dslforum-org:service:DeviceInfo:1 offers the action GetSecurityPort
where the used port value can be retrieved.
Example messages for a request and a response with the security port can be found in the
examples chapter.
 4.2.1 Remark
In the [TR064] specification, SSL encryption is recommended only for some SOAP
actions. The support for access over an encrypted HTTPS link is not specified in details.
Therefore FRITZ! decided to use the explained action GetSecurityPort.
Version: 59 11/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 4.3 Two-Factor-Authentication
Some actions are protected via Two-Factor-Authentication (2FA) to prevent remote config
changes. Even when credentials are known, the physical presence is needed to confirm
changes on the FRITZ! CPE device configuration. For more information see X_Auth
service documentation.
 5 Discovery and Communication
Base on the message flow in the [TR064] specification, the following interaction sequence
is expected by the CPE.
Version: 59 12/37 2025-08-06
Figure 1: CPE Discovery and Communications base on the
[TR064] specification
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 5.1 Device Description
The TR-064 root device description on FRITZ!OS devices is delivered as response of search
request with M-SEARCH.
A typical response will contain the following URL:
 http://192.168.1 78 .1:49000/tr64desc.xml
The UPnP Device Architecture 1.1 and 2.0 describe the protocol in details.
A client CAN (re-)identify a FRITZ!OS device by its MAC address.
The MAC address is part of the UUID and serialNumber.
Some more contents of the UPnP root device description can be found in the table below.
XML Tag Description
UDN UUID with default MAC as last part, to used for (re-) identification
of a FRITZ!OS device by its MAC
serialNumber Device MAC address
serviceType URN with service name and ID
serviceId URN with service ID
controlURL Path to be used for TR-064 SOAP actions
SCPDURL URL path to TR-064 service SCPD
UUID = Universally Unique Identifier
URN = Uniform Resource Name
 6 Actions
Actions may change in their lifetime. The names of actions are kept unchanged as long as
possible, but the arguments may change if needed.
The number of arguments may differ and the meaning of a state variable may change.
To detect such a change, it is recommended to parse all SCPDXML files after attaching to
a device.
In general actions return 402 if the number of arguments for an action is not as expected
or an unexpected argument is used.
If an unknown action is used the returned code is 401. This return code is used for
obsoleted actions, too.
 6.1 Actions and User Rights
The following user rights are supported:
• no rights
• configuration
• phone contents and phone configuration
• NAS contents and NAS configuration
User rights for configuration include user rights for phone contents and phone
configuration.
The user rights may be limited to LAN or granted to LAN and WAN access.
There are some more settings for user rights:
• access from LAN only with authentication
• access from LAN without authentication
Version: 59 13/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
• no configured user and no configured password
Every action has needed user rights. The following user rights mappings to SOAP actions
shall give an overview. Every action without explicit “required rights” field in subdocuments
needs config rights.
C – Configuration
A – restricted configuration for applications
P – Phone contents and phone configuration
N – NAS contents and NAS configuration
H – Home automation contents and configuration
- – No rights
If a user is not authenticated, 401 (“Unauthorized”) will be returned.
If a user is authenticated but has not the needed rights, 606 (“Action not authorized”) will
be returned. If an action needs 2FA, the status code 866 (“second factor authentication
required”), 867 (“second factor authentication blocked”) or 868 (“second factor
authentication busy”) will be returned and the 2FA procedure, as described in X_Auth
service description, has to be passed.
Version: 59 14/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 6.1.1 Service DeviceInfo
Action User Rights
GetDeviceLog C
GetInfo C
GetSecurityPort -
SetProvisioningCode C
X_AVM-DE_GetDeviceLogPath C
 6.1.2 Service DeviceConfig
Action User Rights
ConfigurationFinished C
ConfigurationStarted C
FactoryReset C
GetPersistentData C
Reboot C
SetPersistentData C
X_AVM-DE_CreateUrlSID C or A or P or N or H
X_AVM-DE_GetConfigFile C
X_AVM-DE_SetConfigFile C
X_GenerateUUID C
X_AVM-DE_GetSupportDataInfo C or A or P or N or H
X_AVM-DE_SendSupportData C or A or P or N or H
X_AVM-DE_GetSupportDataEnable -
X_AVM-DE_SetSupportDataEnable C
 6.1.3 Service Layer3Forwarding
Action User Rights
AddForwardingEntry C
DeleteForwardingEntry C
GetDefaultConnectionService C
GetForwardNumberOfEntries C
GetGenericForwardingEntry C
GetSpecificForwardingEntry C
SetDefaultConnectionService C
SetForwardingEntryEnabled C
Version: 59 15/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 6.1.4 Service LANConfigSecurity
Action User Rights
GetInfo -
SetConfigPassword C or A or P or N or H
X_AVM-DE_GetAnonymousLogin -
X_AVM-DE_GetCurrentUser C or A or P or N or H
X_AVM-DE_GetUserList -
 6.1.5 Service ManagementServer
Action User Rights
GetInfo C
SetConnectionRequestAuthentication C
SetManagementServerPassword C
SetManagementServerURL C
SetManagementServerUsername C
SetPeriodicInform C
SetUpgradeManagement C
X_AVM-DE_GetTR069FirmwareDownloadEnabled C
X_AVM-DE_SetTR069FirmwareDownloadEnabled C
X_SetTR069Enable C
 6.1.6 Service Time
Action User Rights
GetInfo C or A or P or N or H
SetNTPServers C
Version: 59 16/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 6.1.7 Service UserInterface
Action User Rights
GetInfo -
X_AVM-DE_CheckUpdate C
X_AVM-DE_DoUpdate C or -, depending on configuration
X_AVM-DE_DoPrepareCGI C
X_AVM-DE_DoManualUpdate C
X_AVM-DE_GetInternationalConfig C
X_AVM-DE_SetInternationalConfig C
X_AVM-DE_GetInfo -
X_AVM-DE_SetConfig C
 6.1.8 Service LANHostConfigManagement
Action User Rights
GetAddressRange C
GetDNSServers C
GetInfo C
GetIPInterfaceNumberOfEntries C
GetIPRoutersList C
GetSubnetMask C
SetAddressRange C
SetDHCPServerEnable C
SetIPInterface C
SetIPRouter C
SetSubnetMask C
 6.1.9 Service LANEthernetInterfaceConfig
Action User Rights
GetInfo C
GetStatistics C
SetEnable C
Version: 59 17/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 6.1.10 Service Hosts
Action User Rights
GetGenericHostEntry -
GetHostNumberOfEntries -
GetSpecificHostEntry -
X_AVM-DE_GetAutoWakeOnLANByMACAddress -
X_AVM-DE_GetChangeCounter -
X_AVM-DE_GetFriendlyName -
X_AVM-DE_GetHostListPath C or A or P
X_AVM-DE_GetInfo -
X_AVM-DE_GetMeshListPath C or P or A
X_AVM-DE_GetSpecificHostEntyByIP C or A or P
X_AVM-DE_HostDoUpdate C or A
X_AVM-DE_HostsCheckUpdate C or A
X_AVM-DE_SetAutoWakeOnLANByMACAddress -
X_AVM-DE_SetDeviceClassUserByIP C or A
X_AVM-DE_SetDeviceClassUserByMAC C or A
X_AVM-DE_SetFriendlyName C or A or P or N or H
X_AVM-DE_SetFriendlyNameByIP C or A
X_AVM-DE_SetFriendlyNameByMAC C or A
X_AVM-DE_SetHostNameByMACAddress C
X_AVM-DE_SetPrioritizationByIP C or A or P
X_AVM-DE_WakeOnLANByMACAddress -
 6.1.11 Service WANCommonInterfaceConfig
Action User Rights
GetCommonLinkProperties C or P or A or N or H
GetTotalBytesReceived C
GetTotalBytesSent C
GetTotalPacketsReceived C
GetTotalPacketsSent C
X_AVM-DE_GetActiveProvider C or P or A
X_AVM-DE_SetWANAccessType C
X_AVM-DE_GetOnlineMonitor C
Version: 59 18/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 6.1.12 Service WANDSLInterfaceConfig
Action User Rights
GetInfo C
GetStatisticsTotal C
X_AVM-DE_GetDSLDiagnoseInfo -
X_AVM-DE_GetDSLInfo -
 6.1.13 Service WANDSLLinkConfig
Action User Rights
GetATMEncapsulation C
GetAutoConfig C
GetDestinationAddress C
GetDSLLinkInfo C
GetInfo C
GetStatistics C
SetATMEncapsulation C
SetDestinationAddress C
SetDSLLinkType C
SetEnable C
 6.1.14 Service WANEthernetLinkConfig
Action User Rights
GetEthernetLinkStatus C
 6.1.15 Service WANPPPConnection
Action User Rights
AddPortMapping C
DeletePortMapping C
ForceTermination C
GetConnectionType C
GetConnectionTypeInfo C
GetExternalIPAddress C
GetGenericPortMappingEntry C
GetInfo C
GetLinkLayerMaxBitRates C
Version: 59 19/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
GetNATRSIPStatus C
GetPortMappingNumberOfEntries C
GetSpecificPortMappingEntry C
GetStatusInfo C
GetUserName C
RequestConnection C
SetConnectionTrigger C
SetIdleDisconnectTime C
SetPassword C
SetRouteProtocolRx C
SetUserName C
X_AVM-DE_GetAutoDisconnectTimeSpan C
X_AVM-DE_SetAutoDisconnectTimeSpan C
X_GetDNSServers C
 6.1.16 Service WANIPConnection
Action User Rights
AddPortMapping C
DeletePortMapping C
ForceTermination C
GetConnectionType C
GetConnectionTypeInfo C
GetExternalIPAddress C
GetGenericPortMappingEntry C
GetInfo C
GetNATRSIPStatus C
GetPortMappingNumberOfEntries C
GetSpecificPortMappingEntry C
GetStatusInfo C
RequestConnection C
SetConnectionTrigger C
SetIdleDisconnectTime C
SetRouteProtocolRx C
X_GetDNSServers C
Version: 59 20/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 6.1.17 Service WLANConfiguration
Action User Rights
GetBasBeaconSecurityProperties C
GetBeaconAdvertisement C
GetBeaconType C or A
GetBSSID C
GetChannelInfo C
GetDefaultWEPKeyIndex (removed) C
GetGenericAssociatedDeviceInfo C or A or P
GetInfo C or A or P
GetPacketStatistics C
GetSecurityKeys A
GetSpecificAssociatedDeviceInfo C or A or P
GetSSID A
GetStatistics C
GetTotalAssociations C or A or P
SetBasBeaconSecurityProperties C
SetBeaconAdvertisement C
SetBeaconType C
SetChannel C
SetConfig C
SetDefaultWEPKeyIndex (removed) C
SetEnable C or A
SetSecurityKeys C for normal AP, A for guest AP
SetSSID C
X_AVM-DE_GetIPTVOptimzed C
X_AVM-DE_GetNightControl -
X_AVM-DE_GetSpecificAssociatedDeviceInfoByIp C or A or P
X_AVM-DE_GetWLANConnectionInfo -
X_AVM-DE_GetWLANDeviceListPath C or A or P
X_AVM-DE_GetWLANExtInfo C or A
X_AVM-DE_GetWLANHybridMode C
X_AVM-DE_GetWPSInfo C or A
X_AVM-DE_SetIPTVOptimzed C
Version: 59 21/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
X_AVM-DE_SetStickSurfEnable C
X_AVM-DE_SetWLANGlobalEnable C or A
X_AVM-DE_SetWLANHybridMode C
X_AVM-DE_SetWPSConfig C or A
X_AVM-DE_SetWPSEnable C or A
X_SetHighFrequencyBand (removed) -
 6.1.18 Service X_VoIP
Action User Rights
GetExistingVoIPNumbers P
GetInfo P
GetInfoEx P
GetMaxVoIPNumbers P
GetVoIPCommonAreaCode P
GetVoIPCommonCountryCode P
GetVoIPEnableAreaCode P
GetVoIPEnableCountryCode P
SetConfig P
SetVoIPCommonAreaCode P
SetVoIPCommonCountryCode P
SetVoIPEnableAreaCode P
SetVoIPEnableCountryCode P
X_AVM-DE_AddVoIPAccount P
X_AVM-DE_DeleteClient P
X_AVM-DE_DelVoIPAccount P
X_AVM-DE_DialGetConfig P
X_AVM-DE_DialHangup P
X_AVM-DE_DialNumber P
X_AVM-DE_DialSetConfig P
X_AVM-DE_GetAlarmClock P
X_AVM-DE_GetClient P
X_AVM-DE_GetClient2 P
X_AVM-DE_GetClient3 P
X_AVM-DE_GetClientByClientId A or P
Version: 59 22/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
X_AVM-DE_GetClients P
X_AVM-DE_GetNumberOfAlarmClocks P
X_AVM-DE_GetNumberOfClients P
X_AVM-DE_GetNumberOfNumbers P
X_AVM-DE_GetNumbers P
X_AVM-DE_GetPhonePort P
X_AVM-DE_GetVoIPAccount P
X_AVM-DE_GetVoIPAccounts P or A
X_AVM-DE_GetVoIPCommonAreaCode P
X_AVM-DE_GetVoIPCommonCountryCode P
X_AVM-DE_GetVoIPStatus P or A
X_AVM-DE_SetAlarmClockEnable P
X_AVM-DE_SetClient P
X_AVM-DE_SetClient2 P
X_AVM-DE_SetClient3 P
X_AVM-DE_SetClient4 P
X_AVM-DE_SetDelayedCallNotification P
X_AVM-DE_SetVoIPCommonAreaCode P
X_AVM-DE_SetVoIPCommonCountryCode P
 6.1.19 Service X_AVM-DE_Storage
Action User Rights
GetInfo C or A
GetUserInfo C
RequestFTPServerWAN A or N only but not C
SetFTPServer C
SetFTPServerWAN C
SetSMBServer C
SetUserConfig C
 6.1.20 Service X_AVM-DE_WebDAVClient
Action User Rights
GetInfo C
SetConfig C
Version: 59 23/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 6.1.21 Service X_AVM-DE_UPnP
Action User Rights
GetInfo C
SetConfig C
 6.1.22 Service X_AVM-DE_OnTel
Action User Rights
AddPhonebook P
DeleteByIndex P
DeleteCallBarringEntryUID P
DeleteDeflection P
DeletePhonebook P
DeletePhonebookEntry P
DeletePhonebookEntryUID P
GetCallBarringEntry P
GetCallBarringEntryByNum P
GetCallBarringList P
GetCallList P
GetDECTHandsetInfo P
GetDECTHandsetList P
GetDeflection P
GetDeflections P
GetInfo P
GetInfoByIndex P
GetNumberOfDeflections P
GetNumberOfEntries P
GetPhonebook P
GetPhonebookEntry P
GetPhonebookEntryUID P
GetPhonebookList P
SetCallBarringEntry P
SetConfig P
SetConfigByIndex P
SetDECTHandsetPhonebook P
SetDeflection P
Version: 59 24/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
SetDeflectionEnable P
SetEnable P
SetEnableByIndex P
SetPhonebookEntry P
SetPhonebookEntryUID P
 6.1.23 Service X_AVM-DE_TAM
Action User Rights
DeleteMessage P
GetInfo P
GetList A or P
GetMessageList P
MarkMessage P
SetEnable P
 6.1.24 Service X_AVM-DE_RemoteAccess
Action User Rights
GetDDNSInfo C or A or P or N or H
GetDDNSProviders C
GetInfo C or A or P or N or H
SetConfig C
SetDDNSConfig C
SetEnable C
SetLetsEncryptEnable C or A
 6.1.25 Service X_AVM-DE_MyFritz
Action User Rights
DeleteServiceByIndex C or A or P
GetInfo C or A or P or N or H
GetNumberOfServices C or A or P
GetServiceByIndex C or A or P
SetServiceByIndex C or A or P
SetMyFRITZ C or A
Version: 59 25/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 6.1.26 Service X_AVM-DE_Speedtest
Action User Rights
GetInfo C
GetStatistics C
ResetStatistics C
SetConfig C
 6.1.27 Service X_AVM-DE_AppSetup
Action User Rights
GetAppMessageFilter C or A or P or N or H
GetAppRemoteInfo C or A
GetBoxSenderId C or A or P or N or H
GetConfig C or A or P or N or H
GetInfo -
RegisterApp C or P or N or H but not A
ResetEvent C
SetAppMessageFilter C or A or P or N or H
SetAppMessageReceiver C or A or P or N or H
SetAppVPN C
SetAppVPNwithPFS C
 6.1.28 Service X_AVM-DE_Homeplug
Action User Rights
DeviceDoUpdate C or A
GetGenericDeviceEntry C or A
GetNumberOfDeviceEntries C or A
GetSpecificDeviceEntry C or A
 6.1.29 Service X_AVM-DE_Homeauto
Action User Rights
GetGenericDeviceInfos C or H
GetInfo C or H
GetSpecificDeviceInfos C or H
SetDeviceName C or A
SetSwitch C or H
Version: 59 26/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 6.1.30 Service X_AVM-DE_Dect
Action User Rights
DectDoUpdate C or A
GetDectListPath C or A
GetGenericDectEntry C or A or P or H
GetNumberOfDectEntries C or A or P or H
GetSpecificDectEntry C or A or P or H
 6.1.31 Service X_AVM-DE_Filelinks
Action User Rights
DeleteFilelinkEntry C or N
GetFileLinkListPath C or N
GetGenericFilelinkEntry C or N
GetNumberOfFilelinkEntries C or N
GetSpecificSpecificEntry C or N
NewFilelinkEntry C or N
SetFilelinkEntry C or N
 6.1.32 Service X_AVM-DE_USPController
Action User Rights
AddUSPContoller C
DeleteUSPControllerByIndex C
GetInfo C or A or P or N or H
GetUSPControllerByIndex C
GetUSPContollerNumberOfEntries C
GetUSPMyFRITZEnable C or A
SetUSPControllerEnableByIndex C
SetUSPMyFRITZEnable C or A or P or N or H
 6.1.33 Service X_AVM-DE_Auth
Action User Rights
GetInfo C or A or P or N or H
GetState C or A or P or N or H
SetConfig C or A or P or N or H
Version: 59 27/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 6.1.34 Service X_AVM-DE_HostFilter
Action User Rights
AddHostEntryToFilterProfile C or A
AddTicketTimeToHostEntryByIP C or A
DisallowWANAccessByIP C or A
DiscardAllTickets C or A
GetFilterProfileByID C or A
GetFilterProfiles C or A
GetHostEntryByIP C or A
GetTicketIDStatus C or A
GetWANAccessByIP C or A
MarkTicket C or A
 6.1.35 Service X_AVM-DE_Media
Action User Rights
GetDVBCEnable -
GetInfo -
GetSearchProgress -
SetDVBCEnable C or A
StationSearch C or A
 6.1.36 Service X_AVM-DE_WANMobileConnection
Action User Rights
GetAccessTechnology C or A
GetBandCapabilities C or A
GetEnabledBandCapabilities C or A
GetInfo C or A
GetInfoEx C or A
GetPreferredAccessTechnology C or A
SetAccessTechnology C or A
SetEnabledBandCapabilities C or A
SetPIN C or A
SetPreferredAccessTechnology C or A
SetPUK C or A
Version: 59 28/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 6.1.37 Service X_AVM-DE_WANFiber
Action User Rights
GetInfo C or A
GetInfoGPON C or A
GetStatistics C or A
Version: 59 29/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 7 State Variables
The vendor specific state variables with the format X_<VENDOR>_StateVariable use the
following string for compatibility reasons.
AVM-DE
e.g. X_AVM-DE_Password
This string is not conform to the UPnP specification [UPNP11], that does not allow a – in
the name of a state variable.
 7.1 Compatibility with clients
At least one client has compatibility problems with these names of state variables.
Herqq UPnP http://www.herqq.org
 8 Support for Multiple Clients and Parallel Access
There are two features that support parallel access of multiple TR-064 clients to one TR064 service.
The first feature is the usage of transactions which is described in chapter 9 Transactions
 and the second feature are multiple internal session IDs.
These internal session IDs are used as parameters for URL access. The action GetCallList
() returns such an URL for example.
Version: 59 30/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 9 Transactions
To support TR-064 access for several clients in parallel the [TR064] specification
recommends the usage of transactions.
The description of the transactions is described in the chapter 6.3 Transactional Approach
and in the chapter 6.5.2 DeviceConfig in [TR064] .
Some message examples can be found on the page 87 and following pages.
 9.1 Usage of Transactions
Transactions cover the execution of SOAP actions without interruption by other TR-064
clients.
To begin a transaction the SOAP action DeviceConfig::ConfigurationStarted has to be
used.
To end a transaction the SOAP action DeviceConfig::ConfigurationFinished has to be
used.
The following sequence of SOAP actions is recommended for the communication with the
FRITZ! TR-064 service.
1. DeviceConfig::X_GenerateUUID
2. DeviceConfig::ConfigurationStarted
3. One or more SOAP actions to be protected against parallel access.
4. DeviceConfig::ConfigurationFinished
A transaction permits to successfully execute SOAP actions for its duration.
A transaction times out 45 seconds after DeviceConfig::ConfigurationStarted was called, if
the client does not call DeviceConfig::ConfigurationFinished to handle connection loss
between TR-064 client and server and similar failures.
 10 Examples
The following examples show unencrypted communication between a TR-064 enabled
CPE and a TR-064 application.
 10.1 SSDP request from application
Multicast HTTP message from application.
M-SEARCH * HTTP/1.1
HOST: 239.255.255.250:1900
MAN: "ssdp:discover"
MX: 5
ST: urn:dslforum-org:device:InternetGatewayDevice:1
 10.2 SSDP response from CPE
Unicast HTTP response message from CPE to application.
HTTP/1.1 200 OK
LOCATION: http://192.168.178.1:49000/tr64desc.xml
SERVER: FRITZBOX UPnP/1.0 AVM FRITZ!Box Fon WLAN 7270 v3 74.04.85
CACHE-CONTROL: max-age=1800
EXT:
ST: urn:dslforum-org:device:InternetGatewayDevice:1
USN: uuid:739f2409-bccb-40e7-8e6c-0024FE6E00C3::urn:dslforumorg:device:InternetGatewayDevice:1
Version: 59 31/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 10.3 Service description request from application
Unicast HTTP GET request message from application to CPE using values of SSDP
discovery response message LOCATION header.
GET /tr64desc.xml HTTP/1.1
HOST: 192.168.178.1:49000
CONNECTION: Close
USER-AGENT: AVM UPnP/1.0 Client 1.0
 10.4 Service description response from CPE
Unicast HTTP response message (trimmed) from CPE to application.
HTTP/1.0 200 OK
Content-Length: 10327
Content-Type: text/xml
Date: Thu, 17 Jun 2010 11:35:47 GMT
Last-Modified: Thu, 17 Jun 2010 10:33:08 GMT
Mime-Version: 1.0
<?xml version="1.0"?>
<root xmlns="urn:dslforum-org:device-1-0">
 <specVersion>
 <major>1</major>
 <minor>0</minor>
 </specVersion>
 <device>

<deviceType>urn:dslforum-org:device:InternetGatewayDevice:1</deviceType>
 <friendlyName>FRITZBOX UPnP/1.0 AVM FRITZ!Box Fon WLAN 7270 v3
74.04.85</friendlyName>
 <manufacturer>AVM</manufacturer>
 <manufacturerURL>https://fritz.com</manufacturerURL>
 <modelDescription>FRITZ!Box Fon WLAN 7270 v3</modelDescription>
 <modelName>FRITZ!Box Fon WLAN 7270 v3</modelName>
 <modelNumber> - avm</modelNumber>
 <modelURL>https://fritz.com</modelURL>
 <UDN>uuid:739f2409-bccb-40e7-8e6c-0024FE6E00C3</UDN>
 <iconList>
 <icon>
 <mimetype>image/gif</mimetype>
 <width>118</width>
 <height>119</height>
 <depth>8</depth>
 <url>/ligd.gif</url>
 </icon>
 </iconList>
 <serviceList>
 <service>
 <serviceType>urn:dslforum-org:service:DeviceInfo:1</serviceType>
 <serviceId>urn:DeviceInfo-com:serviceId:DeviceInfo1</serviceId>
 <controlURL>/upnp/control/deviceinfo</controlURL>
 <eventSubURL>/upnp/control/deviceinfo</eventSubURL>
 <SCPDURL>/deviceinfoSCPD.xml</SCPDURL>
 </service>
Version: 59 32/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
...
 </serviceList>
 <deviceList>
 <device>
 <deviceType>urn:dslforum-org:device:LANDevice:1</deviceType>
 <friendlyName>LANDevice - FRITZ!Box Fon WLAN 7270 v3</friendlyName>
 <manufacturer>AVM</manufacturer>
 <manufacturerURL>https://fritz.com</manufacturerURL>
 <modelDescription>LANDevice - FRITZ!Box Fon WLAN 7270
v3</modelDescription>
 <modelName>LANDevice - FRITZ!Box Fon WLAN 7270 v3</modelName>
 <modelNumber> - avm</modelNumber>
 <modelURL>https://fritz.com</modelURL>
 <UDN>uuid:75802409-bccb-40e7-8e6b-0024FE6E00C3</UDN>
 <UPC>AVM TR-064</UPC>
 <serviceList>
 <service>

<serviceType>urn:dslforum-org:service:WLANConfiguration:1</serviceType>
 <serviceId>urn:WLANConfigurationcom:serviceId:WLANConfiguration1</serviceId>
 <controlURL>/upnp/control/wlanconfig1</controlURL>
 <eventSubURL>/upnp/control/wlanconfig1</eventSubURL>
 <SCPDURL>/wlanconfigSCPD.xml</SCPDURL>
 </service>
...
 </serviceList>
 </device>
...
 </deviceList>
 <presentationURL>http://fritz.box</presentationURL>
 </device>
</root>
 10.5 SOAP Request with Action GetSecurityPort
Unicast SOAP message from application to CPE using values from SCPD description file
and SSDP discovery response message.
POST /upnp/control/deviceinfo HTTP/1.1
HOST: 192.168.178.1:49000
CONTENT-LENGTH: 267
CONTENT-TYPE: text/xml; charset="utf-8"
SOAPACTION: "urn:dslforum-org:service:DeviceInfo:1#GetSecurityPort"
USER-AGENT: AVM UPnP/1.0 Client 1.0
<?xml version="1.0"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/"
s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
<s:Body><u:GetSecurityPort xmlns:u="urn:dslforumorg:service:DeviceInfo:1"></u:GetSecurityPort>
</s:Body>
</s:Envelope>
Version: 59 33/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 10.6 SOAP Response for Action GetSecurityPort
Unicast SOAP message from CPE to application.
HTTP/1.1 200 OK
DATE: Thu, 17 Jun 2010 11:53:59 GMT
SERVER: FRITZBOX UPnP/1.0 AVM FRITZ!Box Fon WLAN 7270 v3 74.04.85
CONNECTION: keep-alive
CONTENT-LENGTH: 324
CONTENT-TYPE: text/xml; charset="utf-8"
EXT:
<?xml version="1.0"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/"
s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><s:Body>
<u:GetSecurityPortResponse xmlns:u="urn:dslforumorg:service:DeviceInfo:1">
<NewSecurityPort>49443</NewSecurityPort>
</u:GetSecurityPortResponse>
</s:Body> </s:Envelope>
Version: 59 34/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 10.7 Content Level Authentication
The used values are listed below.
Variable Value Remark
Username admin
Password gurkensalat
Realm F!Box SOAP-Auth Not writeable.
Nonce F758BE72FB999CEA
Auth b4f67585f22b0af7c4615db5a18faa14
Nonce (next value) 0B9813494DD27C93
 10.7.1 Initial Client Request
<?xml version="1.0" encoding="utf-8"?>
<s:Envelope s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"
xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" >
 <s:Header>
 <h:InitChallenge
 xmlns:h="http://soap-authentication.org/digest/2001/10/"
 s:mustUnderstand="1">
 <UserID>admin</UserID>
 </h:InitChallenge >
 </s:Header>
 <s:Body>
 <u:GetHostNumberOfEntries xmlns:u="urn:dslforum-org:service:Hosts:1">
 </u:GetHostNumberOfEntries>
 </s:Body>
</s:Envelope>
 10.7.2 Server Response to Initial Client Request
<?xml version="1.0" encoding="utf-8"?>
<s:Envelope s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"
xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" >
 <s:Header>
 <h:Challenge
 xmlns:h="http://soap-authentication.org/digest/2001/10/"
 s:mustUnderstand="1">
 <Status>Unauthenticated</Status>
 <Nonce>F758BE72FB999CEA</Nonce>
 <Realm>F!Box SOAP-Auth</Realm>
 </h:Challenge>
 </s:Header>
 <s:Body>
 <sFault>
 <errorCode>503</errorCode>
 <errorDescription>Auth. failed</errorDescription>
 </s:Fault>
 </s:Body>
</s:Envelope>
Version: 59 35/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 10.7.3 Client Request with Authentication
<?xml version="1.0" encoding="utf-8"?>
<s:Envelope s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"
xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" >
 <s:Header>
 <h:ClientAuth
 xmlns:h="http://soap-authentication.org/digest/2001/10/"
 s:mustUnderstand="1">
 <Nonce>F758BE72FB999CEA</Nonce>
 <Auth>b4f67585f22b0af7c4615db5a18faa14</Auth>
 <UserID>admin</UserID>
 <Realm>F!Box SOAP-Auth</Realm>
 </h:ClientAuth>
 </s:Header>
 <s:Body>
 <u:GetHostNumberOfEntries xmlns:u="urn:dslforum-org:service:Hosts:1">
 </u:GetHostNumberOfEntries>
 </s:Body>
</s:Envelope>
 10.7.4 Server Response for successful Authentication
<?xml version="1.0" encoding="utf-8"?>
<s:Envelope s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"
xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" >
 <s:Header>
 <h:NextChallenge
 xmlns:h="http://soap-authentication.org/digest/2001/10/"
 s:mustUnderstand="1">
 <Status>Authenticated</Status>
 <Nonce>0B9813494DD27C93</Nonce>
 <Realm>F!Box SOAP-Auth</Realm>
 </h:NextChallenge>
 </s:Header>
 <s:Body>
 <u:GetHostNumberOfEntriesResponse
 xmlns:u="urn:dslforum-org:service:Hosts:1">
 <HostNumberOfEntries>42</HostNumberOfEntries>
 </u:GetHostNumberOfEntriesResponse>
 </s:Body>
</s:Envelope>
Version: 59 36/37 2025-08-06
FRITZ! TR-064 - First Steps © FRITZ! GMBH
 11 Appendix
 11.1 References
TR064: DSL Forum, LAN-Side DSL CPE Configuration, 2004
SOAPAUTH01: Robert Cunnings, Rich Salz, SOAP Extensions: Basic and Digest
Authentication draft-cunnings-salz-soap-auth-00.txt, 2001,
https://datatracker.ietf.org/doc/html/draft-cunnings-salz-soap-auth-00
UPNP11: UPnP Forum, UPnP Device Architecture version 1.1, 2008
Version: 59 37/37 2025-08-06