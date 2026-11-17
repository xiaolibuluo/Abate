package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25480 {
    data class Success(val payload: String) : Outcome25480()
    data object Failure : Outcome25480()
}

fun describe25480(outcome: Outcome25480): String = when (outcome) {
    is Outcome25480.Success -> outcome.payload
    Outcome25480.Failure -> "failed"
}
class Counter25481 {
    var current: Int by Delegates.observable(59) { _, _, new -> println(new) }

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
@JvmName("step25482")
fun step25482(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 66
}
