package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	// Initialize the list of IP addresses
	var addresses []string

	// Connect to the device that provides the list of IP addresses
	conn, err := net.Dial("tcp", "192.168.137.171:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Connected to master")
	// Read the list of IP addresses from the connection
	reader := bufio.NewReader(conn)
	for {
		address, err := reader.ReadString('\n')
		fmt.Println(address)
		if err != nil && err != io.EOF {
			panic(err)
		}

		addresses = strings.Split(address, ", ")
		if err == io.EOF {
			break
		}
	}
	fmt.Println("Recieved Succesfully")
	// Create a new slice to hold the combined data from all files
	var combinedData []string
	conn.Close()
	// Iterate over the list of addresses and listen for incoming connections
	for _, address := range addresses {
		fmt.Printf("Listening on %s...\n", address)

		conn, err := net.Dial("tcp", address)

		// Read the data from the connection
		reader := bufio.NewReader(conn)
		data, err := readTextFile(reader)
		if err != nil {
			fmt.Println("Error reading data:", err.Error())
			conn.Close()
			continue
		}

		// Append the data to the combined data slice
		combinedData = append(combinedData, data)

		// Close the connection
		conn.Close()

		fmt.Printf("Received data from %s\n", conn.RemoteAddr())

	}

	// Open a new file to write the combined data to
	outputFile, err := os.Create("combined.txt")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Write the combined data to the new file
	writer := bufio.NewWriter(outputFile)
	for _, data := range combinedData {
		_, err := writer.WriteString(data)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	fmt.Println("Combined data written to combined.txt")
}

// Read a text file from a reader
func readTextFile(reader *bufio.Reader) (string, error) {
	var data string
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return "", err
		}
		data += line
		if err == io.EOF {
			break
		}
	}
	return data, nil
}
