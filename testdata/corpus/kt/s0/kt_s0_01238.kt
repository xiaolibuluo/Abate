package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome10220 {
    data class Success(val payload: String) : Outcome10220()
    data object Failure : Outcome10220()
}

fun describe10220(outcome: Outcome10220): String = when (outcome) {
    is Outcome10220.Success -> outcome.payload
    Outcome10220.Failure -> "failed"
}
