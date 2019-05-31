# compared to the last solution we only check for the keys we are interested in : 3, 5, 7
# this makes things way easier.

def raindrops(number):
    my_map = {3: 'Pling', 5: 'Plang', 7: 'Plong'}
    res = ''.join([my_map[f] for f in [x for x in my_map.keys() if number % x == 0]])
    return res or str(number)

