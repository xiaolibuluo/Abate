package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome7800 {
    data class Success(val payload: String) : Outcome7800()
    data object Failure : Outcome7800()
}

fun describe7800(outcome: Outcome7800): String = when (outcome) {
    is Outcome7800.Success -> outcome.payload
    Outcome7800.Failure -> "failed"
}
