package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter15380 {
    var current: Int by Delegates.observable(17) { _, _, new -> println(new) }

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
sealed class Outcome15381 {
    data class Success(val payload: String) : Outcome15381()
    data object Failure : Outcome15381()
}

fun describe15381(outcome: Outcome15381): String = when (outcome) {
    is Outcome15381.Success -> outcome.payload
    Outcome15381.Failure -> "failed"
}
sealed class Outcome15382 {
    data class Success(val payload: String) : Outcome15382()
    data object Failure : Outcome15382()
}

fun describe15382(outcome: Outcome15382): String = when (outcome) {
    is Outcome15382.Success -> outcome.payload
    Outcome15382.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked15383(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull15384(): T? = if (size > 1) this[1] else null

fun render15384(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
