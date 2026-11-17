package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked24170(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull24171(): T? = if (size > 1) this[1] else null

fun render24171(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome24172 {
    data class Success(val payload: String) : Outcome24172()
    data object Failure : Outcome24172()
}

fun describe24172(outcome: Outcome24172): String = when (outcome) {
    is Outcome24172.Success -> outcome.payload
    Outcome24172.Failure -> "failed"
}
@JvmName("step24173")
fun step24173(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 19
}
class Counter24174 {
    var current: Int by Delegates.observable(26) { _, _, new -> println(new) }

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
@JvmName("step24175")
fun step24175(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
