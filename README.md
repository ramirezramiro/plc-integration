# PLC Integration

---

This repository contains scripts to capture data packets from an Ethernet port, pack its data into Modbus dataframe and send the content to a Mongo database.

First, create a socket object and sniff the IP address and port of the Ethernet port for data collection.

Secondly, once there is a stable socket connection, initiate a loop to receive data from the client and pack it into Modbus register format.

Finally, send the data to the MongoDB database.

Note that you may need to modify the MongoDB connection string to match your specific MongoDB configuration.

<br>

---

### Wrap up

- Scripts in Go and Python for data collection of Ethernet packets


---

### Thank you! :heart: 

You can find me on

- GitHub
