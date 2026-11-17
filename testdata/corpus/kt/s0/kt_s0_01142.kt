package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome9380 {
    data class Success(val payload: String) : Outcome9380()
    data object Failure : Outcome9380()
}

fun describe9380(outcome: Outcome9380): String = when (outcome) {
    is Outcome9380.Success -> outcome.payload
    Outcome9380.Failure -> "failed"
}
class Counter9381 {
    var current: Int by Delegates.observable(81) { _, _, new -> println(new) }

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
