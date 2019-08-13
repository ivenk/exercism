import java.util.Optional;

// new solution after looking at frenata's amazing solution
class Twofer {
    String twofer(String name) {
	return String.format("One for %s, one for me.", Optional.ofNullable(name).orElse("you"));
    }
}
