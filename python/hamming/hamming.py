def distance(strand_a, strand_b):
    distance = 0

    if len(strand_a) != len(strand_b):
        raise ValueError("ValueError")

    # enumerate returns the index and the value
    for i, x in enumerate(strand_a):
        if x != strand_b[i]:
            distance += 1

    return distance
