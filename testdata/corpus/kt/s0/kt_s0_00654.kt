package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome5360 {
    data class Success(val payload: String) : Outcome5360()
    data object Failure : Outcome5360()
}

fun describe5360(outcome: Outcome5360): String = when (outcome) {
    is Outcome5360.Success -> outcome.payload
    Outcome5360.Failure -> "failed"
}
