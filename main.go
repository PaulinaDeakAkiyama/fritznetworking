package main

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"fritznetworking/captivePortal"
	"fritznetworking/internal"
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

// -------- HTTP CLIENT (skip TLS verify like PowerShell bc local cert is causing issues) --------
func createHTTPClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // same behavior as your script
	}
	return &http.Client{Transport: tr}
}

// -------- MAIN --------
func main() {

	client := createHTTPClient()

	// Get number of hosts
	data, err := captivePortal.InvokeFritzRequest(client, "Hosts", "GetHostNumberOfEntries", nil)
	if err != nil {
		panic(err)
	}

	var env Envelope

	if err := xml.Unmarshal(data, &env); err != nil {
		panic(err)
	}

	count := env.Body.GetHostNumberOfEntriesResponse.NewHostNumberOfEntries
	fmt.Printf("Hosts: %d\n", count)

	// Enumerate hosts
	for i := 1; i <= 10; i++ {
		args := map[string]string{
			"NewIndex": fmt.Sprintf("%d", i),
		}

		data, err := captivePortal.InvokeFritzRequest(client, "Hosts", "GetGenericHostEntry", args)
		if err != nil {
			fmt.Printf("Failed index %d: %v\n", i, err)
			continue
		}

		var env Envelope

		if err := xml.Unmarshal(data, &env); err != nil {
			fmt.Printf("XML parse error index %d: %v\n", i, err)
			continue
		}

		host := env.Body.GetGenericHostEntryResponse
		if host == nil {
			continue
		}

		if host.NewActive == 0 {
			continue
		}

		fmt.Printf("Index: %d\n", i)
		fmt.Printf("  HostName: %s\n", host.NewHostName)
		fmt.Printf("  IP: %s\n", host.NewIPAddress)
		fmt.Printf("  MAC: %s\n", host.NewMACAddress)
		fmt.Printf("  Active: %d\n\n", host.NewActive)
	}

	fmt.Println("blocking paulinas phone...")

	err := captivePortal.InvokeFritzRequest(client, "Hosts", "GetSpecificHostEntry", map[string]string{"NewMACAddress": internal.PaulinasMobileIP}, 0)

	err = captivePortal.BlockHost(client, internal.PaulinasMobileIP)

}
