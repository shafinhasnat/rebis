import socket

class RebisClient:
    def __init__(self, host: str, port: int) -> None:
        self.sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.sock.connect((host, port))
        print("Connected to server")
    def rebisQuery(self, query:str) -> str:
        self.sock.send(query.encode())
        response = self.sock.recv(1024).decode()
        response = response[2:len(response)-1]
        return response
class RebisQuery:
    def __init__(self, client: RebisClient) -> None:
        self.client = client
    def set(self, key: str, value: str) -> str:
        query = f"SET {key} {value}"
        response = self.client.rebisQuery(query)
        return response
    def get(self, key: str) -> str:
        query = f"GET {key}"
        response = self.client.rebisQuery(query)
        return response
    def delete(self, key: str) -> str:
        query = f"DELETE {key}"
        response = self.client.rebisQuery(query)
        return response
    def reset(self) -> str:
        query = f"RESET"
        response = self.client.rebisQuery(query)
        return response

