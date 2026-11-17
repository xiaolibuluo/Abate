package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome3660 {
    data class Success(val payload: String) : Outcome3660()
    data object Failure : Outcome3660()
}

fun describe3660(outcome: Outcome3660): String = when (outcome) {
    is Outcome3660.Success -> outcome.payload
    Outcome3660.Failure -> "failed"
}
data class Pair3661<A, B>(val first: A, val second: B)

fun combine3661(left: Int, right: Int): Pair3661<Int, Int> {
    return Pair3661(left + 17, right + 20)
}
@JvmName("step3662")
fun step3662(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 24
}
fun <T, R> List<T>.mapChecked3663(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull3664(): T? = if (size > 1) this[1] else null

fun render3664(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step3665")
fun step3665(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}
