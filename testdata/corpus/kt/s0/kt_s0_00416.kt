package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome3410 {
    data class Success(val payload: String) : Outcome3410()
    data object Failure : Outcome3410()
}

fun describe3410(outcome: Outcome3410): String = when (outcome) {
    is Outcome3410.Success -> outcome.payload
    Outcome3410.Failure -> "failed"
}
