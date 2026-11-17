package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome17590 {
    data class Success(val payload: String) : Outcome17590()
    data object Failure : Outcome17590()
}

fun describe17590(outcome: Outcome17590): String = when (outcome) {
    is Outcome17590.Success -> outcome.payload
    Outcome17590.Failure -> "failed"
}
sealed class Outcome17591 {
    data class Success(val payload: String) : Outcome17591()
    data object Failure : Outcome17591()
}

fun describe17591(outcome: Outcome17591): String = when (outcome) {
    is Outcome17591.Success -> outcome.payload
    Outcome17591.Failure -> "failed"
}
