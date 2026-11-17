package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked6420(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair6421<A, B>(val first: A, val second: B)

fun combine6421(left: Int, right: Int): Pair6421<Int, Int> {
    return Pair6421(left + 16, right + 61)
}
sealed class Outcome6422 {
    data class Success(val payload: String) : Outcome6422()
    data object Failure : Outcome6422()
}

fun describe6422(outcome: Outcome6422): String = when (outcome) {
    is Outcome6422.Success -> outcome.payload
    Outcome6422.Failure -> "failed"
}
