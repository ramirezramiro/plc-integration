import socket
import struct
import pymongo

# Create a socket object
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

# Bind the socket to a specific IP address and port
ip_address = "192.168.1.100"  # Replace with the IP address of the Ethernet port
port = 5000  # Replace with the port number of the Ethernet port
sock.bind((ip_address, port))

# Connect to the MongoDB database
client = pymongo.MongoClient("mongodb://localhost:27017/")
db = client["mydatabase"]
collection = db["modbus_data"]

# Listen for incoming connections
sock.listen()

# Wait for a client to connect
print("Waiting for a client to connect...")
client_sock, client_addr = sock.accept()
print("Client connected:", client_addr)

# Receive data from the client, pack it into Modbus format, and send it to the MongoDB database
while True:
    data = client_sock.recv(1024)  # Receive up to 1024 bytes of data
    if not data:
        break
    print("Received data:", data.decode())
    
    # Pack the data into Modbus register format
    # Note: This assumes a 16-bit unsigned integer format (i.e. Modbus register)
    num_bytes = len(data)
    num_registers = num_bytes // 2
    register_format = ">" + "H" * num_registers  # ">" indicates big-endian byte order
    registers = struct.unpack(register_format, data)
    print("Packed data:", registers)
    
    # Send the packed data to the MongoDB database
    document = {"data": list(registers)}
    collection.insert_one(document)
    print("Data sent to MongoDB")

# Close the connection and disconnect from the MongoDB database
print("Closing connection...")
client_sock.close()
sock.close()
client.close()