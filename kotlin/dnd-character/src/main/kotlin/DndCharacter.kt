import kotlin.math.roundToInt
import kotlin.random.Random

class DndCharacter {

    val strength: Int = ability()
    val dexterity: Int = ability()
    val constitution: Int = ability()
    val intelligence: Int = ability()
    val wisdom: Int = TODO("Initialize value to complete the task")
    val charisma: Int = TODO("Initialize value to complete the task")
    val hitpoints: Int = TODO("Initialize value to complete the task")

    companion object {

        fun ability(): Int {
            val rolles = mutableListOf<Int>()
            repeat(4) {
                rolles.add(rollDice())
            }
            val sum = rolles.sorted().dropLast(1).sum()
            return sum
        }

        fun modifier(score: Int): Int {
            val div = score.minus(10).div(2)
            return div
        }

        private fun rollDice() : Int = Random.nextInt(5);
    }

}
