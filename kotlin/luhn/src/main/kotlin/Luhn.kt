object Luhn {

    fun isValid(candidate: String): Boolean {
        if (candidate.length < 2) return false

        return candidate
                .replace(" ", "")
                .map { it.toString().toInt() }
                .withIndex()
                .map { if (it.index.rem(2) == 0) it.value.doubleWithCap() else it.value }
                .sum().rem(10) == 0
    }

    fun Int.doubleWithCap(): Int {
        val doubled = this.times(2)
        return if (doubled > 9) doubled - 9 else doubled
    }
}
