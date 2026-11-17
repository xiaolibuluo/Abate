package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome17560 {
    data class Success(val payload: String) : Outcome17560()
    data object Failure : Outcome17560()
}

fun describe17560(outcome: Outcome17560): String = when (outcome) {
    is Outcome17560.Success -> outcome.payload
    Outcome17560.Failure -> "failed"
}
