class ArmstrongNumbers {
	boolean isArmstrongNumber(int numberToCheck) {
		int len = String.valueOf(numberToCheck).length();
		int number = numberToCheck;
		int result = 0;

		for(; number > 0; number = number/10) {
		    result += Math.pow((number % 10), len);
        }
		return (result == numberToCheck);
	}
}
