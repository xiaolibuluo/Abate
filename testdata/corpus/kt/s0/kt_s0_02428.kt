package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19970 {
    data class Success(val payload: String) : Outcome19970()
    data object Failure : Outcome19970()
}

fun describe19970(outcome: Outcome19970): String = when (outcome) {
    is Outcome19970.Success -> outcome.payload
    Outcome19970.Failure -> "failed"
}
