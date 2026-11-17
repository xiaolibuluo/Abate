package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull18930(): T? = if (size > 1) this[1] else null

fun render18930(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair18931<A, B>(val first: A, val second: B)

fun combine18931(left: Int, right: Int): Pair18931<Int, Int> {
    return Pair18931(left + 80, right + 34)
}
class Counter18932 {
    var current: Int by Delegates.observable(87) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull18933(): T? = if (size > 1) this[1] else null

fun render18933(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step18934")
fun step18934(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 101
}
fun <T, R> List<T>.mapChecked18935(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome18936 {
    data class Success(val payload: String) : Outcome18936()
    data object Failure : Outcome18936()
}

fun describe18936(outcome: Outcome18936): String = when (outcome) {
    is Outcome18936.Success -> outcome.payload
    Outcome18936.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked18937(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
