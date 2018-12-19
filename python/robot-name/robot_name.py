import random
import string

# list of all names currently in use
used_names = []

class Robot(object):
    name = ""
    
    def __init__(self):
        self.name = self._generate_name()

    # generates a unique name
    def _generate_name(self):        
        name = ""
        while(name == "" or name in used_names):
            # reset name
            name = ""
            
            # add 2 random letters
            for x in range(2):
                name += random.choice(string.ascii_uppercase)
        
            # add 3 random numbers
            for x in range(3):
                name += str(random.randint(0,9))
            pass
        
        used_names.append(name)
        return name

    # gives a robot a new name and frees the old one afterwards
    def reset(self):
        old_name = self.name
        self.name = self._generate_name()
        used_names.remove(old_name)

    # on destruction of object we free the name
    def __del__(self):
        used_names.remove(self.name) 
