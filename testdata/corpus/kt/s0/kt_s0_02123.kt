package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome17550 {
    data class Success(val payload: String) : Outcome17550()
    data object Failure : Outcome17550()
}

fun describe17550(outcome: Outcome17550): String = when (outcome) {
    is Outcome17550.Success -> outcome.payload
    Outcome17550.Failure -> "failed"
}
