#tried to make my code look more like python
class Matrix(object):
    def __init__(self, matrix_string):
        self.matrix_list = []
        for x in matrix_string.split("\n"):
            self.matrix_list.append([int(y) for y in x.split(" ")])
        
    def row(self, index):
        return self.matrix_list[index-1]
        
    def column(self, index):
        return
        
