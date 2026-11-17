package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome17270 {
    data class Success(val payload: String) : Outcome17270()
    data object Failure : Outcome17270()
}

fun describe17270(outcome: Outcome17270): String = when (outcome) {
    is Outcome17270.Success -> outcome.payload
    Outcome17270.Failure -> "failed"
}
@JvmName("step17271")
fun step17271(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 56
}
@JvmName("step17272")
fun step17272(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 63
}
fun <T, R> List<T>.mapChecked17273(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked17274(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull17275(): T? = if (size > 1) this[1] else null

fun render17275(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
