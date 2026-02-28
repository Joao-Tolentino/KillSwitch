package main

// Imports
import (
	"encoding/json"
	"encoding/hex"
	"log"
	"os"
	"time"
	"strconv"
	"net/http"
	"crypto/hmac"
	"crypto/sha256"	
)

// Define the Payload to receive
type Payload struct {
	Hash string `json:"hash"`
	Timestamp int64 `json:"timestamp"`
}

// Recomputes the HMAC from the signal
func computeHMAC(secret string, timestamp int64) string {
	message := strconv.FormatInt(timestamp, 10)

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))

	return hex.EncodeToString(h.Sum(nil))
}

// Validate the Hmac using environment variable secret
func verifyHMAC(secret string, payload Payload) bool {
	expected := computeHMAC(secret, payload.Timestamp)
	return hmac.Equal([]byte(expected), []byte(payload.Hash))
}

// Kill Switch signal handler
func handler(w http.ResponseWriter, r *http.Request) {
	// Gets the env secret
	secret := os.Getenv("KS_SECRET")
	if secret == "" {
		http.Error(w, "Missing secret", 500)
		return
	}

	// Check the method used, only accept Post
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Check the MIME for application/json
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Unsupported content type", http.StatusUnsupportedMediaType)
		return
	}

	// Try to decode the JSON using the Payload struct
	var payload Payload
	// Not accept if it isnt formatted correctly
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	// Check HMAC
	if !verifyHMAC(secret, payload) {
		http.Error(w, "Invalid signature", 401)
		return
	}

	// Check timestamp freshness (e.g., 60s)
	if time.Now().Unix()-payload.Timestamp > 60 {
		http.Error(w, "Expired timestamp", 401)
		return
	}

	// Activate Kill Switch if signal is correct
	log.Println("Kill Switch Activated! \nGoodBye World! o7")
	ks()
}

// Main function
func main(){
	// Setup the handler and Server in specified PORT
	http.HandleFunc("/", handler)

	// Print the server is running
	log.Println("Server is Running!")

	// Listen in localhost:PORT
	err := http.ListenAndServe("localhost:8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}