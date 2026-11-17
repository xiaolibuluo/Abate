package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair23210<A, B>(val first: A, val second: B)

fun combine23210(left: Int, right: Int): Pair23210<Int, Int> {
    return Pair23210(left + 11, right + 30)
}
fun <T, R> List<T>.mapChecked23211(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter23212 {
    var current: Int by Delegates.observable(25) { _, _, new -> println(new) }

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
@JvmName("step23213")
fun step23213(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 32
}
fun <T> List<T>.secondOrNull23214(): T? = if (size > 1) this[1] else null

fun render23214(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked23215(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome23216 {
    data class Success(val payload: String) : Outcome23216()
    data object Failure : Outcome23216()
}

fun describe23216(outcome: Outcome23216): String = when (outcome) {
    is Outcome23216.Success -> outcome.payload
    Outcome23216.Failure -> "failed"
}
