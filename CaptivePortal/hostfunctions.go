package captivePortal

import (
	"fmt"
	"net/http"
	"time"
)

func BlockHost(client *http.Client, ip string) error {
	data, err := InvokeFritzRequest(
		client,
		"X_AVM-DE_HostFilter",
		"DisallowWANAccess",
		map[string]string{
			"NewIPv4Adress": ip,
		},
	)
	if err != nil {
		return err
	}
	fmt.Printf("Block response: %s\n", string(data))
	return nil
}

func AllowHost(client *http.Client, ip string) error {
	_, err := InvokeFritzRequest(
		client,
		"X_AVM-DE_HostFilter",
		"AllowWANAccess",
		map[string]string{
			"NewIPv4Address": ip,
		},
	)
	return err
}

func UnblockAfter(client *http.Client, ip string, durationMinutes int) {
	go func() {
		fmt.Printf("Unblocking %s in %d minutes...\n", ip, durationMinutes)
		time.Sleep(time.Duration(durationMinutes) * time.Minute)
		err := AllowHost(client, ip)
		if err != nil {
			fmt.Printf("Failed to unblock %s: %v\n", ip, err)
			return
		}
		fmt.Printf("Host %s unblocked\n", ip)
	}()
}

func ReblockAfter(client *http.Client, ip string, durationMinutes int) {
	go func() {
		time.Sleep(time.Duration(durationMinutes) * time.Minute)
		err := BlockHost(client, ip)
		if err != nil {
			fmt.Printf("Failed to re-block %s: %v\n", ip, err)
			return
		}
		fmt.Printf("Host %s blocked again\n", ip)
	}()
}
