package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24910 {
    data class Success(val payload: String) : Outcome24910()
    data object Failure : Outcome24910()
}

fun describe24910(outcome: Outcome24910): String = when (outcome) {
    is Outcome24910.Success -> outcome.payload
    Outcome24910.Failure -> "failed"
}
class Counter24911 {
    var current: Int by Delegates.observable(94) { _, _, new -> println(new) }

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
