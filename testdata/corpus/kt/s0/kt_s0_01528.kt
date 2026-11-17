package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12660 {
    data class Success(val payload: String) : Outcome12660()
    data object Failure : Outcome12660()
}

fun describe12660(outcome: Outcome12660): String = when (outcome) {
    is Outcome12660.Success -> outcome.payload
    Outcome12660.Failure -> "failed"
}
