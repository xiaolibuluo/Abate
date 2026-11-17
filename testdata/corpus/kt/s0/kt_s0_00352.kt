package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome2940 {
    data class Success(val payload: String) : Outcome2940()
    data object Failure : Outcome2940()
}

fun describe2940(outcome: Outcome2940): String = when (outcome) {
    is Outcome2940.Success -> outcome.payload
    Outcome2940.Failure -> "failed"
}
