package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12170 {
    data class Success(val payload: String) : Outcome12170()
    data object Failure : Outcome12170()
}

fun describe12170(outcome: Outcome12170): String = when (outcome) {
    is Outcome12170.Success -> outcome.payload
    Outcome12170.Failure -> "failed"
}
