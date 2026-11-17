package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome18990 {
    data class Success(val payload: String) : Outcome18990()
    data object Failure : Outcome18990()
}

fun describe18990(outcome: Outcome18990): String = when (outcome) {
    is Outcome18990.Success -> outcome.payload
    Outcome18990.Failure -> "failed"
}
