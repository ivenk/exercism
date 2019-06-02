from itertools import chain

gen_first = "{} bottles of beer on the wall, {} bottles of beer.#Take one down and pass it around, {} bottles of beer on the wall."

def recite(start, take=1):
    return list(chain.from_iterable([z.split('#') for z in list(chain.from_iterable([y for y in zip([gen_first.format(x, x, x-1) for x in range(start, start-take, -1)], ['']*take)]))]))[:-1]

