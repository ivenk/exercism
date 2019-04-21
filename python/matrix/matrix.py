class Matrix(object):
    def __init__(self, matrix_string):
        # the matrix represented by a list of rows which are represented by a list aswell
        self.matrix_list = []
        for x in matrix_string.split("\n"):
            self.matrix_list.append([int(y) for y in x.split(" ")])
        
    def row(self, index):
        return self.matrix_list[index-1]
        
    def column(self, index):
        # should always have the correct order since lists are ordered
        return [z[index-1] for z in self.matrix_list]
        
