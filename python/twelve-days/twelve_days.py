song_dic = {
    1:("first", "a Partridge in a Pear Tree"),
    2:("second", "two Turtle Doves"),
    3:("third", "three French Hens"),
    4:("fourth", "four Calling Birds"),
    5:("fifth", "five Gold Rings"),
    6:("sixth", "six Geese-a-Laying"),
    7:("seventh", "seven Swans-a-Swimming"),
    8:("eighth", "eight Maids-a-Milking"),
    9:("ninth", "nine Ladies Dancing"),
    10:("tenth", "ten Lords-a-Leaping"),
    11:("eleventh", "eleven Pipers Piping"),
    12:("twelfth", "twelve Drummers Drumming"),
}

default_line = "On the {} day of Christmas my true love gave to me:"

def recite(start_verse, end_verse):
    lines = []
    while start_verse <= end_verse:
        end_list = [(song_dic[x])[1] for x in range(start_verse, 0, -1)]
        if len(end_list) > 1 :
            end_list[-1] = "and {}".format(end_list[-1]) # last item needs and and .
            
        end_list[-1] = "{}.".format(end_list[-1]) # last item needs .
        endings = ', '.join(end_list)
        line = "{} {}".format(default_line.format(song_dic[start_verse][0]), endings)
        lines.append(line)
        start_verse += 1
        
    return lines 


