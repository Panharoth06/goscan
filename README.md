# GoScanner

**GoScanner** is a lightweight, high-performance network service scanner written in Go.  
It performs concurrent port scanning, banner grabbing, and TLS intelligence collection. Its output is colorized for readability.



## Features (Current State)

- **Concurrent scanning** with configurable number of workers  
- **Port scanning** for TCP services  
- **Banner grabbing** for common protocols (HTTP, FTP, SSH, SMTP, IMAP, POP3, etc.)  
- **TLS intelligence**:
  - TLS version detection
  - Cipher suite extraction
  - Certificate common name (CN) and issuer
  - Certificate expiration date  
- **Colorized CLI output** for easier reading  
- **Structured service information** for further automation or reporting  



## Installation

1. Clone the repository:

```bash
git clone https://github.com/Panharoth06/goscanner.git
cd goscanner
```

2. Build the binary 
```bash
go build -o goscanner main.go
```

3. Run the scanner 
```bash
./goscanner -d scanme.nmap.org
```

## Usage
```bash
Usage: gorecon -d <target> [options]

Options:
  -d, -domain string        Target host or IP to scan
  -w, -workers int          Number of concurrent workers (default: 500)
  -h, -help                 Show help

```

### Example
```bash
./gorecon -d scanme.nmap.org -w 300
```

## Architecture 
- Worker Pool: Goroutines handle concurrent port scanning
- GrabService Engine: Handles banner grabbing, TLS handshake, and protocol detection
- TLS Fingerprint Module: Extracts certificate information and cipher details
- Colorized CLI: Uses ANSI codes for better readability

## Requirement
- Go 1.20+
- Linux, macOS, or Windows terminal with ANSI support
- Network access to target hosts

## Current Limitations
- Service detection is port-based and uses heuristics; some non-standard ports may be missed
- Binary protocols (e.g., PostgreSQL, MySQL) are only partially supported
- STARTTLS detection is not implemented
- JSON output or structured export not implemented yet
- Vulnerability detection is limited (TLS weak versions, expired certificates are not yet flagged)

## Next Steps / Roadmap
1. Add modular probe system for service-specific detection
2. STARTTLS and opportunistic TLS detection
3. JSON / CSV export for automation and reporting
4. Weak TLS and certificate risk detection
5. Subdomain and HTTP fingerprinting
6. JA3 TLS fingerprinting for advanced reconnaissance