def is_isogram(string):
    li = []
    for x in [c for c in string.lower() if c.isalpha()]:
        if x in li:
            return False
        li.append(x)

    return True

    
