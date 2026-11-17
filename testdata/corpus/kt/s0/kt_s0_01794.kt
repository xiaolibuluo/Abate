package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome14730 {
    data class Success(val payload: String) : Outcome14730()
    data object Failure : Outcome14730()
}

fun describe14730(outcome: Outcome14730): String = when (outcome) {
    is Outcome14730.Success -> outcome.payload
    Outcome14730.Failure -> "failed"
}
class Counter14731 {
    var current: Int by Delegates.observable(52) { _, _, new -> println(new) }

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
