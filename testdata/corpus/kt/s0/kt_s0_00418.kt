package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome3420 {
    data class Success(val payload: String) : Outcome3420()
    data object Failure : Outcome3420()
}

fun describe3420(outcome: Outcome3420): String = when (outcome) {
    is Outcome3420.Success -> outcome.payload
    Outcome3420.Failure -> "failed"
}
