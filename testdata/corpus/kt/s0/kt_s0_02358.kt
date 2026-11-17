package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull19380(): T? = if (size > 1) this[1] else null

fun render19380(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry19381 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19381 = "registry-83"
    }
}
fun <T, R> List<T>.mapChecked19382(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair19383<A, B>(val first: A, val second: B)

fun combine19383(left: Int, right: Int): Pair19383<Int, Int> {
    return Pair19383(left + 97, right + 66)
}
@JvmName("step19384")
fun step19384(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 7
}
sealed class Outcome19385 {
    data class Success(val payload: String) : Outcome19385()
    data object Failure : Outcome19385()
}

fun describe19385(outcome: Outcome19385): String = when (outcome) {
    is Outcome19385.Success -> outcome.payload
    Outcome19385.Failure -> "failed"
}
fun <T> List<T>.secondOrNull19386(): T? = if (size > 1) this[1] else null

fun render19386(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome19387 {
    data class Success(val payload: String) : Outcome19387()
    data object Failure : Outcome19387()
}

fun describe19387(outcome: Outcome19387): String = when (outcome) {
    is Outcome19387.Success -> outcome.payload
    Outcome19387.Failure -> "failed"
}
