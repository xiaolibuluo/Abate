package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome7320 {
    data class Success(val payload: String) : Outcome7320()
    data object Failure : Outcome7320()
}

fun describe7320(outcome: Outcome7320): String = when (outcome) {
    is Outcome7320.Success -> outcome.payload
    Outcome7320.Failure -> "failed"
}
