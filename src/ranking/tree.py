#!/usr/bin/env python3

from questions import *

class Decisional_tree:
    def __init__(self):
        self.qt = Questions()
        self.step = 0
        self.score = 0
        self.number_of_steps = 3
        self.last_action = "First"
        self.list_of_commands = []
        self.action_order = []

    def parse_conf(self):
        try:
            conf_file = open("commands.conf", "r")
            self.reorder_from_conf(conf_file)
        except FileNotFoundError:
            return
        except AttributeError:
            print("ERROR: Function given doesn't exist")

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
            _dict = {"Question": getattr(self.qt, line[0].capitalize()), "Order": int(line[1]), "YES": getattr(self.qt, line[3].capitalize()), "NO": getattr(self.qt, line[5].strip().capitalize())}
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
            line_input = input()
            if line_input == "YES" or line_input == "NO":
                self.last_action = line_input
            else:
                self.step -= 1
                self.score -= 10
            if str(line_input) == "QUIT":
                loop_status = False


def main():
    tree = Decisional_tree()
    tree.parse_conf()
    
    tree.get_line_loop()
    print("The Score:", tree.score)

if (__name__ == "__main__"):
    main()
