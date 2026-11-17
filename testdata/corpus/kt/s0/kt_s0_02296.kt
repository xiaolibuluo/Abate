package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step18830")
fun step18830(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
class Counter18831 {
    var current: Int by Delegates.observable(47) { _, _, new -> println(new) }

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
data class Pair18832<A, B>(val first: A, val second: B)

fun combine18832(left: Int, right: Int): Pair18832<Int, Int> {
    return Pair18832(left + 54, right + 36)
}
data class Pair18833<A, B>(val first: A, val second: B)

fun combine18833(left: Int, right: Int): Pair18833<Int, Int> {
    return Pair18833(left + 61, right + 47)
}
fun <T> List<T>.secondOrNull18834(): T? = if (size > 1) this[1] else null

fun render18834(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome18835 {
    data class Success(val payload: String) : Outcome18835()
    data object Failure : Outcome18835()
}

fun describe18835(outcome: Outcome18835): String = when (outcome) {
    is Outcome18835.Success -> outcome.payload
    Outcome18835.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked18836(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
