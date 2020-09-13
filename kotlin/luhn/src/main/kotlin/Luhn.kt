object Luhn {

    fun isValid(candidate: String): Boolean {
        if (candidate.trim().length < 2) return false
        var trimed = candidate.replace(" ", "")
        if(trimed.any {!it.isDigit()}) {
            return false
        }

        return trimed
                .filter { it.toString().toIntOrNull() != null }
                .map { it.toString().toInt() }
                .reversed()
                .withIndex()
                .map { if (it.index.rem(2) == 1) it.value.doubleWithCap() else it.value }
                .sum().rem(10) == 0
    }

    private fun Int.doubleWithCap(): Int {
        val doubled = this.times(2)
        return if (doubled > 9) doubled - 9 else doubled
    }

}
