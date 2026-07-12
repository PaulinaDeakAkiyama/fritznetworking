// package main

// import (
// 	"crypto/tls"
// 	"encoding/xml"
// 	"fmt"
// 	"fritznetworking/captivePortal"
// 	"fritznetworking/internal"
// 	"net/http"
// )

// // -------- HTTP CLIENT (skip TLS verify like PowerShell bc local cert is causing issues) --------
// func createHTTPClient() *http.Client {
// 	tr := &http.Transport{
// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // same behavior as your script
// 	}
// 	return &http.Client{Transport: tr}
// }

// // -------- MAIN --------
// func main() {

// 	client := createHTTPClient()

// 	// Get number of hosts
// 	data, err := captivePortal.InvokeFritzRequest(client, "Hosts", "GetHostNumberOfEntries", nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var env internal.Envelope

// 	if err := xml.Unmarshal(data, &env); err != nil {
// 		panic(err)
// 	}

// 	count := env.Body.GetHostNumberOfEntriesResponse.NewHostNumberOfEntries
// 	fmt.Printf("Hosts: %d\n", count)

// 	// Enumerate hosts
// 	for i := 1; i <= 10; i++ {
// 		args := map[string]string{
// 			"NewIndex": fmt.Sprintf("%d", i),
// 		}

// 		data, err := captivePortal.InvokeFritzRequest(client, "Hosts", "GetGenericHostEntry", args)
// 		if err != nil {
// 			fmt.Printf("Failed index %d: %v\n", i, err)
// 			continue
// 		}

// 		var env internal.Envelope

// 		if err := xml.Unmarshal(data, &env); err != nil {
// 			fmt.Printf("XML parse error index %d: %v\n", i, err)
// 			continue
// 		}

// 		host := env.Body.GetGenericHostEntryResponse
// 		if host == nil {
// 			continue
// 		}

// 		if host.NewActive == 0 {
// 			continue
// 		}

// 		fmt.Printf("Index: %d\n", i)
// 		fmt.Printf("  HostName: %s\n", host.NewHostName)
// 		fmt.Printf("  IP: %s\n", host.NewIPAddress)
// 		fmt.Printf("  MAC: %s\n", host.NewMACAddress)
// 		fmt.Printf("  Active: %d\n\n", host.NewActive)
// 	}

// 	fmt.Println("blocking paulinas phone...")

// 	err = captivePortal.BlockHost(client, internal.PaulinasMobileIP)

// }

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var srv *http.Server

func main() {

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    "127.0.0.1:5000",
		Handler: mux,
	}

	// API routes
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hi", hiHandler)
	mux.HandleFunc("/request-access", requestAccessHandler)
	mux.HandleFunc("/admin", auth(adminHandler))
	mux.HandleFunc("/admin/shutdown", auth(shutdownHandler))

	fmt.Println("Server is running on port" + "127.0.0.1:5000")
	// Start server on port specified above

	go func() {
		if err := srv.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	fmt.Println("Waiting for interrupt signal to gracefully shutdown the server...")
	<-sig
	shutdownServer(srv)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "www/testpage.html")
}
func hiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}
func requestAccessHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "www/blockedwifi.html")
}
func adminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, admin!")
	// redirect to fritz box webpage
}
func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Shutting down server...")
	shutdownServer(srv)
}

func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok || user != "admin" || pass != "secret" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func shutdownServer(srv *http.Server) {
	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
	fmt.Println("Server gracefully stopped")
}
