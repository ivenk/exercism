/* Solution inspired by the amazing solutions of GrahamsLea and cowlike */
class Series(val series: String) {
    init {
        // only accepting numerical strings
        require(series.all { it.isDigit() })
    }

    fun getLargestProduct(span: Int): Long {
        return when {
            span > series.length -> throw IllegalArgumentException("fail")
            span == 0 -> 1L
            else -> series.windowed(span, 1, false) {
                it.fold(1L) { x, y -> x * Character.getNumericValue(y) }
            }.max()?: 1L
        }
    }
}
