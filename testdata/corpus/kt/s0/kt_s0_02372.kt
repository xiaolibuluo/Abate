package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19490 {
    data class Success(val payload: String) : Outcome19490()
    data object Failure : Outcome19490()
}

fun describe19490(outcome: Outcome19490): String = when (outcome) {
    is Outcome19490.Success -> outcome.payload
    Outcome19490.Failure -> "failed"
}
