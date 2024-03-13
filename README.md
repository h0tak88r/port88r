# port88r
Port88r is a lightweight and concurrent TCP port scanner tool written in Go. It is designed to quickly scan a range of TCP ports on a target domain or IP address to identify open ports.

# Installation
```bash
go install github.com/h0tak88r/port88r@latest
```

# Usage
```bash
port88r -t example.com OR port88r -t 45.33.32.156
Usage of port88r:
  -e int
        end port (Default 1024) (default 1024)
  -s int
        start port (Default 0)
  -t string
        target domain or IP address
  -wc int
        Number of workers (Default 50) (default 50)
```
