package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// Listen for incoming connections from the Master server
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		// handle error
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Waiting for master server to connect...")

	conn, err := listener.Accept()

	fmt.Println("Master server connected successfully")

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Receive the file from the master server and write it to a new file on the local system
	file, err := os.Create("received-file.txt")
	if err != nil {
		// handle error
		panic(err)
	}
	defer file.Close()

	// Read the data from the connection and write it to the file
	_, err = io.Copy(file, conn)
	if err != nil {
		// handle error
		panic(err)
	}

	// --Send a string to the server
	// message := "ok"
	// _, err = conn.Write([]byte(message))
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("file has been received successfully")

	conn.Close()
	listener.Close()
	//Listen for incoming connections from any client
	listener, err = net.Listen("tcp", ":9090")
	if err != nil {
		// handle error
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Waiting for clients to connect...")

	conn, err = listener.Accept()

	fmt.Println("A Client connected successfully")

	if err != nil {
		// handle error
		panic(err)
	}
	defer conn.Close()

	//Open the file to be sent--
	file1, err := os.Open("received-file.txt")
	if err != nil {
		// handle error
		panic(err)
	}
	defer file1.Close()

	// Send the file to the slave
	_, err = io.Copy(conn, file1)
	if err != nil {
		// handle error
		panic(err)
	}

	fmt.Println("file sent successfully to Client")

	/////////////////////////////
	// Once the file is received, you can start serving the file to a client by creating a TCP listener and accepting incoming connections
	// 	listener, err = net.Listen("tcp", ":8090")
	// 	if err != nil {
	// 		// handle error
	// 		panic(err)
	// 	}
	// 	defer listener.Close()

	// 	//fmt.Println("Waiting for clients to connect...")

	// 	for {
	// 		fmt.Println("Waiting for clients to connect...")
	// 		conn, err := listener.Accept()
	// 		if err != nil {
	// 			// handle error
	// 			panic(err)
	// 		}
	// 		fmt.Println("Client connected")

	// 		// Read the message from the slave
	// 		buf := make([]byte, 1024)
	// 		n, err := conn.Read(buf)
	// 		if err != nil {
	// 			// handle error
	// 			panic(err)
	// 		}
	// 		msg := string(buf[:n])

	// 		if msg == "send_file" {
	// 			// Open the file to be sent
	// 			file1, err := os.Open("received-file.txt")
	// 			if err != nil {
	// 				// handle error
	// 				panic(err)
	// 			}
	// 			defer file1.Close()

	// 			// Send the file to the slave
	// 			_, err = io.Copy(conn, file1)
	// 			if err != nil {
	// 				// handle error
	// 				panic(err)
	// 			}

	// 			fmt.Println("File sent to Client")
	// 		}

	// 		conn.Close()
	// 	}
}
