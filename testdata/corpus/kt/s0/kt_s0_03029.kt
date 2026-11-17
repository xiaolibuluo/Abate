package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair24880<A, B>(val first: A, val second: B)

fun combine24880(left: Int, right: Int): Pair24880<Int, Int> {
    return Pair24880(left + 48, right + 40)
}
sealed class Outcome24881 {
    data class Success(val payload: String) : Outcome24881()
    data object Failure : Outcome24881()
}

fun describe24881(outcome: Outcome24881): String = when (outcome) {
    is Outcome24881.Success -> outcome.payload
    Outcome24881.Failure -> "failed"
}
