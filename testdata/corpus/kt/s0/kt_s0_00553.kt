package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step4500")
fun step4500(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 35
}
data class Pair4501<A, B>(val first: A, val second: B)

fun combine4501(left: Int, right: Int): Pair4501<Int, Int> {
    return Pair4501(left + 42, right + 59)
}
fun <T, R> List<T>.mapChecked4502(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome4503 {
    data class Success(val payload: String) : Outcome4503()
    data object Failure : Outcome4503()
}

fun describe4503(outcome: Outcome4503): String = when (outcome) {
    is Outcome4503.Success -> outcome.payload
    Outcome4503.Failure -> "failed"
}
