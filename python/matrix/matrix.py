class Matrix(object):
    def __init__(self, matrix_string):
        # the matrix represented by a list of rows which are represented by a list aswell
        self.matrix_list = [[int(y) for y in x.split(" ")] for x in matrix_string.split("\n")]
        
    def row(self, index):
        # copy : return by value instead of reference
        return self.matrix_list[index-1].copy()
        
    def column(self, index):
        # should always have the correct order since lists are ordered
        return [z[index-1] for z in self.matrix_list]
        
