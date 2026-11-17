package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19760 {
    data class Success(val payload: String) : Outcome19760()
    data object Failure : Outcome19760()
}

fun describe19760(outcome: Outcome19760): String = when (outcome) {
    is Outcome19760.Success -> outcome.payload
    Outcome19760.Failure -> "failed"
}
class Counter19761 {
    var current: Int by Delegates.observable(92) { _, _, new -> println(new) }

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
sealed class Outcome19762 {
    data class Success(val payload: String) : Outcome19762()
    data object Failure : Outcome19762()
}

fun describe19762(outcome: Outcome19762): String = when (outcome) {
    is Outcome19762.Success -> outcome.payload
    Outcome19762.Failure -> "failed"
}
data class Pair19763<A, B>(val first: A, val second: B)

fun combine19763(left: Int, right: Int): Pair19763<Int, Int> {
    return Pair19763(left + 9, right + 27)
}
@JvmName("step19764")
fun step19764(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 16
}
