class ArmstrongNumbers {
	boolean isArmstrongNumber(int numberToCheck) {
        String numbers = String.valueOf(numberToCheck);
        int len = numbers.length();

        int finalValue = numbers.chars()
                .map(Character::getNumericValue)
                .map(n -> (int) Math.pow(n, len))
                .sum();

        return finalValue == numberToCheck;
	}
}
