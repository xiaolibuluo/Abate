package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome23870 {
    data class Success(val payload: String) : Outcome23870()
    data object Failure : Outcome23870()
}

fun describe23870(outcome: Outcome23870): String = when (outcome) {
    is Outcome23870.Success -> outcome.payload
    Outcome23870.Failure -> "failed"
}
