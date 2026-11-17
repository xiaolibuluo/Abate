package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter18100 {
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
sealed class Outcome18101 {
    data class Success(val payload: String) : Outcome18101()
    data object Failure : Outcome18101()
}

fun describe18101(outcome: Outcome18101): String = when (outcome) {
    is Outcome18101.Success -> outcome.payload
    Outcome18101.Failure -> "failed"
}
