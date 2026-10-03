# Developer & Technical Documentation

This document provides a technical guide to the **KillSwitch** application's architecture and execution.

---

## System Architecture

The application is built using Go's standard library for HTTP handling (`net/http`) and cryptography (`crypto/hmac`, `crypto/sha256`).

```mermaid
graph TD
    Client[External Client] -->|POST JSON| Server[HTTP Handler on 8888]
    Server -->|Parse| Payload[{hash, timestamp}]
    Server -->|Read| Env[os.Getenv KS_SECRET]
    Env -->|Compute HMAC| Verification
    Payload -->|Check Freshness| Verification
    Verification -->|Success| Exec[ks function]
    Exec -->|Optional| OSShutdown[exec.Command shutdown]
```

---

## Directory Structure & File Roles

```
.
├── main.go             # Server initialization, HMAC logic, and HTTP handler
├── killSwitch.go       # The ks() function containing the (commented) exec.Command
├── go.mod              # Module definition
├── README.md           # General overview
└── Documentation.md    # Technical documentation
```

---

## Workflow

The execution flow of KillSwitch:
1. **Initialization**: Standard `http.ListenAndServe` starts listening on port 8888.
2. **Action Trigger**: An incoming POST request hits the `/` route.
3. **Processing**: The handler checks that the `KS_SECRET` is set, the method is POST, and the Content-Type is `application/json`. It parses the timestamp, computes the expected SHA256 HMAC using `KS_SECRET`, and compares it to the provided hash string. It then verifies the timestamp is within 60 seconds of `time.Now().Unix()`.
4. **Execution**: If valid, it triggers the `ks()` package method to execute the payload.

---

## Launcher Compilation Guide

If you need to compile or recompile the standalone executable for KillSwitch, follow the standard Go commands.

### Compilation or Execution Commands

Execute the following commands in order within your terminal:

```powershell
# Resolve any dependencies (though it relies solely on stdlib)
go mod tidy

# Build the executable
go build -o killswitch.exe .

# Or run directly for development
go run .
```
