package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter22950 {
    var current: Int by Delegates.observable(61) { _, _, new -> println(new) }

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
sealed class Outcome22951 {
    data class Success(val payload: String) : Outcome22951()
    data object Failure : Outcome22951()
}

fun describe22951(outcome: Outcome22951): String = when (outcome) {
    is Outcome22951.Success -> outcome.payload
    Outcome22951.Failure -> "failed"
}
