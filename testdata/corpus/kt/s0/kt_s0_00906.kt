package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome7430 {
    data class Success(val payload: String) : Outcome7430()
    data object Failure : Outcome7430()
}

fun describe7430(outcome: Outcome7430): String = when (outcome) {
    is Outcome7430.Success -> outcome.payload
    Outcome7430.Failure -> "failed"
}
data class Pair7431<A, B>(val first: A, val second: B)

fun combine7431(left: Int, right: Int): Pair7431<Int, Int> {
    return Pair7431(left + 68, right + 36)
}
data class Pair7432<A, B>(val first: A, val second: B)

fun combine7432(left: Int, right: Int): Pair7432<Int, Int> {
    return Pair7432(left + 75, right + 47)
}
