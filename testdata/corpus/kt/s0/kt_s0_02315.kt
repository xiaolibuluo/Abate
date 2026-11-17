package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19020 {
    data class Success(val payload: String) : Outcome19020()
    data object Failure : Outcome19020()
}

fun describe19020(outcome: Outcome19020): String = when (outcome) {
    is Outcome19020.Success -> outcome.payload
    Outcome19020.Failure -> "failed"
}
