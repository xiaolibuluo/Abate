package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked25200(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair25201<A, B>(val first: A, val second: B)

fun combine25201(left: Int, right: Int): Pair25201<Int, Int> {
    return Pair25201(left + 83, right + 31)
}
sealed class Outcome25202 {
    data class Success(val payload: String) : Outcome25202()
    data object Failure : Outcome25202()
}

fun describe25202(outcome: Outcome25202): String = when (outcome) {
    is Outcome25202.Success -> outcome.payload
    Outcome25202.Failure -> "failed"
}
fun <T> List<T>.secondOrNull25203(): T? = if (size > 1) this[1] else null

fun render25203(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked25204(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter25205 {
    var current: Int by Delegates.observable(14) { _, _, new -> println(new) }

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
class Counter25206 {
    var current: Int by Delegates.observable(21) { _, _, new -> println(new) }

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
