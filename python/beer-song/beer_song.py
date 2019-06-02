from itertools import chain

gen_first = "{} bottles of beer on the wall, {} bottles of beer.#Take one down and pass it around, {} bottles of beer on the wall."

def recite(start, take=1):
    return list(chain.from_iterable([z for z in [y.split('#') for y in [gen_first.format(x, x, x-1) for x in range(start, start-take, -1)]]]))

