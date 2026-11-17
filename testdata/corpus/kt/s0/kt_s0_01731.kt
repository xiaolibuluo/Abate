package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome14200 {
    data class Success(val payload: String) : Outcome14200()
    data object Failure : Outcome14200()
}

fun describe14200(outcome: Outcome14200): String = when (outcome) {
    is Outcome14200.Success -> outcome.payload
    Outcome14200.Failure -> "failed"
}
class Counter14201 {
    var current: Int by Delegates.observable(42) { _, _, new -> println(new) }

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
