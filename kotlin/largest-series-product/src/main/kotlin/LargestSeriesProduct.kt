class Series(series: String) {

    private val numbers = series.trim().map { it.toString().toInt() }

    fun getLargestProduct(span: Int): Long {
        if (span > numbers.size) throw IllegalArgumentException("Fail")

        if (span == 0) return numbers.getOrElse(0) { 1 }.toLong()

        var max = 0L
        numbers.forEachIndexed { index, _ ->
            if (index + span > numbers.size) return@forEachIndexed
            val sum = numbers.subList(index, index + span).reduce { x, y -> x * y }.toLong()
            max = maxOf(max, sum)
        }
        return max;
    }
}
