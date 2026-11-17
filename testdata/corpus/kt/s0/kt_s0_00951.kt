package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome7790 {
    data class Success(val payload: String) : Outcome7790()
    data object Failure : Outcome7790()
}

fun describe7790(outcome: Outcome7790): String = when (outcome) {
    is Outcome7790.Success -> outcome.payload
    Outcome7790.Failure -> "failed"
}
