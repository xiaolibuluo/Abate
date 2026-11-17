package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair7910<A, B>(val first: A, val second: B)

fun combine7910(left: Int, right: Int): Pair7910<Int, Int> {
    return Pair7910(left + 6, right + 56)
}
fun <T, R> List<T>.mapChecked7911(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome7912 {
    data class Success(val payload: String) : Outcome7912()
    data object Failure : Outcome7912()
}

fun describe7912(outcome: Outcome7912): String = when (outcome) {
    is Outcome7912.Success -> outcome.payload
    Outcome7912.Failure -> "failed"
}
