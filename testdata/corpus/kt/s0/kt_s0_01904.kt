package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter15700 {
    var current: Int by Delegates.observable(45) { _, _, new -> println(new) }

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
@JvmName("step15701")
fun step15701(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 52
}
sealed class Outcome15702 {
    data class Success(val payload: String) : Outcome15702()
    data object Failure : Outcome15702()
}

fun describe15702(outcome: Outcome15702): String = when (outcome) {
    is Outcome15702.Success -> outcome.payload
    Outcome15702.Failure -> "failed"
}
