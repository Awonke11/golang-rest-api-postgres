import requests
import json

def getData(url):
    response_api = requests.get(url)
    return response_api.json()

print(getData("http://127.0.0.1:8080/api/users"))