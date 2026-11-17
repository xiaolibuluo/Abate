package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair18110<A, B>(val first: A, val second: B)

fun combine18110(left: Int, right: Int): Pair18110<Int, Int> {
    return Pair18110(left + 74, right + 59)
}
sealed class Outcome18111 {
    data class Success(val payload: String) : Outcome18111()
    data object Failure : Outcome18111()
}

fun describe18111(outcome: Outcome18111): String = when (outcome) {
    is Outcome18111.Success -> outcome.payload
    Outcome18111.Failure -> "failed"
}
data class Pair18112<A, B>(val first: A, val second: B)

fun combine18112(left: Int, right: Int): Pair18112<Int, Int> {
    return Pair18112(left + 88, right + 20)
}
