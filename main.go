package main

import (
	"log"
)

func main() {
	handler, err := connectToEthernetPort()
	if err != nil {
		log.Fatal("Error connecting to Ethernet port: ", err)
	}
	defer handler.Close()

	data, err := readModbusData(handler)
	if err != nil {
		log.Fatal("Error reading Modbus data: ", err)
	}

	err = sendDataToMongoDB(data)
	if err != nil {
		log.Fatal("Error sending data to MongoDB: ", err)
	}

	log.Println("Modbus data saved to MongoDB successfully!")
}
