package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome30 {
    data class Success(val payload: String) : Outcome30()
    data object Failure : Outcome30()
}

fun describe30(outcome: Outcome30): String = when (outcome) {
    is Outcome30.Success -> outcome.payload
    Outcome30.Failure -> "failed"
}
