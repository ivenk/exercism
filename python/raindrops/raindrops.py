def raindrops(number):
    # +1 due to range exluding the max number; at the end we need to append the number itself
    factors = [x for x in range(1, (int(number/2) + 1)) if number % x == 0]
    factors.append(number)
    res = ""
    if 3 in factors:
        res += "Pling"
    if 5 in factors:
        res += "Plang"
    if 7 in factors:
        res += "Plong"

    if res == "":
        return str(number)
        
    return res
