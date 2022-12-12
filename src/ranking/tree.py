#!/usr/bin/env python3

class Decisional_tree:
    def __init__(self):
        self.step = 0
        self.score = 0
        self.number_of_steps = 3
        self.last_action = "First"
        self.list_of_commands = []
        self.action_order = []
        self.commands = {
            "MALAISE": self.is_malaise,
            "SYMPTOME": self.is_symptome,
            "CARDIAC_ARREST": self.is_cardiac_arrest,
            "NONE": self.Nothing
        }

    def Nothing(self):
        self.last_action = "Nothing"

    def parse_conf(self):
        try:
            conf_file = open("commands.conf", "r")
            self.reorder_from_conf(conf_file)
        except FileNotFoundError:
            return


    def create_dictionary(self, line): #put in dictionary and after that we put it in a list of dictionary
        _dict = {}
        #command = ["MALAISE", "CARDIAC_ARREST", "SYMPTOME"]
        #if line[0] in command:
        match line[0]:
            case "MALAISE":
                _dict = {"Question": self.is_malaise, "Order": int(line[1]), "YES": self.commands[line[3]], "NO": self.commands[line[5].strip()]}
                self.list_of_commands.append(_dict)
                #self.action_order.append()
            case "CARDIAC_ARREST":
                _dict = {"Question": self.is_cardiac_arrest, "Order": int(line[1]), "YES": self.commands[line[3]], "NO": self.commands[line[5].strip()]}
                self.list_of_commands.append(_dict)
                #self.action_order.append()
            case "SYMPTOME":
                _dict = {"Question": self.is_symptome, "Order": int(line[1]), "YES": self.commands[line[3]], "NO": self.commands[line[5].strip()]}
                self.list_of_commands.append(_dict)
                #self.action_order.append()

    def reorder_from_conf(self, file): #create a list of dictionary
        line = []
        for i in file:
            line = i.split(" ")
            self.create_dictionary(line)

            
    def ask_question(self): #question to ask from order of list and reponse yes or no will auomatically ask the right question
        if (self.last_action == "First"):
            self.list_of_commands[self.step]["Question"]()
            self.step += 1
            return
        else:
            self.list_of_commands[self.step - 1][self.last_action]()
            self.step += 1
            return


    def is_malaise(self):
        print("\nDid the victim fainted ?")
        if (self.last_action == "YES" or self.last_action == "First"):
            self.score += 10


    def is_cardiac_arrest(self):
        print("\nIs the victim in cardiac arrest ?")
        if (self.last_action == "YES" or self.last_action == "First"):
            self.score += 10
            


    def is_symptome(self):
        print("\nDoes the victim have any of the following symptoms ?\n"
              "\t- Unconscious, don't speak anymore, don't open your eyes, don't watch, respond when you speak to him, reacts\n"
              "\t- Difficulty breathing, to other BP related to breathing\n"
              "\t- Signs of shock, pallor, sweating")
        if (self.last_action == "YES" or self.last_action == "First"):
            self.score += 10


    def get_score(self):
        print("Score =", self.score)


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
            if str(line_input) == "QUIT":
                loop_status = False


def main():
    tree = Decisional_tree()
    tree.parse_conf()

    tree.get_line_loop()
    print("The Score:", tree.score)
    

if (__name__ == "__main__"):
    main()
