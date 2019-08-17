import java.util.HashMap;
import java.util.Map;

class RnaTranscription {
    private static final Map<Character, Character> map;
    static {
    	 map = new HashMap<>();
    	 map.put('G', 'C');
    	 map.put('C', 'G');
    	 map.put('T', 'A');
    	 map.put('A', 'U');
	}

    public String transcribe(String dnaStrand) {
	char[] solution = new char[dnaStrand.length()];

	if (map == null) {
		System.out.println("test");
	}

	for (int i = 0; i< dnaStrand.length(); i++) {
	    solution[i] = map.get(dnaStrand.toCharArray()[i]);
	}
	return new String(solution);
    }
}
