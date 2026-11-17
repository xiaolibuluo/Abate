package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair15690<A, B>(val first: A, val second: B)

fun combine15690(left: Int, right: Int): Pair15690<Int, Int> {
    return Pair15690(left + 32, right + 12)
}
sealed class Outcome15691 {
    data class Success(val payload: String) : Outcome15691()
    data object Failure : Outcome15691()
}

fun describe15691(outcome: Outcome15691): String = when (outcome) {
    is Outcome15691.Success -> outcome.payload
    Outcome15691.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked15692(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
