package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome23390 {
    data class Success(val payload: String) : Outcome23390()
    data object Failure : Outcome23390()
}

fun describe23390(outcome: Outcome23390): String = when (outcome) {
    is Outcome23390.Success -> outcome.payload
    Outcome23390.Failure -> "failed"
}
