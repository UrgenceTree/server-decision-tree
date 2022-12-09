#!/usr/bin/env python3

class Decisional_tree:
    def __init__(self):
        self.step = 0
        self.score = 0
        self.number_of_steps = 3
        self.last_action = "Nothing"
        self.list_of_commands = []
        self.commands = {
            "MALAISE": self.is_malaise,
            "SYMPTOME": self.is_symptome,
            "CARDIAC_ARREST": self.is_cardiac_arrest,
            "NONE": "Nothing"
        }

    def parse_conf(self):
        try:
            conf_file = open("commands.conf", "r")
            self.reorder_from_conf(conf_file)
        except FileNotFoundError:
            return

    def create_dictionary(self, line): #put in dictionary and after that we put it in a list of dictionary
        _dict = {}
        match line[0]:
            case "MALAISE":
                _dict = {"Question": self.is_malaise, "Order": int(line[1]),  "YES": self.commands[line[3]], "NO": self.commands[line[5].strip()]}
                self.list_of_commands.append(_dict)
            case "CARDIAC_ARREST":
                _dict = {"Question": self.is_cardiac_arrest, "Order": int(line[1]),  "YES": self.commands[line[3]], "NO": self.commands[line[5].strip()]}
                self.list_of_commands.append(_dict)
            case "SYMPTOME":
                _dict = {"Question": self.is_symptome, "Order": int(line[1]),  "YES": self.commands[line[3]], "NO": self.commands[line[5].strip()]}
                self.list_of_commands.append(_dict)

    def reorder_from_conf(self, file): #create a list of dictionary
        line = []
        for i in file:
            line = i.split(" ")
            self.create_dictionary(line)
                
            
    def ask_question(self): #question to ask from order of list and reponse yes or no will auomatically ask the right question
        if (self.step == 0):
            self.list_of_commands[self.step]["Question"]()
            self.step += 1
            return
        else:
            self.list_of_commands[self.step - 1][self.last_action]()
            self.step += 1
            return

    def is_malaise(self):
        print("is_malaise")
        if self.last_action == "Oui":
            self.score += 10


    def is_cardiac_arrest(self):
        print("is_cardiac_arrest")
        if self.last_action == "Oui":
            self.score += 10


    def is_symptome(self):
        print("is_symptome")
        if self.last_action == "Oui":
            self.score += 10


    def get_line_loop(self):
        loop_status = True

        while loop_status:
            line_input = input()
            self.last_action = line_input
            #                                      call ask_question
            #                                      have to update self.last_action to "yes" or "NO"

            # to put in another function who will dispatch it :
            self.ask_question()

            if str(line_input) == "QUIT":
                loop_status = False


def main():
    tree = Decisional_tree()
    tree.parse_conf()

    #tree.ask_question()        #my function to test the call of function already in the order given in the conf file
    #tree.get_line_loop()
    

if (__name__ == "__main__"):
    main()
