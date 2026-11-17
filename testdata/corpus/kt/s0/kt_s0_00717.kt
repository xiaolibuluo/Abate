package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter5910 {
    var current: Int by Delegates.observable(25) { _, _, new -> println(new) }

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
sealed class Outcome5911 {
    data class Success(val payload: String) : Outcome5911()
    data object Failure : Outcome5911()
}

fun describe5911(outcome: Outcome5911): String = when (outcome) {
    is Outcome5911.Success -> outcome.payload
    Outcome5911.Failure -> "failed"
}
