package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12680 {
    data class Success(val payload: String) : Outcome12680()
    data object Failure : Outcome12680()
}

fun describe12680(outcome: Outcome12680): String = when (outcome) {
    is Outcome12680.Success -> outcome.payload
    Outcome12680.Failure -> "failed"
}
