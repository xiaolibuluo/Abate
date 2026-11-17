package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome18090 {
    data class Success(val payload: String) : Outcome18090()
    data object Failure : Outcome18090()
}

fun describe18090(outcome: Outcome18090): String = when (outcome) {
    is Outcome18090.Success -> outcome.payload
    Outcome18090.Failure -> "failed"
}
data class Pair18091<A, B>(val first: A, val second: B)

fun combine18091(left: Int, right: Int): Pair18091<Int, Int> {
    return Pair18091(left + 55, right + 56)
}
