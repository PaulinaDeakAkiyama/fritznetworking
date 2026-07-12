package captivePortal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AccessRequest struct {
	Duration int `json:"duration"`
}

func startPortal(client *http.Client, targetIP string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/request-access", func(w http.ResponseWriter, r *http.Request) {
		var req AccessRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request", 400)
			return
		}

		if req.Duration <= 0 || req.Duration > 180 {
			http.Error(w, "Invalid duration (1–180 min)", 400)
			return
		}

		// Allow access now
		err = AllowHost(client, targetIP)
		if err != nil {
			http.Error(w, "Failed to allow access", 500)
			return
		}

		// Schedule re-block (IMPORTANT)
		UnblockAfter(client, targetIP, req.Duration)

		msg := fmt.Sprintf("Access granted for %d minutes", req.Duration)
		w.Write([]byte(msg))
	})

	// http.HandleFunc("/login", login)

	fmt.Println("Captive portal running on 192.168.18.20:8080")
	http.ListenAndServe("192.168.18.20:8080", nil)

}
