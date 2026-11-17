package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome4950 {
    data class Success(val payload: String) : Outcome4950()
    data object Failure : Outcome4950()
}

fun describe4950(outcome: Outcome4950): String = when (outcome) {
    is Outcome4950.Success -> outcome.payload
    Outcome4950.Failure -> "failed"
}
data class Pair4951<A, B>(val first: A, val second: B)

fun combine4951(left: Int, right: Int): Pair4951<Int, Int> {
    return Pair4951(left + 45, right + 69)
}
