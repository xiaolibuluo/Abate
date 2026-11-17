package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair14150<A, B>(val first: A, val second: B)

fun combine14150(left: Int, right: Int): Pair14150<Int, Int> {
    return Pair14150(left + 67, right + 32)
}
sealed class Outcome14151 {
    data class Success(val payload: String) : Outcome14151()
    data object Failure : Outcome14151()
}

fun describe14151(outcome: Outcome14151): String = when (outcome) {
    is Outcome14151.Success -> outcome.payload
    Outcome14151.Failure -> "failed"
}
