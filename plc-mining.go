package main

import (
	"context"
	"log"
	"time"

	"fmt"
	"net"

	"github.com/goburrow/modbus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ModbusData struct {
	Coil     bool
	Register uint16
}

func eth_read() {
	// Set up a TCP connection to the Ethernet port
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error connecting to Ethernet port: ", err)
	}
	defer conn.Close()

	// Enter a loop to read data from the Ethernet port and print the status
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal("Error reading from Ethernet port: ", err)
		}

		// Print the status of the data received from the Ethernet port
		fmt.Printf("Received %d bytes from Ethernet port: %s\n", n, string(buf[:n]))

		// Wait for 5 seconds before reading again
		time.Sleep(5 * time.Second)
	}
}

func modbus_conn() {
	// Define the Modbus connection parameters
	handler := modbus.NewTCPClientHandler("localhost:502")
	handler.Timeout = 5 * time.Second
	err := handler.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer handler.Close()

	// Create a Modbus client and read data from the device
	client := modbus.NewClient(handler)
	data, err := client.ReadCoils(0, 10)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the MongoDB database
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(ctx)

	// Insert the Modbus data into the MongoDB database
	collection := mongoClient.Database("test").Collection("modbus_data")
	for i, d := range data {
		modbusData := ModbusData{Coil: d != 0, Register: uint16(i)}
		_, err = collection.InsertOne(ctx, modbusData)
		if err != nil {
			log.Fatal(err)
		}
	}
}
