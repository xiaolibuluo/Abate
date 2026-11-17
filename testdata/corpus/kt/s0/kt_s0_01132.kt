package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome9290 {
    data class Success(val payload: String) : Outcome9290()
    data object Failure : Outcome9290()
}

fun describe9290(outcome: Outcome9290): String = when (outcome) {
    is Outcome9290.Success -> outcome.payload
    Outcome9290.Failure -> "failed"
}
