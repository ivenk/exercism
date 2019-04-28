def is_isogram(string):
    li = []
    for x in [c for c in string if c.isalpha()]:
        x = x.lower()
        if x in li:
            return False
        li.append(x)

    return True
