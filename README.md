# KillSwitch

[![License: AGPL v3](https://img.shields.io/badge/License-AGPL_v3-blue.svg)](LICENSE)
[![Platform: Windows](https://img.shields.io/badge/Platform-Windows-0078D6.svg?logo=windows&logoColor=white)](#)

A secure, local Go web server designed to act as an emergency shutdown handler. It receives authenticated HTTP POST requests and executes an immediate machine shutdown (when fully uncommented) upon verifying a cryptographic signature.

---

## Features

- **HMAC-SHA256 Authentication**: Validates incoming requests by hashing a UNIX timestamp with your private `KS_SECRET` environment variable.
- **Replay Attack Prevention**: Automatically rejects any valid signal if its timestamp is older than 60 seconds.
- **Lightweight Go Server**: Runs natively via Go's `net/http` package on `localhost:8888`.

---

## Quick Start

1. Clone or download the repository.
2. Ensure you have the Go runtime installed.
3. Set your secret key environment variable:
   - Windows: `set KS_SECRET=my_super_secret`
   - Linux/Mac: `export KS_SECRET=my_super_secret`
4. Run `go run .` or build the executable.

---

## Configuration Details

You **must** set the `KS_SECRET` environment variable before running the application, or the server will return a 500 error on any incoming payload. The incoming payload requires a JSON structure containing a `hash` and a `timestamp`.

---

## Usage Guidelines

- Start the server on `localhost:8888`.
- Send a POST request with `Content-Type: application/json` containing:
  ```json
  {
    "hash": "your_generated_hmac_hex",
    "timestamp": 1696000000
  }
  ```
- If the signature is valid, the `ks()` function is triggered in `killSwitch.go` (uncomment the `os/exec` import and execution code to enable actual system shutdown).

---

## Technical Documentation

For developers interested in directory structures, code architecture, or compilation guidelines, please refer to the **[Documentation.md](Documentation.md)** file.

---

## License

This project is licensed under the **GNU Affero General Public License Version 3 (AGPLv3)**. See the LICENSE file for details.
