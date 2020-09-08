class CustomSet(content: List<Int>) {

    var content = content.toMutableList()

    constructor(vararg input: Int) : this(input.toList())

    fun isEmpty(): Boolean {
        return content.isEmpty()
    }

    fun isSubset(other: CustomSet): Boolean {
        return content.containsAll(other.content)
    }

    fun isDisjoint(other: CustomSet): Boolean {
        return !content.any { other.contains(it) }
    }

    fun contains(other: Int): Boolean {
        return content.contains(other)
    }

    fun intersection(other: CustomSet): CustomSet {
        return CustomSet(content.filter { other.contains(it) });
    }

    fun add(other: Int) {
        content.add(other)
    }

    override fun equals(other: Any?): Boolean {
        println(content.toList())
        if (other is CustomSet) println(other.content.toList())

        return if (other is CustomSet) other.content == this.content else false
    }

    operator fun plus(other: CustomSet): CustomSet {
        val newList = content.toMutableList()
        newList.addAll(other.content)
        return CustomSet().apply { this.content = newList }
    }

    operator fun minus(other: CustomSet): CustomSet {
        val newList = content.toMutableList()
        newList.removeAll(other.content)
        return CustomSet().apply { this.content = newList }
    }
}
