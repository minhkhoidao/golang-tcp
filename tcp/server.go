package tcp

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	fmt.Println("Received:", text)

	// Create a temporary file for the audio output
	audioFile := "output.wav"
	cmd := exec.Command("flite", "-t", text, "-o", audioFile)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error generating speech:", err.Error())
		return
	}

	// Read the audio file
	audioData, err := ioutil.ReadFile(audioFile)
	if err != nil {
		fmt.Println("Error reading audio file:", err.Error())
		return
	}

	// Send the audio data back to the client
	conn.Write(audioData)

	// Clean up
	os.Remove(audioFile)
}

func StartTCPServer() {
	server, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("Error starting TCP server:", err.Error())
		os.Exit(1)
	}
	defer server.Close()

	fmt.Println("TCP Server started, listening on :9090")

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}
		go handleConnection(conn)
	}
}
