package tcp

import (
	"fmt"
	"io/ioutil"
	"net"
)

func SendTextToServer(text string) error {
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		return fmt.Errorf("error connecting: %v", err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, text+"\n")

	audioData, err := ioutil.ReadAll(conn)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}

	// Save the received audio data to a file
	err = ioutil.WriteFile("response.wav", audioData, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	fmt.Println("Audio saved to response.wav")
	return nil
}
