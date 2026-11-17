package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step14410")
fun step14410(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 17
}
class Counter14411 {
    var current: Int by Delegates.observable(24) { _, _, new -> println(new) }

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
data class Pair14412<A, B>(val first: A, val second: B)

fun combine14412(left: Int, right: Int): Pair14412<Int, Int> {
    return Pair14412(left + 31, right + 53)
}
sealed class Outcome14413 {
    data class Success(val payload: String) : Outcome14413()
    data object Failure : Outcome14413()
}

fun describe14413(outcome: Outcome14413): String = when (outcome) {
    is Outcome14413.Success -> outcome.payload
    Outcome14413.Failure -> "failed"
}
sealed class Outcome14414 {
    data class Success(val payload: String) : Outcome14414()
    data object Failure : Outcome14414()
}

fun describe14414(outcome: Outcome14414): String = when (outcome) {
    is Outcome14414.Success -> outcome.payload
    Outcome14414.Failure -> "failed"
}
