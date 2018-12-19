import string

def is_pangram(sentence):
    # build a dic containing all ascii_lowercase letters and the amount they occure
    dic = { x: False for x in string.ascii_lowercase }

    # fill the dic with values for occurences
    for c in sentence:
        # case insensitiv
        c = c.lower()
        if c in dic:
            dic[c] = True

    # count how many individual letters are used
    n = 0
    for _, v in dic.items():
        if v:
            n += 1

    # if all letters are used its true otherwise false
    if n == 26:
        return True
    else:
        return False
