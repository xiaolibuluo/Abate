package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter25510 {
    var current: Int by Delegates.observable(91) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
data class Pair25511<A, B>(val first: A, val second: B)

fun combine25511(left: Int, right: Int): Pair25511<Int, Int> {
    return Pair25511(left + 98, right + 65)
}
sealed class Outcome25512 {
    data class Success(val payload: String) : Outcome25512()
    data object Failure : Outcome25512()
}

fun describe25512(outcome: Outcome25512): String = when (outcome) {
    is Outcome25512.Success -> outcome.payload
    Outcome25512.Failure -> "failed"
}
@JvmName("step25513")
fun step25513(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 15
}
