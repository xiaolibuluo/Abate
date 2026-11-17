package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked25670(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome25671 {
    data class Success(val payload: String) : Outcome25671()
    data object Failure : Outcome25671()
}

fun describe25671(outcome: Outcome25671): String = when (outcome) {
    is Outcome25671.Success -> outcome.payload
    Outcome25671.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked25672(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter25673 {
    var current: Int by Delegates.observable(29) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked25674(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked25675(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull25676(): T? = if (size > 1) this[1] else null

fun render25676(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
