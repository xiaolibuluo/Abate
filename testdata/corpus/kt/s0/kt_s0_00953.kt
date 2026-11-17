package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome7810 {
    data class Success(val payload: String) : Outcome7810()
    data object Failure : Outcome7810()
}

fun describe7810(outcome: Outcome7810): String = when (outcome) {
    is Outcome7810.Success -> outcome.payload
    Outcome7810.Failure -> "failed"
}
