package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter2070 {
    var current: Int by Delegates.observable(77) { _, _, new -> println(new) }

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
sealed class Outcome2071 {
    data class Success(val payload: String) : Outcome2071()
    data object Failure : Outcome2071()
}

fun describe2071(outcome: Outcome2071): String = when (outcome) {
    is Outcome2071.Success -> outcome.payload
    Outcome2071.Failure -> "failed"
}
