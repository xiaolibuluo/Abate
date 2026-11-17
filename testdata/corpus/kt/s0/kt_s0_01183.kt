package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome9730 {
    data class Success(val payload: String) : Outcome9730()
    data object Failure : Outcome9730()
}

fun describe9730(outcome: Outcome9730): String = when (outcome) {
    is Outcome9730.Success -> outcome.payload
    Outcome9730.Failure -> "failed"
}
