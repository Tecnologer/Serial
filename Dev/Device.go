package dev

import (
	"fmt"
	"log"

	"go.bug.st/serial.v1"
)

func PrintPortsList() {
	ports := getPorts()

	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
}
func getPorts() []string {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}

	return ports
}
func ConectTo(index int) {
	mode := &serial.Mode{
		BaudRate: 1000000,
		Parity:   serial.EvenParity,
	}
	var portName = getPorts()[index]
	fmt.Printf("Connecting to %s\n", portName)
	port, err := serial.Open(portName, mode)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully connected to %s\n", portName)
	buffer := make([]byte, 100)
	for {
		fmt.Printf("Reading from %s\n", portName)
		n, err := port.Read(buffer)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}
		fmt.Printf("%v", string(buffer[:n]))
	}
}

func Test() {
	// Retrieve the port list
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	// Print the list of detected ports
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}

	// Open the first serial port detected at 9600bps N81
	mode := &serial.Mode{
		BaudRate: 1000000,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open("COM3", mode)
	if err != nil {
		log.Fatal(err)
	}

	// Send the string "10,20,30\n\r" to the serial port
	n, err := port.Write([]byte{0x00, 0x01, 0x04, 0x01, 0x01, 0xFA, 0x04})
	n, err = port.Write([]byte{0x00, 0x01, 0x04, 0x00, 0x13, 0xE9, 0x04})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)

	// Read and print the response
	buff := make([]byte, 100)

	for {
		// Reads up to 100 bytes
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}
		fmt.Println(buff[:n])
	}
}
