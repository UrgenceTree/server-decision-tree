##
## EPITECH PROJECT, 2022
## ranking [WSL: Ubuntu-22.04]
## File description:
## questions
##

class Questions:
    def __init__(self):
        self.answers = ["Yes.", "No."]


    def Malaise(self, tree):
        while 1:
            print("\nDid the victim fainted ?")     
            line = input().strip()
            print("Line:", line)
            if line in self.answers:
                break;
        if (line == "Yes."):
            tree.score += 10
        tree.last_action = line
            

    def Cardiac_arrest(self, tree):
        while 1:
            print("\nIs the victim in cardiac arrest ?")     
            line = input().strip()
            if line in self.answers:
                break;
        if line == "Yes.":
            tree.score = 100
        tree.last_action = line
            
        

    def Symptome(self, tree):
        while 1:
            print("\nDoes the victim have any of the following symptoms ?\n"
              "\t- Unconscious, don't speak anymore, don't open your eYes., don't watch, respond when you speak to him, reacts\n"
              "\t- Difficulty breathing, to other BP related to breathing\n"
              "\t- Signs of shock, pallor, sweating")     
            line = input().strip()
            if line in self.answers:
                break;    
        
        if (line == "Yes."):
            tree.score += 10
        tree.last_action = line
            
    
    def Nothing(self, tree):
        tree.last_action = "Nothing"