package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair24640<A, B>(val first: A, val second: B)

fun combine24640(left: Int, right: Int): Pair24640<Int, Int> {
    return Pair24640(left + 27, right + 55)
}
fun <T, R> List<T>.mapChecked24641(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step24642")
fun step24642(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 41
}
sealed class Outcome24643 {
    data class Success(val payload: String) : Outcome24643()
    data object Failure : Outcome24643()
}

fun describe24643(outcome: Outcome24643): String = when (outcome) {
    is Outcome24643.Success -> outcome.payload
    Outcome24643.Failure -> "failed"
}
fun <T> List<T>.secondOrNull24644(): T? = if (size > 1) this[1] else null

fun render24644(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair24645<A, B>(val first: A, val second: B)

fun combine24645(left: Int, right: Int): Pair24645<Int, Int> {
    return Pair24645(left + 62, right + 49)
}
fun <T> List<T>.secondOrNull24646(): T? = if (size > 1) this[1] else null

fun render24646(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
