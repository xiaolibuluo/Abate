package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome8760 {
    data class Success(val payload: String) : Outcome8760()
    data object Failure : Outcome8760()
}

fun describe8760(outcome: Outcome8760): String = when (outcome) {
    is Outcome8760.Success -> outcome.payload
    Outcome8760.Failure -> "failed"
}
