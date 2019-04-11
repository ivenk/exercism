#tried to make my code look more like python
class Matrix(object):
    def __init__(self, matrix_string):
        self.matrix = matrix_string

    def row(self, index):
        return [int(i) for i in self.matrix.split("\n")[index-1].split(" ")]
        
    def column(self, index):
        return [int(i.split(" ")[index-1]) for i in self.matrix.split("\n")]
        
