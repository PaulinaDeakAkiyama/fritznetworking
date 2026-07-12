package internal

// -------- SOAP RESPONSE STRUCTS --------
type Envelope struct {
	Body Body `xml:"Body"`
}

type Body struct {
	GetSecurityPortResponse        *GetSecurityPortResponse        `xml:"GetSecurityPortResponse"`
	GetHostNumberOfEntriesResponse *GetHostNumberOfEntriesResponse `xml:"GetHostNumberOfEntriesResponse"`
	GetGenericHostEntryResponse    *GetGenericHostEntryResponse    `xml:"GetGenericHostEntryResponse"`
}

type GetSecurityPortResponse struct {
	NewSecurityPort int `xml:"NewSecurityPort"`
}

type GetHostNumberOfEntriesResponse struct {
	NewHostNumberOfEntries int `xml:"NewHostNumberOfEntries"`
}

type GetGenericHostEntryResponse struct {
	NewHostName   string `xml:"NewHostName"`
	NewIPAddress  string `xml:"NewIPAddress"`
	NewMACAddress string `xml:"NewMACAddress"`
	NewActive     int    `xml:"NewActive"`
}
