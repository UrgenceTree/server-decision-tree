#!/usr/bin/env python3

class Decisional_tree:
    def __init__(self):
        self.step = 0
        self.number_of_steps = 3
        self.score = 0
        self.commands = {
            0: self.is_malaise,
            1: self.is_symptome,
            2: self.is_cardiac_arrest,
        }

    def parse_conf(self):
        try:
            conf_file = open("commands.conf", "r")
            self.commands.clear()
            self.reorder_from_conf(conf_file)
        except FileNotFoundError:
            return

    def reorder_from_conf(self, file):
        for i in file:
            match i.split(" ")[0]:
                case "MALAISE":
                    self.commands[int(i.split(" ")[1])] = self.is_malaise
                case "CARDIAC_ARREST":
                    self.commands[int(i.split(" ")[1])] = self.is_cardiac_arrest
                case "SYMPTOME":
                    self.commands[int(i.split(" ")[1])] = self.is_symptome
                case "NB_COMMANDS":
                    self.number_of_steps = int(i.split(" ")[1])

    def is_malaise(self, line_input):
        if str(line_input) == "Oui":
            self.score += 10


    def is_cardiac_arrest(self, line_input):
        if line_input == "Oui":
            self.score += 10


    def is_symptome(self, line_input):
        if line_input == "Oui":
            self.score += 10


    def get_line_loop(self):
        loop_status = True

        while loop_status:
            line_input = input()

            # to put in another function who will dispatch it :
            self.is_cardiac_arrest(line_input)
            self.is_symptome(line_input)
            self.is_malaise(line_input)

            if str(line_input) == "QUIT":
                loop_status = False


def main():
    tree = Decisional_tree()
    tree.parse_conf()
    
    tree.get_line_loop()
    
    #for i in range (tree.number_of_steps):
    #    tree.commands[i]()
    

if (__name__ == "__main__"):
    main()
