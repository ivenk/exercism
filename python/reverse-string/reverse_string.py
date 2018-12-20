
def reverse(text):
    i = len(text) - 1
    new_text = ""
    while i >= 0:
        new_text += text[i]
        i = i -1

    return new_text
