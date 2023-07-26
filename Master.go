package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	slaves := []string{"192.168.137.227:9090", "192.168.137.227:9091", "192.168.137.227:9092"}
	distributeFile(slaves)
	// Define the IP address and port number to listen on
	ip := "192.168.137.171"
	port := "8080"
	listener, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Master listening on " + ip + ":" + port)

	// Accept incoming connections and handle them in separate goroutines
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn, slaves)
	}
}

type Request struct {
	ChunkIndex int `json:"chunk_index"`
}

type Response struct {
	SlaveIP   string `json:"slave_ip"`
	SlavePort int    `json:"slave_port"`
}

func handleConnection(conn net.Conn, slaves []string) {	
	numSlaves := 3
	// Return IP and port number of slave node
	for i := 0; i < numSlaves; i++ {
		if i < numSlaves-1 {
			slaves[i] = slaves[i] + ", "
		}
		_, err := conn.Write([]byte(slaves[i]))
		if err != nil {
			fmt.Println("Error sending data to client:", err)
			conn.Close()
			return
		}
	}
	fmt.Println("sending successfully")
	conn.Close()
}

func distributeFile(slaves []string) {
	// Open the file for reading
	file, err := os.Open("largefile.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Calculate the chunk size based on the file size and number of slaves
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	numSlaves := 3
	chunkSize := fileSize / int64(numSlaves)

	// Connect to each slave node and send a chunk of the file
	for i := 0; i < numSlaves; i++ {
		// Calculate the offset and size of the chunk
		offset := int64(i) * chunkSize
		size := chunkSize
		if i == numSlaves-1 {
			size = fileSize - offset
		}
		// Read the chunk from the file
		buffer := make([]byte, size)
		_, err = file.ReadAt(buffer, offset)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		// Connect to the slave node
		conn, err := net.Dial("tcp", slaves[i])
		if err != nil {
			fmt.Println("Error connecting to slave:", err)
			return
		}
		defer conn.Close()

		// Send the chunk to the slave node
		_, err = conn.Write(buffer)
		if err != nil {
			fmt.Println("Error sending data to slave:", err)
			return
		}
		fmt.Println("Send successfully")		
	}

	fmt.Println("File distributed to slave nodes.")
}
