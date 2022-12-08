#!/usr/bin/env python3

import json
import sys

class JsonFile:
    def put_in_dictionary(self, _id, name, age, city, description, rank):
        null = None
        dictionary = {"id": _id, "name": name, "age": age, "city": city, "description": description, "rank": rank, "user_id": null, "create_at": ""}
        self.put_in_json_file(dictionary)
    
    def put_in_json_file(self, dictionary):
        file = open("data.json", "w") 
        jsonString = json.dumps(dictionary, indent=4)
        print(jsonString, file=file)


def main():
    js = JsonFile()
    test = js.put_in_dictionary("0001", "Ashwin SAGODIRA", 20, "KB", "I am a student", 1)

if (__name__ == "__main__"):
    main()