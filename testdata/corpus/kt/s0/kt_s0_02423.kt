package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19920 {
    data class Success(val payload: String) : Outcome19920()
    data object Failure : Outcome19920()
}

fun describe19920(outcome: Outcome19920): String = when (outcome) {
    is Outcome19920.Success -> outcome.payload
    Outcome19920.Failure -> "failed"
}
@JvmName("step19921")
fun step19921(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 9
}
fun <T> List<T>.secondOrNull19922(): T? = if (size > 1) this[1] else null

fun render19922(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter19923 {
    var current: Int by Delegates.observable(23) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull19924(): T? = if (size > 1) this[1] else null

fun render19924(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull19925(): T? = if (size > 1) this[1] else null

fun render19925(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair19926<A, B>(val first: A, val second: B)

fun combine19926(left: Int, right: Int): Pair19926<Int, Int> {
    return Pair19926(left + 44, right + 50)
}
@JvmName("step19927")
fun step19927(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 51
}
fun <T, R> List<T>.mapChecked19928(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
