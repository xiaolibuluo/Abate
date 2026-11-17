package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24360 {
    data class Success(val payload: String) : Outcome24360()
    data object Failure : Outcome24360()
}

fun describe24360(outcome: Outcome24360): String = when (outcome) {
    is Outcome24360.Success -> outcome.payload
    Outcome24360.Failure -> "failed"
}
