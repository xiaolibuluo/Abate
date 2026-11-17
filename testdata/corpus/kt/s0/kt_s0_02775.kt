package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome22910 {
    data class Success(val payload: String) : Outcome22910()
    data object Failure : Outcome22910()
}

fun describe22910(outcome: Outcome22910): String = when (outcome) {
    is Outcome22910.Success -> outcome.payload
    Outcome22910.Failure -> "failed"
}
