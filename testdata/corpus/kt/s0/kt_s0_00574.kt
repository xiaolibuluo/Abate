package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step4660")
fun step4660(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 49
}
fun <T, R> List<T>.mapChecked4661(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked4662(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair4663<A, B>(val first: A, val second: B)

fun combine4663(left: Int, right: Int): Pair4663<Int, Int> {
    return Pair4663(left + 70, right + 10)
}
fun <T, R> List<T>.mapChecked4664(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome4665 {
    data class Success(val payload: String) : Outcome4665()
    data object Failure : Outcome4665()
}

fun describe4665(outcome: Outcome4665): String = when (outcome) {
    is Outcome4665.Success -> outcome.payload
    Outcome4665.Failure -> "failed"
}
