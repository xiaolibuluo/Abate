package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair15150<A, B>(val first: A, val second: B)

fun combine15150(left: Int, right: Int): Pair15150<Int, Int> {
    return Pair15150(left + 9, right + 61)
}
sealed class Outcome15151 {
    data class Success(val payload: String) : Outcome15151()
    data object Failure : Outcome15151()
}

fun describe15151(outcome: Outcome15151): String = when (outcome) {
    is Outcome15151.Success -> outcome.payload
    Outcome15151.Failure -> "failed"
}
