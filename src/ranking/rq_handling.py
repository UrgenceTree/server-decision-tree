##
## EPITECH PROJECT, 2022
## server-decision-tree [WSL: Ubuntu-22.04]
## File description:
## requests
##

import requests

class request:
    def __init__(self):
        self.header = {"Content-Type": "application/json"}
        self.data = {'name': "Unknown", 'age': "18", "situation": ".", "rank": 1, "status": False, "telephone": "0", "location": "Paris", "chunk_id": "0"}
        self.url_patient = "http://localhost:8002/api/patient-info"
        self.url_status = "http://localhost:8002/api/patient-info/status/"
        self.url_situation = "http://localhost:8002/api/patient-info/situation/"
        self.id = ""
        self.step = 0
    
    def create_patient(self):
        post = requests.post(self.url_patient, json=self.data, headers=self.header)
        response_in_json = post.json()
        self.id = str(response_in_json["data"]["_id"])
        self.url_status += self.id
        self.url_situation += self.id
    
    def update_status(self):
        self.data["status"] = True
        data = {"status": True}
        requests.put(self.url_status, json=data, headers=self.header)
    
    def update_situation(self, string):
        if self.step == 0:
            self.data["situation"] = string
        else:
            self.data["situation"] += ", " + string
        data = {"situation": self.data["situation"]}
        requests.put(self.url_situation, json=data, headers=self.header)
        self.step += 1
    
    def update_score(self, nb):
        self.data["rank"] = nb
        data = {"rank": self.data["rank"]}
        requests.put(self.url_situation, json=data, headers=self.header)