def proteins(strand):
    dic = {
        "AUG" : "Methionine",
        "UUU" : "Phenylalanine",
        "UUC" : "Phenylalanine",
        "UUA" : "Leucine",
        "UUG" : "Leucine",
        "UCU" : "Serine",
        "UCC" : "Serine",
        "UCA" : "Serine",
        "UCG" : "Serine",
        "UAU" : "Tyrosine",
        "UAC" : "Tyrosine",
        "UGU" : "Cysteine",
        "UGC" : "Cysteine",
        "UGG" : "Tryptophan",
        "UAA" : "STOP",
        "UAG" : "STOP",
        "UGA" : "STOP",
    }

    increase = 3
    counter = 3

    result = []
    
    while counter <= (len(strand)):
        c = strand[(counter-increase):counter]
        v = dic[c]
        if v == None :
            raise Exception("No match found for %s", c)
        elif v == "STOP":
            break;
        else:
            result.append(dic[c])
        
        counter += increase

    return result
        
    
