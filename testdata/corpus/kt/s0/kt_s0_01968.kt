package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome16270 {
    data class Success(val payload: String) : Outcome16270()
    data object Failure : Outcome16270()
}

fun describe16270(outcome: Outcome16270): String = when (outcome) {
    is Outcome16270.Success -> outcome.payload
    Outcome16270.Failure -> "failed"
}
sealed class Outcome16271 {
    data class Success(val payload: String) : Outcome16271()
    data object Failure : Outcome16271()
}

fun describe16271(outcome: Outcome16271): String = when (outcome) {
    is Outcome16271.Success -> outcome.payload
    Outcome16271.Failure -> "failed"
}
data class Pair16272<A, B>(val first: A, val second: B)

fun combine16272(left: Int, right: Int): Pair16272<Int, Int> {
    return Pair16272(left + 24, right + 13)
}
data class Pair16273<A, B>(val first: A, val second: B)

fun combine16273(left: Int, right: Int): Pair16273<Int, Int> {
    return Pair16273(left + 31, right + 24)
}
