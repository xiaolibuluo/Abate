package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome23850 {
    data class Success(val payload: String) : Outcome23850()
    data object Failure : Outcome23850()
}

fun describe23850(outcome: Outcome23850): String = when (outcome) {
    is Outcome23850.Success -> outcome.payload
    Outcome23850.Failure -> "failed"
}
