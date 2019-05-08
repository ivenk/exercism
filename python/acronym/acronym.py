import string
import re

def abbreviate(words):
    return ''.join([w.strip('_')[0].upper() for w in re.findall(r"[\w']+", words)])
    # return ''.join([w.strip('_')[0].upper() for w in words.replace(' -', '').split()])
    
