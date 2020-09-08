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
        if (other is CustomSet) return other.content.containsAll(content) && content.containsAll(other.content) else return false
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
