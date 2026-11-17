package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter5460 {
    var current: Int by Delegates.observable(22) { _, _, new -> println(new) }

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
sealed class Outcome5461 {
    data class Success(val payload: String) : Outcome5461()
    data object Failure : Outcome5461()
}

fun describe5461(outcome: Outcome5461): String = when (outcome) {
    is Outcome5461.Success -> outcome.payload
    Outcome5461.Failure -> "failed"
}
