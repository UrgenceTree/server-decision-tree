##
## EPITECH PROJECT, 2022
## ranking [WSL: Ubuntu-22.04]
## File description:
## questions
##

class Questions:
    def __init__(self):
        self.answers_yes = ["yes", "Yes.", "absolutely", "Absolutely." ,"indeed", "Indeed.", "totally", "Totally."]
        self.answers_no = ["no", "No.", "not", "Not.", "is not", "He's not.", "he's not" ]


    def Malaise(self, tree):
        while 1:
            # print("\nDid the victim fainted ?")
            if tree.answer_data[0] in self.answers_yes:
                break;
            if tree.answer_data[0] in self.answers_no:
                break;
        if (tree.answer_data[0] in self.answers_yes):
            tree.score += 10
            tree.last_action = "Yes."
            # rq.update_situation("Malaise")
            # rq.update_score(tree.score)
        else:   
            tree.last_action = "No."
            
            

    def Cardiac_arrest(self, tree):
        while 1:
            # print("\nIs the victim in cardiac arrest ?")
            if tree.answer_data[1] in self.answers_yes:
                break;
            if tree.answer_data[1] in self.answers_no:
                break;
        if tree.answer_data[1] in self.answers_yes:
            tree.score += 100
            tree.last_action = "Yes."
            # rq.update_situation("Cardiac arrest")
            # rq.update_score(tree.score)
        else:
            tree.last_action = "No."
            
        

    def Symptome(self, tree):
        while 1:
            # print("\nDoes the victim have any of the following symptoms ?\n"
            #   "\t- Unconscious, don't speak anymore, don't open your eYes., don't watch, respond when you speak to him, reacts\n"
            #   "\t- Difficulty breathing, to other BP related to breathing\n"
            #   "\t- Signs of shock, pallor, sweating")
            if tree.answer_data[2] in self.answers_yes:
                break;    
            if tree.answer_data[2] in self.answers_no:
                break;

        if (tree.answer_data[2] in self.answers_yes):
            tree.score += 10
            tree.last_action = "Yes." 
            # rq.update_situation("Symptome")
            # rq.update_score(tree.score)
        else:
            tree.last_action = "No."
            
    
    def Nothing(self, tree):
        tree.last_action = "Nothing"