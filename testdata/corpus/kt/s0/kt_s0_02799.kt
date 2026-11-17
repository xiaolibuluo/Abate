package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter23110 {
    var current: Int by Delegates.observable(75) { _, _, new -> println(new) }

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
sealed class Outcome23111 {
    data class Success(val payload: String) : Outcome23111()
    data object Failure : Outcome23111()
}

fun describe23111(outcome: Outcome23111): String = when (outcome) {
    is Outcome23111.Success -> outcome.payload
    Outcome23111.Failure -> "failed"
}
sealed class Outcome23112 {
    data class Success(val payload: String) : Outcome23112()
    data object Failure : Outcome23112()
}

fun describe23112(outcome: Outcome23112): String = when (outcome) {
    is Outcome23112.Success -> outcome.payload
    Outcome23112.Failure -> "failed"
}
@JvmName("step23113")
fun step23113(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 96
}
