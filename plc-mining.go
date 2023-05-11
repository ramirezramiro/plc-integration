package main

import (
	"context"
	"time"

	"github.com/goburrow/modbus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ModbusData []byte

func connectToEthernetPort() (*modbus.TCPClientHandler, error) {
	handler := modbus.NewTCPClientHandler("localhost:502")
	handler.Timeout = 5 * time.Second
	err := handler.Connect()
	if err != nil {
		return nil, err
	}
	return handler, nil
}

func readModbusData(handler *modbus.TCPClientHandler) (ModbusData, error) {
	client := modbus.NewClient(handler)
	data, err := client.ReadCoils(0, 10)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func sendDataToMongoDB(data ModbusData) error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	defer mongoClient.Disconnect(ctx)

	collection := mongoClient.Database("test").Collection("modbus_data")
	_, err = collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
