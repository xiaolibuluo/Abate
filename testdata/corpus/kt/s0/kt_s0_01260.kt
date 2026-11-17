package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair10400<A, B>(val first: A, val second: B)

fun combine10400(left: Int, right: Int): Pair10400<Int, Int> {
    return Pair10400(left + 42, right + 30)
}
sealed class Outcome10401 {
    data class Success(val payload: String) : Outcome10401()
    data object Failure : Outcome10401()
}

fun describe10401(outcome: Outcome10401): String = when (outcome) {
    is Outcome10401.Success -> outcome.payload
    Outcome10401.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked10402(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step10403")
fun step10403(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 63
}
