def is_isogram(string):
    li = []
    for x in [c for c in string if c.isalpha()]:
        if x in li:
            return false
        li.append(x)

    return true
