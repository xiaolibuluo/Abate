package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome17910 {
    data class Success(val payload: String) : Outcome17910()
    data object Failure : Outcome17910()
}

fun describe17910(outcome: Outcome17910): String = when (outcome) {
    is Outcome17910.Success -> outcome.payload
    Outcome17910.Failure -> "failed"
}
fun <T> List<T>.secondOrNull17911(): T? = if (size > 1) this[1] else null

fun render17911(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked17912(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step17913")
fun step17913(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 29
}
fun <T, R> List<T>.mapChecked17914(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter17915 {
    var current: Int by Delegates.observable(43) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
fun <T> List<T>.secondOrNull17916(): T? = if (size > 1) this[1] else null

fun render17916(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
