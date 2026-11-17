package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked16910(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair16911<A, B>(val first: A, val second: B)

fun combine16911(left: Int, right: Int): Pair16911<Int, Int> {
    return Pair16911(left + 73, right + 23)
}
fun <T> List<T>.secondOrNull16912(): T? = if (size > 1) this[1] else null

fun render16912(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome16913 {
    data class Success(val payload: String) : Outcome16913()
    data object Failure : Outcome16913()
}

fun describe16913(outcome: Outcome16913): String = when (outcome) {
    is Outcome16913.Success -> outcome.payload
    Outcome16913.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked16914(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step16915")
fun step16915(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 101
}
fun <T> List<T>.secondOrNull16916(): T? = if (size > 1) this[1] else null

fun render16916(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome16917 {
    data class Success(val payload: String) : Outcome16917()
    data object Failure : Outcome16917()
}

fun describe16917(outcome: Outcome16917): String = when (outcome) {
    is Outcome16917.Success -> outcome.payload
    Outcome16917.Failure -> "failed"
}
