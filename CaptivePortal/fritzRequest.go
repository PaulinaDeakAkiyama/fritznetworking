package captivePortal

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

// -------- CONFIG --------
// Replace with your credentials
const username = "fritzadmin"
const password = "masteroffritz2002"

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

// -------- HTTP CLIENT (skip TLS verify like PowerShell) --------
func CreateHTTPClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // same behavior as your script
	}
	return &http.Client{Transport: tr}
}

// -------- SOAP HELPERS --------
func BuildSoapEnvelope(service, action string, args map[string]string) string {
	argsXML := ""
	for k, v := range args {
		argsXML += fmt.Sprintf("<%s>%s</%s>", k, v, k)
	}
	return fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
  <s:Body>
    <u:%s xmlns:u="urn:dslforum-org:service:%s:1">
      %s
    </u:%s>
  </s:Body>
</s:Envelope>`, action, service, argsXML, action)
}

func InvokeFritzRequest(client *http.Client, service, action string, args map[string]string) ([]byte, error) {

	port, err := GetSecurityPort(client)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Security Port: %d\n", port)

	url := fmt.Sprintf("https://fritz.box:%d/upnp/control/%s", port, lower(service))
	body := BuildSoapEnvelope(service, action, args)
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", fmt.Sprintf(`"urn:dslforum-org:service:%s:1#%s"`, service, action))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// simple lowercase helper (to match PowerShell behavior)
func lower(s string) string {
	return string(bytes.ToLower([]byte(s)))
}

// -------- GET SECURITY PORT --------
func GetSecurityPort(client *http.Client) (int, error) {
	body := `<?xml version="1.0" encoding="utf-8"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
<s:Body>
<u:GetSecurityPort xmlns:u="urn:dslforum-org:service:DeviceInfo:1"/>
</s:Body>
</s:Envelope>`
	req, err := http.NewRequest("POST", "http://fritz.box:49000/upnp/control/deviceinfo", bytes.NewBufferString(body))
	if err != nil {
		return 0, err
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", `"urn:dslforum-org:service:DeviceInfo:1#GetSecurityPort"`)
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	var env Envelope
	if err := xml.Unmarshal(data, &env); err != nil {
		return 0, err
	}
	return env.Body.GetSecurityPortResponse.NewSecurityPort, nil
}
