package main

import (
	"flag"
	"fmt"
	"net"
	"regexp"
	"sync"
)

func worker(target string, ports <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for port := range ports {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}

		conn.Close()
		results <- port
	}
}

func main() {

	var (
		target             string
		startPort, endPort int
		wg                 sync.WaitGroup
		workersCount       int
	)

	// Define flags
	flag.StringVar(&target, "t", "", "target domain or IP address")
	flag.IntVar(&startPort, "s", 0, "start port (Default 0)")
	flag.IntVar(&endPort, "e", 1024, "end port (Default 1024)")
	flag.IntVar(&workersCount, "wc", 50, "Number of workers (Default 50)")
	flag.Parse()

	// Validate user input
	if !validateInput(target) {
		fmt.Println("Wrong input please input either domain or ip Example: port88r -t example.com OR port88r -t 45.33.32.156")
		flag.Usage()
		return
	}

	// Create channels for ports and results
	ports := make(chan int)
	results := make(chan int)

	// Launch workers
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go worker(target, ports, results, &wg)
	}

	// Send ports to workers
	go func() {
		for port := startPort; port <= endPort; port++ {
			ports <- port
		}
		close(ports)
	}()

	// Receive results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results
	for port := range results {
		fmt.Printf("port %d is open: %s:%d \n", port, target, port)
	}
}

// Function that validates user input as either a domain or IP address
func validateInput(input string) bool {
	// Define regex for domain and IP address
	domainRegex := regexp.MustCompile(`^([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`)
	ipRegex := regexp.MustCompile(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)

	return domainRegex.MatchString(input) || ipRegex.MatchString(input)
}
