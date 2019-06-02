from itertools import chain

class School(object):
    def __init__(self):
        self.r = {}
        for x in range(1, 8):
            self.r[x] = []

    def add_student(self, name, grade):
        self.r[grade].append(name)

    def roster(self):
        # iterate over the keys, get the entries for each key and sort them. After that we flatten it all to one list
        return list(chain.from_iterable([sorted(self.r[x]) for x in self.r])) 

    def grade(self, grade_number):
        return sorted(self.r[grade_number])
