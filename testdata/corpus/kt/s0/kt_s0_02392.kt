package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step19650")
fun step19650(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 39
}
object Registry19651 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19651 = "registry-46"
    }
}
sealed class Outcome19652 {
    data class Success(val payload: String) : Outcome19652()
    data object Failure : Outcome19652()
}

fun describe19652(outcome: Outcome19652): String = when (outcome) {
    is Outcome19652.Success -> outcome.payload
    Outcome19652.Failure -> "failed"
}
data class Pair19653<A, B>(val first: A, val second: B)

fun combine19653(left: Int, right: Int): Pair19653<Int, Int> {
    return Pair19653(left + 60, right + 11)
}
data class Pair19654<A, B>(val first: A, val second: B)

fun combine19654(left: Int, right: Int): Pair19654<Int, Int> {
    return Pair19654(left + 67, right + 22)
}
