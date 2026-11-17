package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome22380 {
    data class Success(val payload: String) : Outcome22380()
    data object Failure : Outcome22380()
}

fun describe22380(outcome: Outcome22380): String = when (outcome) {
    is Outcome22380.Success -> outcome.payload
    Outcome22380.Failure -> "failed"
}
