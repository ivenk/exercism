class CustomSet(content: List<Int>) {

    var content = content.toMutableList()

    constructor(vararg input: Int) : this(input.toList())

    fun isEmpty(): Boolean {
        return content.isEmpty()
    }

    fun isSubset(other: CustomSet): Boolean {
        return other.content.containsAll(content)
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
        return if (other is CustomSet) other.content.containsAll(content) && content.containsAll(other.content) else false
    }

    operator fun plus(other: CustomSet): CustomSet {
        return CustomSet(content + other.content)
    }

    operator fun minus(other: CustomSet): CustomSet {
        return CustomSet(content - other.content)
    }
}
