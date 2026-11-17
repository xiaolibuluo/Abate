package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull15350(): T? = if (size > 1) this[1] else null

fun render15350(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step15351")
fun step15351(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 82
}
sealed class Outcome15352 {
    data class Success(val payload: String) : Outcome15352()
    data object Failure : Outcome15352()
}

fun describe15352(outcome: Outcome15352): String = when (outcome) {
    is Outcome15352.Success -> outcome.payload
    Outcome15352.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked15353(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair15354<A, B>(val first: A, val second: B)

fun combine15354(left: Int, right: Int): Pair15354<Int, Int> {
    return Pair15354(left + 6, right + 62)
}
fun <T, R> List<T>.mapChecked15355(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
