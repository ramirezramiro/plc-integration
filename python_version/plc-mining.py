import time
import logging
from pymodbus.client.sync import ModbusTcpClient

logging.basicConfig(level=logging.ERROR)

def connect_to_ethernet_port():
    handler = ModbusTcpClient('localhost', 502)
    handler.timeout = 5
    result = handler.connect()
    if not result:
        raise ConnectionError("Error connecting to Ethernet port")
    return handler

def read_modbus_data(handler):
    client = handler
    result = client.read_coils(0, 10)
    if not result.isError():
        return result.bits
    raise Exception("Error reading Modbus data")

if __name__ == "__main__":
    try:
        handler = connect_to_ethernet_port()
        data = read_modbus_data(handler)
        print("Modbus data:", data)
    except Exception as e:
        logging.error(str(e))
