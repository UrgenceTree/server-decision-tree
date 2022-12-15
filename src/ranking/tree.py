#!/usr/bin/env python3

from questions import *
import sys
import requests
import configparser

class Decisional_tree:
    def __init__(self):
        self.qt = Questions()
        self.step = 0
        self.score = 0
        self.number_of_steps = 3
        self.last_action = "First"
        self.list_of_commands = []
        self.action_order = []
        self.serv_data_ip = ""
        self.serv_data_port = 0


    def init_env(self):
        config = configparser.ConfigParser()
        config.read('./env/serv_data_info.env')
        self.serv_data_ip = config['server_data_info']['SERV_DATA_IP']
        self.serv_data_port = config['server_data_info']['SERV_DATA_PORT']

    def handle_env(self):
        self.init_env()
        if ((int(self.serv_data_port) < 1025) or (int(self.serv_data_port) > 65535)):
            print("ERROR: The PORT in the env must be between 1025 and 65535", file=sys.stderr)
            exit(84)
        if (self.serv_data_ip == None):
            print("ERROR: Bad ip adress", file=sys.stderr)
            exit(84)
        #print("serv_data_ip =", self.serv_data_ip)
        #print("serv_data_port =", self.serv_data_port)

    def parse_conf(self):
        conf_file = open("../server-decision-tree/src/ranking/commands.conf", "r")
        self.reorder_from_conf(conf_file)

    def reorder_from_conf(self, file): #create a list of dictionary
        line = []
        for i in file:
            line = i.split(" ")
            self.create_dictionary(line)

    def create_dictionary(self, line): #put in dictionary and after that we put it in a list of dictionary
        _dict = {}
        
        if line[0] == "NB_COMMANDS":
            self.number_of_steps = int(line[1])
            return
        else:
            _dict = {"Question": getattr(self.qt, line[0].capitalize()), "Order": int(line[1]), "Yes.": getattr(self.qt, line[3].capitalize()), "No.": getattr(self.qt, line[5].strip().capitalize())}
            self.list_of_commands.append(_dict)
            
    def ask_question(self): #question to ask from order of list and reponse yes or no will auomatically ask the right question
        if (self.last_action == "First"):
            self.list_of_commands[self.step]["Question"](self)
            self.step += 1
            return
        else:
            self.list_of_commands[self.step - 1][self.last_action](self)
            self.step += 1
            return

    def get_line_loop(self):
        loop_status = True

        while loop_status:
            # Call ask_question
            # Have to update self.last_action to "yes" or "NO"
            self.ask_question()
            if self.last_action == "Nothing":
                break
            if self.last_action == "Yes." or self.last_action == "No.":
                self.last_action = self.last_action
            else:
                self.step -= 1
                self.score -= 10
            if str(self.last_action) == "QUIT":
                loop_status = False


def main():
    tree = Decisional_tree()
    # try:
    tree.parse_conf()
    tree.get_line_loop()
    
    # The function 'handle_env' get and set variables env to the class 'Decisional_tree'
    # and handle errors from env variables set in env/server_data_info.env
    #tree.handle_env()

    #     print("\nThe Score:", tree.score)
    # except KeyboardInterrupt:
    #     print("ERROR: Keyboard Interrupt")
    # except EOFError:
    #     print("ERROR: End of File")
    # except FileNotFoundError:
    #     print("ERROR: File not found")
    # except AttributeError:
    #     print("ERROR: Function given doesn't exist")

if (__name__ == "__main__"):
    main()
