package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter25520 {
    var current: Int by Delegates.observable(7) { _, _, new -> println(new) }

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
sealed class Outcome25521 {
    data class Success(val payload: String) : Outcome25521()
    data object Failure : Outcome25521()
}

fun describe25521(outcome: Outcome25521): String = when (outcome) {
    is Outcome25521.Success -> outcome.payload
    Outcome25521.Failure -> "failed"
}
sealed class Outcome25522 {
    data class Success(val payload: String) : Outcome25522()
    data object Failure : Outcome25522()
}

fun describe25522(outcome: Outcome25522): String = when (outcome) {
    is Outcome25522.Success -> outcome.payload
    Outcome25522.Failure -> "failed"
}
data class Pair25523<A, B>(val first: A, val second: B)

fun combine25523(left: Int, right: Int): Pair25523<Int, Int> {
    return Pair25523(left + 28, right + 33)
}
