# FTP Brute Force Attack Tool

This tool is a simple FTP brute force password cracker written in Go. It utilizes a wordlist to attempt login to an FTP server, with the option to use FTPS (FTP over TLS) and to delay login attempts for better stealth.

## Features

- Brute force attack for FTP login.
- Option to use FTPS (FTP over TLS).
- Configurable number of concurrent connections.
- Option to delay login attempts (`-time` flag).
- Simple color output to indicate success or failure.

## Prerequisites

- Go (version 1.18 or higher)
- FTP server with exposed login (the tool works on servers accepting FTP or FTPS connections).

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/Mogra12/FTP-BruteForce
    cd FTP-BruteForce
    ```

2. Install dependencies (if needed):
    ```bash
    go mod tidy
    ```

## Usage

You can build and run the application with the following command:

```bash
go run main.go -wl <path_to_wordlist> -h <hostname:port> [options]
