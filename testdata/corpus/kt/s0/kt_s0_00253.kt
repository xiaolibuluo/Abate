package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome2060 {
    data class Success(val payload: String) : Outcome2060()
    data object Failure : Outcome2060()
}

fun describe2060(outcome: Outcome2060): String = when (outcome) {
    is Outcome2060.Success -> outcome.payload
    Outcome2060.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked2061(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair2062<A, B>(val first: A, val second: B)

fun combine2062(left: Int, right: Int): Pair2062<Int, Int> {
    return Pair2062(left + 78, right + 9)
}
