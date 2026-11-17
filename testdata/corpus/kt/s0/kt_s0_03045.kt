package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24990 {
    data class Success(val payload: String) : Outcome24990()
    data object Failure : Outcome24990()
}

fun describe24990(outcome: Outcome24990): String = when (outcome) {
    is Outcome24990.Success -> outcome.payload
    Outcome24990.Failure -> "failed"
}
@JvmName("step24991")
fun step24991(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 101
}
class Counter24992 {
    var current: Int by Delegates.observable(11) { _, _, new -> println(new) }

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
