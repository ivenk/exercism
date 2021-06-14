class Matrix(private val matrixAsString: String) {

    private val matrix : List<List<Int>> = matrixAsString
            .lines()
            .map{
                it
                        .trim()
                        .replace("\\s+".toRegex(), " ")
                        .split(" ")
                        .map { it.toInt() }
            }

    fun column(colNr: Int): List<Int> {
        return matrix.map { it[colNr-1] }
    }

    fun row(rowNr: Int): List<Int> {
        return matrix[rowNr-1]
    }
}
