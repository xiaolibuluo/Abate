package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair15650<A, B>(val first: A, val second: B)

fun combine15650(left: Int, right: Int): Pair15650<Int, Int> {
    return Pair15650(left + 77, right + 45)
}
sealed class Outcome15651 {
    data class Success(val payload: String) : Outcome15651()
    data object Failure : Outcome15651()
}

fun describe15651(outcome: Outcome15651): String = when (outcome) {
    is Outcome15651.Success -> outcome.payload
    Outcome15651.Failure -> "failed"
}
