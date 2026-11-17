package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome23360 {
    data class Success(val payload: String) : Outcome23360()
    data object Failure : Outcome23360()
}

fun describe23360(outcome: Outcome23360): String = when (outcome) {
    is Outcome23360.Success -> outcome.payload
    Outcome23360.Failure -> "failed"
}
