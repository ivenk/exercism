class HighScores(object):
    def __init__(self, scores):
        self.scores = scores

    def latest(self):
        return self.scores[len(self.scores) - 1]

    def personal_best(self):
        best = 0
        for score in self.scores:
            if score > best:
                best = score
        return best

    def personal_top_three(self):
        self.scores.sort(reverse=True)
        return self.scores[:3]
        
        
