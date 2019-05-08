from collections import defaultdict # defaultdict makes incrementing way nicer
import re

def word_count(phrase):
    word_dict = defaultdict(int)
    words_raw = re.split("(\\n)|[,]| |_|\t", phrase)
    words = [w.strip().strip("@^':.!%$&§").lower() for w in words_raw if w != '' and w != ' ' and w != '\n' and w!= None]
    print(words)
    for w in words:
        if w != '' and w != ' ' and w != '\n' and w!= None:
            word_dict[w] += 1 

    return dict(word_dict)
