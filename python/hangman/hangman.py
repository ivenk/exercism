# Game status categories
# Change the values as you see fit
STATUS_WIN = "win"
STATUS_LOSE = "lose"
STATUS_ONGOING = "ongoing"


class Hangman(object):
    def __init__(self, word):
        self.remaining_guesses = 9
        self.status = STATUS_ONGOING
        self.word = word
        self.letters = []
        self.found_letters = []
        
        # we find all letters 
        for letter in word:
                if letter not in self.letters:
                    self.letters.append(letter)
               
    def guess(self, char):
        if self.status == STATUS_WIN :
            raise ValueError("ValueError")
        
        if self.remaining_guesses >= 0:
            if (char in self.letters) and (char not in self.found_letters) :
                self.found_letters.append(char)
                if len(self.letters) == len(self.found_letters):
                    self.status = STATUS_WIN
            else:
                self.remaining_guesses -= 1
                if self.remaining_guesses == 0:
                    self.status = STATUS_LOSE
        else:
            raise ValueError("ValueError")

        
    def get_masked_word(self):
        str = ""
        for l in self.word:
            if l in self.found_letters:
                str += l
            else:
                str += "_"
        return str

    def get_status(self):
        return self.status
