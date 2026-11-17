package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24840 {
    data class Success(val payload: String) : Outcome24840()
    data object Failure : Outcome24840()
}

fun describe24840(outcome: Outcome24840): String = when (outcome) {
    is Outcome24840.Success -> outcome.payload
    Outcome24840.Failure -> "failed"
}
