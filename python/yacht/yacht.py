# Score categories
# Change the values as you see fit
YACHT = "YACHT"
ONES = "ONES"
TWOS = "TWOS"
THREES = "THREES"
FOURS = "FOURS"
FIVES = "FIVES"
SIXES = "SIXES"
FULL_HOUSE = "FULL_HOUSE"
FOUR_OF_A_KIND = "FOUR_OF_A_KIND"
LITTLE_STRAIGHT = "LITTLE_STRAIGHT"
BIG_STRAIGHT = "BIG_STRAIGHT"
CHOICE = "CHOICE"


def score(dice, category):
    if len(dice) != 5:
        raise Exception("Need 5 guesses")
    
    if category == YACHT:
        return __yacht(dice)
    elif category == ONES:
        return __number(dice, 1)
    elif category == TWOS:
        return __number(dice, 2)
    elif category == THREES:
        return __number(dice, 3)
    elif category == FOURS:
        return __number(dice, 4)
    elif category == FIVES:
        return __number(dice, 5)
    elif category == SIXES:
        return __number(dice, 6)
    elif category == FULL_HOUSE:
        return __full_house(dice)
    elif category == FOUR_OF_A_KIND:
        return __four_of_a_kind(dice)
    elif category == LITTLE_STRAIGHT:
        return __little_straight(dice)
    elif category == BIG_STRAIGHT:
        return __big_straight(dice)
    elif category == CHOICE:
        return __choice(dice)

    return 0

# returns the points the dice scored in the yacht category
def __yacht(dice):
    if dice.count(dice[0]) == 5:
        return 50
    else:
        return 0
    
# returns the points scored in the number categorie identified by the number given
def __number(dice, number):
    return dice.count(number) * number

# returns the points scored in the full house categorie
def __full_house(dice):
    f = 0
    s = 0
    u = 0
    
    for x in dice:
        # count occurence of first dice row
        if dice[0] == x:
            f += 1

        # if we find a different dice value we store it
        # in failing cases this might get reasigned a couple of times but should not be an issue
        if x != dice[0]:
            u = x

    # we count the occurence of the second dice value
    s = dice.count(u)
    
    # if we have a full house
    if (f == 2 and s == 3) or (f == 3 and s == 2):
        return sum(dice)
                    
    return 0

def __four_of_a_kind(dice):
    # this looks strange but i feel like its the most efficient way of doing this
    # we count the occurence of the value of the first dice throw
    # we return the points if we found our 4 elements
    if (dice.count(dice[0]) >= 4): 
        return dice[0] * 4

    # if the first element did not occure 4 times the second one has to
    # if it doesnt there cannot be four of a kind in a list of len 5
    if (dice.count(dice[1]) >= 4):
        return dice[1] * 4
      
    return 0


def __little_straight(dice):
    # always wanted to use list comprehension
    if sorted(dice) == [x for x in range(1, 6)]:
        return 30
    return 0 

def __big_straight(dice):
    if sorted(dice) == [x for x in range(2, 7)]:
        return 30
    return 0

def __choice(dice):
    return sum(dice)
    
