from itertools import chain

plant_dict = {
    'V':'Violets',
    'R':'Radishes',
    'G':'Grass',
    'C':'Clover',
}

class Garden(object):
    def __init__(self, diagram, students="Alice Bob Charlie David Eve Fred Ginny Harriet Ileana Joseph Kincaid Larry".split()):
        self.garden = diagram.split('\n')
        self.students = sorted(students)
        
    def plants(self, student):
        return list(chain.from_iterable([[plant_dict[x] for x in g[self.students.index(student) * 2: self.students.index(student) * 2 +2]] for g in self.garden]))
        
        
        
