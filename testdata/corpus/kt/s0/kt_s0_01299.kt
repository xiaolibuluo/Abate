package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome10730 {
    data class Success(val payload: String) : Outcome10730()
    data object Failure : Outcome10730()
}

fun describe10730(outcome: Outcome10730): String = when (outcome) {
    is Outcome10730.Success -> outcome.payload
    Outcome10730.Failure -> "failed"
}
