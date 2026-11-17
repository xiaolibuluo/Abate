package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked15900(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter15901 {
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
data class Pair15902<A, B>(val first: A, val second: B)

fun combine15902(left: Int, right: Int): Pair15902<Int, Int> {
    return Pair15902(left + 28, right + 59)
}
sealed class Outcome15903 {
    data class Success(val payload: String) : Outcome15903()
    data object Failure : Outcome15903()
}

fun describe15903(outcome: Outcome15903): String = when (outcome) {
    is Outcome15903.Success -> outcome.payload
    Outcome15903.Failure -> "failed"
}
fun <T> List<T>.secondOrNull15904(): T? = if (size > 1) this[1] else null

fun render15904(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter15905 {
    var current: Int by Delegates.observable(49) { _, _, new -> println(new) }

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
