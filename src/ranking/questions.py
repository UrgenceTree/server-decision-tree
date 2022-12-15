##
## EPITECH PROJECT, 2022
## ranking [WSL: Ubuntu-22.04]
## File description:
## questions
##

class Questions:
    def __init__(self):
        self.answers_yes = ["yes", "Yes."]
        self.answers_no = ["no", "No."]
        


    def Malaise(self, tree, rq):
        while 1:
            print("\nDid the victim fainted ?")     
            line = input().strip()
            if line in self.answers_yes:
                break;
            if line in self.answers_no:
                break;
        if (line in self.answers_yes):
            tree.score += 10
            tree.last_action = "Yes."
            rq.update_situation("Malaise")
        else:   
            tree.last_action = "No."
            
            

    def Cardiac_arrest(self, tree, rq):
        while 1:
            print("\nIs the victim in cardiac arrest ?")     
            line = input().strip()
            if line in self.answers_yes or self.answers_no:
                break;
        if line in self.answers_yes:
            tree.score = 100
            tree.last_action = "Yes."
            rq.update_situation("Cardiac arrest")
        else:    
            tree.last_action = "No."
            
        

    def Symptome(self, tree, rq):
        while 1:
            print("\nDoes the victim have any of the following symptoms ?\n"
              "\t- Unconscious, don't speak anymore, don't open your eYes., don't watch, respond when you speak to him, reacts\n"
              "\t- Difficulty breathing, to other BP related to breathing\n"
              "\t- Signs of shock, pallor, sweating")     
            line = input().strip()
            if line in self.answers_yes:
                break;    
        
        if (line in self.answers_yes):
            tree.score += 10
            tree.last_action = "Yes." 
            rq.update_situation("Symptome")
        else:   
            tree.last_action = "No."
            
    
    def Nothing(self, tree, rq):
        tree.last_action = "Nothing"