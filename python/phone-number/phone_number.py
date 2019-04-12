class Phone(object):
    def __init__(self, phone_number):
        # we use list comprehension in order to remove all chars that are not numeric        
        self.number = ''.join([i for i in phone_number if i.isdigit()])
        # remove leading nation code
        if len(self.number) == 11 :
            print("maybe here ?")
            if self.number[0] == "1":
                self.number = self.number[1:]
                print("i get called")

        # not the correct length
        if len(self.number) != 10:
            raise ValueError("ValueError")

        # check if the first digit in area code and in region code is bigger then 1
        # we know its a digit therefore casting should always work
        if (int(self.number[0]) <= 1) or (int(self.number[3]) <= 1):
            raise ValueError("ValueError")

        self.area_code = self.number[:3]

    def pretty(self):
        return "(%s) %s-%s" % (self.number[:3], self.number[3:6], self.number[6:])
        
        
