/**
 * Solution inspired by Testare's and hartror's solution.
 */
object Bob {
    fun hey(input: String): String {
        val trim = input.trim()

        if (trim.isEmpty()) return "Fine. Be that way!"

        return when(Pair(isQuestion(trim), isYelling(trim))) {
            Pair(true, false) -> "Sure."
            Pair(true, true) -> "Calm down, I know what I'm doing!"
            Pair(false, true) -> "Whoa, chill out!"
            else -> "Whatever."
        }
    }

    private fun isYelling(message : String) : Boolean {
        return message.any { it.isLetter() } && message == message.toUpperCase()
    }

    private fun isQuestion(message: String) : Boolean{
        return message.endsWith("?")
    }
}
