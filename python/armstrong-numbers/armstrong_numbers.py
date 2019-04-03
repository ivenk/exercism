def is_armstrong(number):
    x = str(number)
    power = len(x)

    value = 0
    for n in x:
        value += int(n)** power

    return value == number
    
