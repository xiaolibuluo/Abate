package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome3020 {
    data class Success(val payload: String) : Outcome3020()
    data object Failure : Outcome3020()
}

fun describe3020(outcome: Outcome3020): String = when (outcome) {
    is Outcome3020.Success -> outcome.payload
    Outcome3020.Failure -> "failed"
}
class Counter3021 {
    var current: Int by Delegates.observable(58) { _, _, new -> println(new) }

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
