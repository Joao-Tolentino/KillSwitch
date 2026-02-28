package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// Format the Payload
type Payload struct {
	Hash      string `json:"hash"`
	Timestamp int64  `json:"timestamp"`
}

// Computes the HMAC signature using the same env secret
func computeHMAC(secret string, timestamp int64) string {
	message := strconv.FormatInt(timestamp, 10)

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))

	return hex.EncodeToString(h.Sum(nil))
}

// Main function
func main() {
	// Gets the secret from the environment
	secret := os.Getenv("BOOT_SECRET")
	if secret == "" {
		panic("BOOT_SECRET not set")
	}

	// add a tag and the HMAC
	ts := time.Now().Unix()
	hash := computeHMAC(secret, ts)

	// Format the payload to be sent
	payload := Payload{
		Hash: hash,
		Timestamp: ts,
	}

	// Define and send the POST request to the server
	body, _ := json.Marshal(payload)
	resp, err := http.Post(
		"http://127.0.0.1:8888/boot",
		"application/json",
		bytes.NewBuffer(body),
	)

	// Display any error
	if err != nil {
		panic(err)
	}

	// Show response status
	fmt.Println("Status:", resp.Status)
}