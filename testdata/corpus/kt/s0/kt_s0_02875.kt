package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked23720(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked23721(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step23722")
fun step23722(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 9
}
fun <T, R> List<T>.mapChecked23723(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step23724")
fun step23724(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 23
}
fun <T> List<T>.secondOrNull23725(): T? = if (size > 1) this[1] else null

fun render23725(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome23726 {
    data class Success(val payload: String) : Outcome23726()
    data object Failure : Outcome23726()
}

fun describe23726(outcome: Outcome23726): String = when (outcome) {
    is Outcome23726.Success -> outcome.payload
    Outcome23726.Failure -> "failed"
}
@JvmName("step23727")
fun step23727(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 44
}
