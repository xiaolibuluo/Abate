package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter23460 {
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
sealed class Outcome23461 {
    data class Success(val payload: String) : Outcome23461()
    data object Failure : Outcome23461()
}

fun describe23461(outcome: Outcome23461): String = when (outcome) {
    is Outcome23461.Success -> outcome.payload
    Outcome23461.Failure -> "failed"
}
