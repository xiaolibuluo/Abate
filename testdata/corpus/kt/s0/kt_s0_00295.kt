package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome2460 {
    data class Success(val payload: String) : Outcome2460()
    data object Failure : Outcome2460()
}

fun describe2460(outcome: Outcome2460): String = when (outcome) {
    is Outcome2460.Success -> outcome.payload
    Outcome2460.Failure -> "failed"
}
