class Darts {
    int score;

    Darts(double x, double y) {
        double distance = pythagoras(x, y);

        if (distance <= 1) {
            score = 10;
        } else if (distance <= 5) {
            score = 5;
        } else if (distance <= 10) {
            score = 1;
        } else {
            score = 0;
        }
    }

    private double pythagoras(double x, double y) {
        return Math.sqrt(Math.pow(x, 2) + Math.pow(y, 2));
    }

    int score() {
        return score;
    }

}
