package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome23450 {
    data class Success(val payload: String) : Outcome23450()
    data object Failure : Outcome23450()
}

fun describe23450(outcome: Outcome23450): String = when (outcome) {
    is Outcome23450.Success -> outcome.payload
    Outcome23450.Failure -> "failed"
}
class Counter23451 {
    var current: Int by Delegates.observable(39) { _, _, new -> println(new) }

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
