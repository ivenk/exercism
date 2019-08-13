class ReverseString {
    String reverse(String inputString) {
        StringBuilder strb = new StringBuilder();
        for (int i = inputString.length() -1; i >= 0; i--) {
            strb.append(inputString.charAt(i));
        }
        return strb.toString();
    }
}
