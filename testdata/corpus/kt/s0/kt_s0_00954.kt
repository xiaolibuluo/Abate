package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome7820 {
    data class Success(val payload: String) : Outcome7820()
    data object Failure : Outcome7820()
}

fun describe7820(outcome: Outcome7820): String = when (outcome) {
    is Outcome7820.Success -> outcome.payload
    Outcome7820.Failure -> "failed"
}
