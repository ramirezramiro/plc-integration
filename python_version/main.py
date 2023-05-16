from pymongo import MongoClient
from pymongo.errors import ConnectionError
from modbus_reader import read_modbus_data

def send_data_to_mongodb(data):
    try:
        client = MongoClient('mongodb://localhost:27017')
        db = client['test']
        collection = db['modbus_data']
        result = collection.insert_one({"data": data})
        if not result.acknowledged:
            raise Exception("Error sending data to MongoDB")
    except ConnectionError as e:
        raise Exception("Error connecting to MongoDB") from e

if __name__ == "__main__":
    try:
        data = read_modbus_data()
        send_data_to_mongodb(data)
        print("Modbus data saved to MongoDB successfully!")
    except Exception as e:
        print(str(e))
