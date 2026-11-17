package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked17640(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair17641<A, B>(val first: A, val second: B)

fun combine17641(left: Int, right: Int): Pair17641<Int, Int> {
    return Pair17641(left + 52, right + 46)
}
@JvmName("step17642")
fun step17642(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 59
}
sealed class Outcome17643 {
    data class Success(val payload: String) : Outcome17643()
    data object Failure : Outcome17643()
}

fun describe17643(outcome: Outcome17643): String = when (outcome) {
    is Outcome17643.Success -> outcome.payload
    Outcome17643.Failure -> "failed"
}
