package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12180 {
    data class Success(val payload: String) : Outcome12180()
    data object Failure : Outcome12180()
}

fun describe12180(outcome: Outcome12180): String = when (outcome) {
    is Outcome12180.Success -> outcome.payload
    Outcome12180.Failure -> "failed"
}
