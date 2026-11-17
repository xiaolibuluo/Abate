package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome22460 {
    data class Success(val payload: String) : Outcome22460()
    data object Failure : Outcome22460()
}

fun describe22460(outcome: Outcome22460): String = when (outcome) {
    is Outcome22460.Success -> outcome.payload
    Outcome22460.Failure -> "failed"
}
sealed class Outcome22461 {
    data class Success(val payload: String) : Outcome22461()
    data object Failure : Outcome22461()
}

fun describe22461(outcome: Outcome22461): String = when (outcome) {
    is Outcome22461.Success -> outcome.payload
    Outcome22461.Failure -> "failed"
}
