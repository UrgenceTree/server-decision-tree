##
## EPITECH PROJECT, 2022
## ranking [WSL: Ubuntu-22.04]
## File description:
## questions
##

class Questions:
    def Malaise(self, tree):
        print("\nDid the victim fainted ?")
        if (tree.last_action == "YES" or tree.last_action == "First"):
            tree.score += 10

    def Cardiac_arrest(self, tree):
        print("\nIs the victim in cardiac arrest ?")
        if (tree.last_action == "YES" or tree.last_action == "First"):
            tree.score += 10

    def Symptome(self, tree):
        print("\nDoes the victim have any of the following symptoms ?\n"
              "\t- Unconscious, don't speak anymore, don't open your eyes, don't watch, respond when you speak to him, reacts\n"
              "\t- Difficulty breathing, to other BP related to breathing\n"
              "\t- Signs of shock, pallor, sweating")
        if (tree.last_action == "YES" or tree.last_action == "First"):
            tree.score += 10
    
    def Nothing(self, tree):
        tree.last_action = "Nothing"