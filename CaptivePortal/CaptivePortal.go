package main

import (
	"fmt"
	"net/http"
	"time"
)

func blockHost(client *http.Client, port int, ip string) error {
	data, err := invokeFritzRequest(
		client,
		"X_AVM-DE_HostFilter",
		"DisallowWANAccess",
		map[string]string{
			"NewIPv4Adress": ip,
		},
		port,
	)
	if err != nil {
		return err
	}
	fmt.Printf("Block response: %s\n", string(data))
	return nil
}

func allowHost(client *http.Client, port int, ip string) error {
	_, err := invokeFritzRequest(
		client,
		"X_AVM-DE_HostFilter",
		"AllowWANAccess",
		map[string]string{
			"NewIPv4Address": ip,
		},
		port,
	)
	return err
}

func unblockAfter(client *http.Client, port int, ip string, durationMinutes int) {
	go func() {
		fmt.Printf("Unblocking %s in %d minutes...\n", ip, durationMinutes)
		time.Sleep(time.Duration(durationMinutes) * time.Minute)
		err := allowHost(client, port, ip)
		if err != nil {
			fmt.Printf("Failed to unblock %s: %v\n", ip, err)
			return
		}
		fmt.Printf("Host %s unblocked\n", ip)
	}()
}

func reblockAfter(client *http.Client, port int, ip string, durationMinutes int) {
	go func() {
		time.Sleep(time.Duration(durationMinutes) * time.Minute)
		err := blockHost(client, port, ip)
		if err != nil {
			fmt.Printf("Failed to re-block %s: %v\n", ip, err)
			return
		}
		fmt.Printf("Host %s blocked again\n", ip)
	}()
}
