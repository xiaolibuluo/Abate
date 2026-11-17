package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome18530 {
    data class Success(val payload: String) : Outcome18530()
    data object Failure : Outcome18530()
}

fun describe18530(outcome: Outcome18530): String = when (outcome) {
    is Outcome18530.Success -> outcome.payload
    Outcome18530.Failure -> "failed"
}
